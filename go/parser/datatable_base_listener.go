// Code generated from D:/Projects/datatable/grammar/Datatable.g4 by ANTLR 4.9. DO NOT EDIT.

package parser // Datatable

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseDatatableListener is a complete listener for a parse tree produced by DatatableParser.
type BaseDatatableListener struct{}

var _ DatatableListener = &BaseDatatableListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseDatatableListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseDatatableListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseDatatableListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseDatatableListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterDefinition is called when production definition is entered.
func (s *BaseDatatableListener) EnterDefinition(ctx *DefinitionContext) {}

// ExitDefinition is called when production definition is exited.
func (s *BaseDatatableListener) ExitDefinition(ctx *DefinitionContext) {}

// EnterAtom is called when production atom is entered.
func (s *BaseDatatableListener) EnterAtom(ctx *AtomContext) {}

// ExitAtom is called when production atom is exited.
func (s *BaseDatatableListener) ExitAtom(ctx *AtomContext) {}

// EnterName is called when production name is entered.
func (s *BaseDatatableListener) EnterName(ctx *NameContext) {}

// ExitName is called when production name is exited.
func (s *BaseDatatableListener) ExitName(ctx *NameContext) {}

// EnterTypeDefinition is called when production typeDefinition is entered.
func (s *BaseDatatableListener) EnterTypeDefinition(ctx *TypeDefinitionContext) {}

// ExitTypeDefinition is called when production typeDefinition is exited.
func (s *BaseDatatableListener) ExitTypeDefinition(ctx *TypeDefinitionContext) {}

// EnterFieldType is called when production fieldType is entered.
func (s *BaseDatatableListener) EnterFieldType(ctx *FieldTypeContext) {}

// ExitFieldType is called when production fieldType is exited.
func (s *BaseDatatableListener) ExitFieldType(ctx *FieldTypeContext) {}

// EnterDirect is called when production direct is entered.
func (s *BaseDatatableListener) EnterDirect(ctx *DirectContext) {}

// ExitDirect is called when production direct is exited.
func (s *BaseDatatableListener) ExitDirect(ctx *DirectContext) {}

// EnterSize is called when production size is entered.
func (s *BaseDatatableListener) EnterSize(ctx *SizeContext) {}

// ExitSize is called when production size is exited.
func (s *BaseDatatableListener) ExitSize(ctx *SizeContext) {}

// EnterChunk is called when production chunk is entered.
func (s *BaseDatatableListener) EnterChunk(ctx *ChunkContext) {}

// ExitChunk is called when production chunk is exited.
func (s *BaseDatatableListener) ExitChunk(ctx *ChunkContext) {}

// EnterTypeDeclaration is called when production typeDeclaration is entered.
func (s *BaseDatatableListener) EnterTypeDeclaration(ctx *TypeDeclarationContext) {}

// ExitTypeDeclaration is called when production typeDeclaration is exited.
func (s *BaseDatatableListener) ExitTypeDeclaration(ctx *TypeDeclarationContext) {}

// EnterDeclaration is called when production declaration is entered.
func (s *BaseDatatableListener) EnterDeclaration(ctx *DeclarationContext) {}

// ExitDeclaration is called when production declaration is exited.
func (s *BaseDatatableListener) ExitDeclaration(ctx *DeclarationContext) {}

// EnterBlock is called when production block is entered.
func (s *BaseDatatableListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BaseDatatableListener) ExitBlock(ctx *BlockContext) {}

// EnterBody is called when production body is entered.
func (s *BaseDatatableListener) EnterBody(ctx *BodyContext) {}

// ExitBody is called when production body is exited.
func (s *BaseDatatableListener) ExitBody(ctx *BodyContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseDatatableListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseDatatableListener) ExitExpression(ctx *ExpressionContext) {}

// EnterField is called when production field is entered.
func (s *BaseDatatableListener) EnterField(ctx *FieldContext) {}

// ExitField is called when production field is exited.
func (s *BaseDatatableListener) ExitField(ctx *FieldContext) {}
