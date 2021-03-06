package core

import (
	"fmt"

	"../parser"
)

type Buffer struct {
	Bytes []byte
	Pos   int
}

func NewBuffer(bytes []byte) Buffer {
	return Buffer{bytes, 0}
}

func (buffer *Buffer) Put(bytes []byte) {
	buffer.Bytes = append(buffer.Bytes, bytes...)
}

func (buffer *Buffer) Get(c int) []byte {
	b := buffer.Bytes[buffer.Pos : buffer.Pos+c]
	buffer.Pos += c
	return b
}

type Types struct {
	Types []Type
}

type Type struct {
	Name               string
	Fields, MetaFields []Field
}

type Field struct {
	Name, Prim, Type, Sec, Direct, Default string
}

type Listener struct {
	parser.DatatableListener

	IsMetaContext            bool
	IsTypeContext            bool
	IsGroupContext           bool
	IsTypeDefinitionContext  bool
	IsFieldContext           bool
	IsFieldDefinitionContext bool
	Types                    []Type
	CurrentType              Type
	CurrentField             Field
}

func (this *Listener) Print() {
	for _, t := range this.Types {
		fmt.Println("Type", t.Name)
		for _, field := range t.Fields {
			fmt.Println("Field", field.Name, field.Prim, field.Type, field.Sec, field.Direct, field.Default)
		}
		for _, field := range t.MetaFields {
			fmt.Println("Meta", field.Name, field.Prim, field.Type, field.Sec, field.Direct, field.Default)
		}
		fmt.Println("")
	}
}

func NewListener() *Listener {
	return new(Listener)
}

func (this *Listener) Init() {
	this.Clear()
}

func (this *Listener) IsMeta() bool {
	return this.IsMetaContext
}

func (this *Listener) SetMeta(b bool) {
	this.IsMetaContext = b
}

func (this *Listener) IsType() bool {
	return this.IsTypeContext
}

func (this *Listener) SetTypeC(b bool) {
	this.IsTypeContext = b
}

func (this *Listener) IsGroup() bool {
	return this.IsGroupContext
}

func (this *Listener) SetGroup(b bool) {
	this.IsGroupContext = b
}

func (this *Listener) IsTypeDef() bool {
	return this.IsTypeDefinitionContext
}

func (this *Listener) SetTypeDef(b bool) {
	this.IsTypeDefinitionContext = b
}

func (this *Listener) IsFieldC() bool {
	return this.IsFieldContext
}

func (this *Listener) SetFieldC(b bool) {
	this.IsFieldContext = b
}

func (this *Listener) IsFieldDef() bool {
	return this.IsFieldDefinitionContext
}

func (this *Listener) SetFieldDef(b bool) {
	this.IsFieldDefinitionContext = b
}

func (this *Listener) AddType() {
	if !this.CurrentType.IsEmpty() {
		this.Types = append(this.Types, this.CurrentType)
		this.Clear()
	}
}

func (this *Listener) SetTypeName(s string) {
	this.CurrentType.Name = s
}

func (this *Listener) Clear() {
	this.CurrentType = Type{}
	this.CurrentField = Field{"", "", "", "", "", ""}
}

func (this *Listener) AddField() {
	if this.CurrentField.IsEmpty() {
		return
	}
	this.CurrentType.Fields = append(this.CurrentType.Fields, this.CurrentField)
	this.CurrentField = Field{"", "", "", "", "", ""}
	this.SetFieldDef(false)
	this.SetFieldC(false)
}

func (this *Listener) AddMetaField() {
	if this.CurrentField.IsEmpty() {
		return
	}
	this.CurrentType.MetaFields = append(this.CurrentType.MetaFields, this.CurrentField)
	this.CurrentField = Field{"", "", "", "", "", ""}
}

func (this *Listener) SetPrimary(s string) {
	this.CurrentField.Prim = s
}

func (this *Listener) SetType(s string) {
	this.CurrentField.Type = s
}

func (this *Listener) SetSecondary(s string) {
	this.CurrentField.Sec = s
}

func (this *Listener) SetDirect(s string) {
	this.CurrentField.Direct = s
}

func (this *Listener) SetDefault(s string) {
	this.CurrentField.Default = s
}

func (this *Listener) SetFieldName(s string) {
	this.CurrentField.Name = s
}

func (this *Type) IsEmpty() bool {
	a := false
	if len(this.Fields) == 0 {
		a = true
	}
	return a
}

func (this *Field) IsEmpty() bool {
	return (this.Name == "") && (this.Prim == "") && (this.Type == "") && (this.Sec == "") && (this.Default == "") && (this.Direct == "")
}
