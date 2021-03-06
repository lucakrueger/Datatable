// Code generated from D:/Projects/datatable/grammar/Datatable.g4 by ANTLR 4.9. DO NOT EDIT.

package parser // Datatable

import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseDatatableVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseDatatableVisitor) VisitDefinition(ctx *DefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDatatableVisitor) VisitAtom(ctx *AtomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDatatableVisitor) VisitName(ctx *NameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDatatableVisitor) VisitTypeDefinition(ctx *TypeDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDatatableVisitor) VisitFieldType(ctx *FieldTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDatatableVisitor) VisitDirect(ctx *DirectContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDatatableVisitor) VisitSize(ctx *SizeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDatatableVisitor) VisitChunk(ctx *ChunkContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDatatableVisitor) VisitTypeDeclaration(ctx *TypeDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDatatableVisitor) VisitDeclaration(ctx *DeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDatatableVisitor) VisitBlock(ctx *BlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDatatableVisitor) VisitBody(ctx *BodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDatatableVisitor) VisitExpression(ctx *ExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDatatableVisitor) VisitField(ctx *FieldContext) interface{} {
	return v.VisitChildren(ctx)
}
