// Code generated from D:/Projects/datatable/grammar/Datatable.g4 by ANTLR 4.9. DO NOT EDIT.

package parser // Datatable

import "github.com/antlr/antlr4/runtime/Go/antlr"

// DatatableListener is a complete listener for a parse tree produced by DatatableParser.
type DatatableListener interface {
	antlr.ParseTreeListener

	// EnterDefinition is called when entering the definition production.
	EnterDefinition(c *DefinitionContext)

	// EnterAtom is called when entering the atom production.
	EnterAtom(c *AtomContext)

	// EnterName is called when entering the name production.
	EnterName(c *NameContext)

	// EnterTypeDefinition is called when entering the typeDefinition production.
	EnterTypeDefinition(c *TypeDefinitionContext)

	// EnterFieldType is called when entering the fieldType production.
	EnterFieldType(c *FieldTypeContext)

	// EnterDirect is called when entering the direct production.
	EnterDirect(c *DirectContext)

	// EnterSize is called when entering the size production.
	EnterSize(c *SizeContext)

	// EnterChunk is called when entering the chunk production.
	EnterChunk(c *ChunkContext)

	// EnterTypeDeclaration is called when entering the typeDeclaration production.
	EnterTypeDeclaration(c *TypeDeclarationContext)

	// EnterDeclaration is called when entering the declaration production.
	EnterDeclaration(c *DeclarationContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// EnterBody is called when entering the body production.
	EnterBody(c *BodyContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterField is called when entering the field production.
	EnterField(c *FieldContext)

	// ExitDefinition is called when exiting the definition production.
	ExitDefinition(c *DefinitionContext)

	// ExitAtom is called when exiting the atom production.
	ExitAtom(c *AtomContext)

	// ExitName is called when exiting the name production.
	ExitName(c *NameContext)

	// ExitTypeDefinition is called when exiting the typeDefinition production.
	ExitTypeDefinition(c *TypeDefinitionContext)

	// ExitFieldType is called when exiting the fieldType production.
	ExitFieldType(c *FieldTypeContext)

	// ExitDirect is called when exiting the direct production.
	ExitDirect(c *DirectContext)

	// ExitSize is called when exiting the size production.
	ExitSize(c *SizeContext)

	// ExitChunk is called when exiting the chunk production.
	ExitChunk(c *ChunkContext)

	// ExitTypeDeclaration is called when exiting the typeDeclaration production.
	ExitTypeDeclaration(c *TypeDeclarationContext)

	// ExitDeclaration is called when exiting the declaration production.
	ExitDeclaration(c *DeclarationContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitBody is called when exiting the body production.
	ExitBody(c *BodyContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitField is called when exiting the field production.
	ExitField(c *FieldContext)
}
