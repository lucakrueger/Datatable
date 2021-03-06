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

// EnterChunk is called when entering the chunk production.
func (this *Listener) EnterChunk(c *parser.ChunkContext) {

}

// EnterDeclaration is called when entering the declaration production.
func (this *Listener) EnterDeclaration(c *parser.DeclarationContext) {

}

// EnterBlock is called when entering the block production.
func (this *Listener) EnterBlock(c *parser.BlockContext) {

}

// EnterDirect is called when entering the direct production.
func (this *Listener) EnterDirect(c *parser.DirectContext) {
	this.SetDirect(c.GetText())
}

// EnterBody is called when entering the body production.
func (this *Listener) EnterBody(c *parser.BodyContext) {

}

// EnterExpression is called when entering the expression production.
func (this *Listener) EnterExpression(c *parser.ExpressionContext) {

}

// ExitDefinition is called when exiting the definition production.
func (this *Listener) ExitDefinition(c *parser.DefinitionContext) {
}

// ExitTypeDeclaration is called when exiting the typeDeclaration production.
func (this *Listener) ExitTypeDeclaration(c *parser.TypeDeclarationContext) {
	this.AddType()
}

// ExitAtom is called when exiting the atom production.
func (this *Listener) ExitAtom(c *parser.AtomContext) {

}

// ExitName is called when exiting the name production.
func (this *Listener) ExitName(c *parser.NameContext) {

}

// ExitTypeDefinition is called when exiting the typeDefinition production.
func (this *Listener) ExitTypeDefinition(c *parser.TypeDefinitionContext) {

}

// ExitFieldType is called when exiting the fieldType production.
func (this *Listener) ExitFieldType(c *parser.FieldTypeContext) {
	this.SetFieldDef(false)
}

// ExitSize is called when exiting the size production.
func (this *Listener) ExitSize(c *parser.SizeContext) {

}

// ExitChunk is called when exiting the chunk production.
func (this *Listener) ExitChunk(c *parser.ChunkContext) {

}

// ExitDirect is called when exiting the direct production.
func (this *Listener) ExitDirect(c *parser.DirectContext) {

}

// ExitDeclaration is called when exiting the declaration production.
func (this *Listener) ExitDeclaration(c *parser.DeclarationContext) {
	this.SetMeta(false)
}

// ExitBlock is called when exiting the block production.
func (this *Listener) ExitBlock(c *parser.BlockContext) {

}

// ExitBody is called when exiting the body production.
func (this *Listener) ExitBody(c *parser.BodyContext) {

}

// ExitExpression is called when exiting the expression production.
func (this *Listener) ExitExpression(c *parser.ExpressionContext) {

}

// ExitField is called when exiting the field production.
func (this *Listener) ExitField(c *parser.FieldContext) {
	if this.IsMeta() {
		this.AddMetaField()
		return
	}
	this.AddField()
}
