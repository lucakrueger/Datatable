// Code generated from D:/Projects/datatable/grammar/Datatable.g4 by ANTLR 4.9. DO NOT EDIT.

package parser // Datatable

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by DatatableParser.
type DatatableVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by DatatableParser#definition.
	VisitDefinition(ctx *DefinitionContext) interface{}

	// Visit a parse tree produced by DatatableParser#atom.
	VisitAtom(ctx *AtomContext) interface{}

	// Visit a parse tree produced by DatatableParser#name.
	VisitName(ctx *NameContext) interface{}

	// Visit a parse tree produced by DatatableParser#typeDefinition.
	VisitTypeDefinition(ctx *TypeDefinitionContext) interface{}

	// Visit a parse tree produced by DatatableParser#fieldType.
	VisitFieldType(ctx *FieldTypeContext) interface{}

	// Visit a parse tree produced by DatatableParser#direct.
	VisitDirect(ctx *DirectContext) interface{}

	// Visit a parse tree produced by DatatableParser#size.
	VisitSize(ctx *SizeContext) interface{}

	// Visit a parse tree produced by DatatableParser#chunk.
	VisitChunk(ctx *ChunkContext) interface{}

	// Visit a parse tree produced by DatatableParser#typeDeclaration.
	VisitTypeDeclaration(ctx *TypeDeclarationContext) interface{}

	// Visit a parse tree produced by DatatableParser#declaration.
	VisitDeclaration(ctx *DeclarationContext) interface{}

	// Visit a parse tree produced by DatatableParser#block.
	VisitBlock(ctx *BlockContext) interface{}

	// Visit a parse tree produced by DatatableParser#body.
	VisitBody(ctx *BodyContext) interface{}

	// Visit a parse tree produced by DatatableParser#expression.
	VisitExpression(ctx *ExpressionContext) interface{}

	// Visit a parse tree produced by DatatableParser#field.
	VisitField(ctx *FieldContext) interface{}
}
