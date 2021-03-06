package core

import (
	"../parser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func (this *Listener) VisitTerminal(node antlr.TerminalNode) {
}

func (this *Listener) VisitErrorNode(node antlr.ErrorNode) {

}

func (this *Listener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	//fmt.Println(ctx.GetText())
}

func (this *Listener) ExitEveryRule(ctx antlr.ParserRuleContext) {

}

func (this *Listener) EnterDefinition(c *parser.DefinitionContext) {
	if c.GetText() == "Meta" {
		this.SetMeta(true)
	}
}

// EnterTypeDeclaration is called when entering the typeDeclaration production.
func (this *Listener) EnterTypeDeclaration(c *parser.TypeDeclarationContext) {
	this.SetTypeDef(true)
}

// EnterAtom is called when entering the atom production.
func (this *Listener) EnterAtom(c *parser.AtomContext) {
	this.SetDefault(c.GetText())
}

// EnterField is called when entering the field production.
func (this *Listener) EnterField(c *parser.FieldContext) {
	this.SetFieldC(true)
}

// EnterName is called when entering the name production.
func (this *Listener) EnterName(c *parser.NameContext) {
	if this.IsTypeDef() {
		this.SetTypeDef(false)
		this.SetTypeName(c.GetText())
	} else if this.IsFieldC() {
		this.SetFieldName(c.GetText())
		this.SetFieldC(false)
	} else if this.IsFieldDef() {
		this.SetType(c.GetText())
		this.SetFieldDef(false)
	}
}

// EnterTypeDefinition is called when entering the typeDefinition production.
func (this *Listener) EnterTypeDefinition(c *parser.TypeDefinitionContext) {
	//this.SetFieldDef(true)
}

// EnterFieldType is called when entering the fieldType production.
func (this *Listener) EnterFieldType(c *parser.FieldTypeContext) {
	this.SetFieldDef(true)
}

// EnterSize is called when entering the size production.
func (this *Listener) EnterSize(c *parser.SizeContext) {
	if this.IsFieldDef() {
		// primary
		this.SetPrimary(c.GetText())
		return
	}
	this.SetSecondary(c.GetText())
}

// EnterDirect is called when entering the direct production.
func (this *Listener) EnterDirect(c *parser.DirectContext) {
	this.SetDirect(c.GetText())
}

// ExitTypeDeclaration is called when exiting the typeDeclaration production.
func (this *Listener) ExitTypeDeclaration(c *parser.TypeDeclarationContext) {
	this.AddType()
}

// ExitFieldType is called when exiting the fieldType production.
func (this *Listener) ExitFieldType(c *parser.FieldTypeContext) {
	this.SetFieldDef(false)
}

// ExitDeclaration is called when exiting the declaration production.
func (this *Listener) ExitDeclaration(c *parser.DeclarationContext) {
	this.SetMeta(false)
}

// ExitField is called when exiting the field production.
func (this *Listener) ExitField(c *parser.FieldContext) {
	if this.IsMeta() {
		this.AddMetaField()
		return
	}
	this.AddField()
}
