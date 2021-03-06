grammar Datatable;

/*
    Keywords
*/
TYPE : 'Type' ;
GROUP : 'Group' ;
META : 'Meta' ;

/*
    Definition
*/
definition
    : GROUP
    | META
    ;

/*
    Name
*/
NAME
    : [a-zA-Z][a-zA-Z0-9]* ;

STRING
    : '"' ~('"')* '"'
    ;

INTEGER
    : [0-9]+
    ;

BOOLEAN
    : 'true'
    | 'false'
    ;

atom
    : STRING
    | INTEGER
    | BOOLEAN
    ;

Digit
    : [0-9]
    ;

name
    : ('@')? NAME ;

typeDefinition
    : (size '-')? name ('-' size)? ('?' atom)? ;

fieldType
    : typeDefinition
    | direct
    ;

direct
    : STRING ;

size
    : INTEGER
    | NAME
    | '?'
    ;

/*
    Start
*/
chunk
    : typeDeclaration* EOF ;

typeDeclaration
    : name ':' 'Type' body ;

declaration
    : name ':' definition body ;

block
    : expression* ;

body
    : '{' block '}' ;

/*
    Expressions
*/
expression
    : declaration
    | field
    ;

field
    : name ':' fieldType ;

/*
    Whitespace
*/
WS
    : [ \t\r\n]+ -> skip
    ;

/*
    Comment
*/
BlockComment
    :   '/*' .*? '*/'
        -> skip
    ;

LineComment
    :   '//' ~[\r\n]*
        -> skip
    ;