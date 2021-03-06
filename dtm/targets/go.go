package targets

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"../core"
	lib "../parser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func Go(types core.Types, std bool) []byte {

	buffer := core.NewBuffer(make([]byte, 0))

	if std {
		// include std libs
		lines, err := ioutil.ReadFile("../libs/std.dt")
		if err != nil {
			fmt.Println(err)
		}

		chars := antlr.NewInputStream(string(lines))
		lexer := lib.NewDatatableLexer(chars)
		tokens := antlr.NewCommonTokenStream(lexer, 0)
		parser := lib.NewDatatableParser(tokens)

		tree := parser.Chunk()
		listener := core.NewListener()
		antlr.ParseTreeWalkerDefault.Walk(listener, tree)
		types.Types = append(types.Types, listener.Types...)
	}

	for _, t := range types.Types {
		buffer.Put([]byte(GoProcessType(t)))
	}

	return buffer.Bytes
}

func GoProcessType(t core.Type) string {
	s := "type " + t.Name + " struct {\n"

	defaultParameters := make([][]string, 0)

	for _, field := range t.Fields {
		fieldstr := "\t"
		if field.Name[0] == '@' {
			fieldstr += field.Name[1:] + " "
			defaultParameters = append(defaultParameters, []string{field.Name[1:], GetType(&field)})
		} else {
			fieldstr += field.Name + " "
		}
		if field.Prim != "" {
			if field.Prim == "?" {
				fieldstr += "[]"
			} else {
				_, eee := strconv.Atoi(field.Prim)
				if eee != nil {
					// variable length
					fieldstr += "[]"
				} else {
					// fixed length
					fieldstr += "[" + field.Prim + "]"
				}
			}
		}
		if field.Direct != "" {
			fieldstr += field.Direct[1 : len(field.Direct)-1]
		} else {
			switch field.Type {
			case "string":
				fieldstr += field.Type
			default:
				fieldstr += field.Type + field.Sec
			}
		}
		fieldstr += "\n"
		s += fieldstr
	}

	s += "}\n\n"

	/*
		Default functions
	*/

	// NewT
	if !HasMeta(&t, "New"+t.Name) {
		newfunc := "func New" + t.Name + "("
		for i, defaults := range defaultParameters {
			newfunc += strings.ToLower(defaults[0]) + " " + defaults[1]
			if i != (len(defaultParameters) - 1) {
				newfunc += ", "
			}
		}
		newfunc += ") " + t.Name + " {\n"
		newfunc += "\treturn " + t.Name + "{"
		for i, field := range t.Fields {
			if field.Default != "" {
				newfunc += field.Default
			} else if IsCons(&field) {
				newfunc += strings.ToLower(field.Name[1:])
			} else {
				reqType := ""
				isarray := false
				if field.Direct != "" {
					reqType = "nil"
				} else {
					reqType = field.Type
				}
				isarray = (field.Prim != "")
				if isarray {
					newfunc += "make([]" + reqType + ", 0)"
				} else {
					switch reqType {
					case "string":
						newfunc += "\"\""
					case "int":
						newfunc += "0"
					case "bool":
						newfunc += "false"
					case "float":
						newfunc += "0.0"
					case "rune":
						newfunc += "''"
					case "byte":
						newfunc += "0"
					default:
						newfunc += "nil"
					}
				}
			}
			if i != (len(t.Fields) - 1) {
				newfunc += ", "
			}
		}
		newfunc += "}\n}\n\n"
		s += newfunc
	}

	// ReadT
	if !HasMeta(&t, "Read"+t.Name) {
		readfunc := "func Read" + t.Name + "(bytes []byte) " + t.Name + " {\n"
		readfunc += "\tbuffer := NewBuffer(bytes)\n\treturn buffer.Fetch" + t.Name + "()\n"
		readfunc += "}\n\n"
		s += readfunc
	}

	// FetchT
	if !HasMeta(&t, "Fetch"+t.Name) {
		fetchfunc := "func (buffer *Buffer) Fetch" + t.Name + "() " + t.Name + "{\n"
		fetchfunc += "\tvar this " + t.Name + "\n"
		for _, field := range t.Fields {
			if IsArray(field) {
				size := ""
				if field.Prim == "?" {
					size = "len(this." + Normalize(field.Name) + ")"
				} else {
					size = "this." + field.Prim
				}
				fetchfunc += "\tfor i := 0; i < " + size + "; i++{\n"
				fetchfunc += "\t\tthis." + Normalize(field.Name) + " = append(this." + Normalize(field.Name) + ", buffer.Fetch" + field.Type + "())\n\t}\n"
			} else {
				fetchfunc += "\tthis." + Normalize(field.Name) + " = buffer.Fetch" + field.Type + field.Sec + "()\n"
			}
		}
		fetchfunc += "\treturn this\n}\n\n"
		s += fetchfunc
	}

	// WriteT
	if !HasMeta(&t, "Write") {
		writefunc := "func (this *" + t.Name + ") Write() []byte {\n"
		writefunc += "\tvar buffer Buffer\n"
		for _, field := range t.Fields {
			if IsArray(field) {
				writefunc += "\tfor _, elem := range this." + Normalize(field.Name) + " {\n"
				writefunc += "\t\tbuffer.Put("
				if IsPrimitive(field.Type) {
					writefunc += "buffer." + strings.Title(field.Type+field.Sec) + "Bytes(this." + Normalize(field.Name) + "))\n"
				} else {
					writefunc += "elem" + ".Write())\n"
				}
				writefunc += "\t}\n"
			} else {
				writefunc += "\tbuffer.Put("
				if IsPrimitive(field.Type) {
					writefunc += "buffer." + strings.Title(field.Type+field.Sec) + "Bytes(this." + Normalize(field.Name) + "))\n"
				} else {
					writefunc += "this." + Normalize(field.Name) + ".Write())\n"
				}
			}
		}
		writefunc += "\treturn buffer.Bytes\n}\n\n"
		s += writefunc
	}

	/*
		Meta functions
	*/
	for _, meta := range t.MetaFields {
		switch meta.Name {
		case "Package":
			s = "package " + meta.Direct[1:len(meta.Direct)-1] + "\n\n" + s
		case "Import":
			s = "import \"" + meta.Direct[1:len(meta.Direct)-1] + "\"\n\n" + s
		default:
			switch meta.Name {
			case "Read" + t.Name:
				s += "func " + meta.Name + "" + meta.Direct[1:len(meta.Direct)-1] + "\n\n"
			case "New" + t.Name:
				s += "func " + meta.Name + "" + meta.Direct[1:len(meta.Direct)-1] + "\n\n"
			default:
				s += "func (this *" + t.Name + ") " + meta.Name + "" + meta.Direct[1:len(meta.Direct)-1] + "\n\n"
			}
		}
	}

	return s
}

func HasMeta(this *core.Type, field string) bool {
	for _, elem := range this.MetaFields {
		if elem.Name == field {
			return true
		}
	}
	return false
}

func GetType(this *core.Field) string {
	s := ""
	if this.Prim != "" {
		s += "[]"
	}
	if this.Direct != "" {
		return s + this.Direct
	}
	return s + this.Type
}

func IsCons(this *core.Field) bool {
	if this.Name[0] == '@' {
		return true
	}
	return false
}

func Normalize(s string) string {
	if s[0] == '@' {
		return s[1:]
	}
	return s
}

func IsArray(this core.Field) bool {
	return (this.Prim != "")
}

func IsPrimitive(s string) bool {
	switch s {
	case "string":
		return true
	case "int":
		return true
	case "bool":
		return true
	case "float":
		return true
	case "rune":
		return true
	case "byte":
		return true
	default:
		return false
	}
}
