package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"./core"
	lib "./parser"
	"./targets"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func main() {
	args := os.Args
	lines, err := ioutil.ReadFile(args[1])
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
	//listener.Print()
	f, err := os.Create(args[2])
	if err != nil {
		fmt.Println(err)
	}
	if len(args) >= 4 {
		bytes := []byte{}
		switch args[3] {
		case "go":
			std := false
			if len(args) >= 5 {
				if args[4] == "std" {
					std = true
				}
			}
			bytes = targets.Go(core.Types{listener.Types}, std)
		case "json":
			bytes = listener.GenerateJSON()
		}
		f.Write(bytes)
		f.Close()
	} else {
		f.Write(listener.GenerateJSON())
		f.Close()
	}
}
