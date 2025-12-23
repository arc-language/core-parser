parser grammar ArcParser;

options {
    tokenVocab=ArcLexer;
}

// Compilation Unit
compilationUnit
    : (importDecl | namespaceDecl | topLevelDecl)* EOF
    ;

// Import Declaration
importDecl
    : IMPORT IMPORT_PATH
    | IMPORT LPAREN importSpec* RPAREN
    ;

importSpec
    : IMPORT_PATH
    ;

// Namespace Declaration
namespaceDecl
    : NAMESPACE IDENTIFIER
    ;

// Top Level Declarations
topLevelDecl
    : functionDecl
    | structDecl
    | variableDecl
    | constDecl
    | externDecl
    ;

// Extern Declaration
externDecl
    : EXTERN NAMESPACE IDENTIFIER LBRACE externMember* RBRACE
    ;

externMember
    : externFunctionDecl
    ;

externFunctionDecl
    : FUNC IDENTIFIER (STRING_LITERAL)? LPAREN externParameterList? RPAREN type?
    ;

// Extern parameters are type-only (no names required)
externParameterList
    : type (COMMA type)* (COMMA ELLIPSIS)?
    | ELLIPSIS
    ;

// Function Declaration
functionDecl
    : FUNC IDENTIFIER LPAREN parameterList? RPAREN type? block
    ;

parameterList
    : parameter (COMMA parameter)* (COMMA ELLIPSIS)?
    | ELLIPSIS
    ;

parameter
    : SELF? IDENTIFIER COLON type
    ;

// Struct Declaration
structDecl
    : STRUCT IDENTIFIER LBRACE structField* RBRACE
    ;

structField
    : IDENTIFIER COLON type
    ;

// Variable Declaration
variableDecl
    : LET IDENTIFIER (COLON type)? ASSIGN expression
    ;

// Constant Declaration
constDecl
    : CONST IDENTIFIER (COLON type)? ASSIGN expression
    ;

// Type System
type
    : primitiveType
    | pointerType
    | referenceType
    | vectorType
    | mapType
    | IDENTIFIER  // Named types (structs, etc.)
    ;

primitiveType
    : INT8 | INT16 | INT32 | INT64
    | UINT8 | UINT16 | UINT32 | UINT64
    | USIZE | ISIZE
    | FLOAT32 | FLOAT64
    | BYTE | BOOL | RUNE
    | STRING
    | VOID
    ;

pointerType
    : STAR type
    ;

referenceType
    : AMP type
    ;

vectorType
    : VECTOR LT type GT
    ;

mapType
    : MAP LT type COMMA type GT
    ;

// Statements
block
    : LBRACE statement* RBRACE
    ;

statement
    : variableDecl
    | constDecl
    | assignmentStmt
    | expressionStmt
    | returnStmt
    | ifStmt
    | deferStmt
    | block
    ;

assignmentStmt
    : leftHandSide ASSIGN expression
    ;

leftHandSide
    : IDENTIFIER
    | STAR expression  // Pointer dereference
    | expression DOT IDENTIFIER  // Field access
    ;

expressionStmt
    : expression
    ;

returnStmt
    : RETURN expression?
    ;

ifStmt
    : IF expression block (ELSE IF expression block)* (ELSE block)?
    ;

deferStmt
    : DEFER expression
    ;

// Expressions (with proper precedence, lowest to highest)
expression
    : logicalOrExpression
    ;

logicalOrExpression
    : logicalAndExpression (OR logicalAndExpression)*
    ;

logicalAndExpression
    : equalityExpression (AND equalityExpression)*
    ;

equalityExpression
    : relationalExpression ((EQ | NE) relationalExpression)*
    ;

relationalExpression
    : additiveExpression ((LT | LE | GT | GE) additiveExpression)*
    ;

additiveExpression
    : multiplicativeExpression ((PLUS | MINUS) multiplicativeExpression)*
    ;

multiplicativeExpression
    : unaryExpression ((STAR | SLASH | PERCENT) unaryExpression)*
    ;

unaryExpression
    : (MINUS | NOT | STAR | AMP) unaryExpression
    | postfixExpression
    ;

postfixExpression
    : primaryExpression (postfixOp)*
    ;

postfixOp
    : DOT IDENTIFIER
    | DOT IDENTIFIER LPAREN argumentList? RPAREN
    | LPAREN argumentList? RPAREN
    ;

primaryExpression
    : literal
    | IDENTIFIER
    | LPAREN expression RPAREN
    | castExpression
    | allocaExpression
    | structLiteral
    ;

literal
    : INTEGER_LITERAL
    | FLOAT_LITERAL
    | STRING_LITERAL
    | CHAR_LITERAL
    | BOOLEAN_LITERAL
    | vectorLiteral
    | mapLiteral
    ;

vectorLiteral
    : LBRACE (expression (COMMA expression)*)? RBRACE
    ;

mapLiteral
    : LBRACE (mapEntry (COMMA mapEntry)*)? RBRACE
    ;

mapEntry
    : expression COLON expression
    ;

structLiteral
    : IDENTIFIER LBRACE (fieldInit (COMMA fieldInit)*)? RBRACE
    ;

fieldInit
    : IDENTIFIER COLON expression
    ;

argumentList
    : expression (COMMA expression)*
    ;

castExpression
    : CAST LT type GT LPAREN expression RPAREN
    ;

allocaExpression
    : ALLOCA LPAREN type (COMMA expression)? RPAREN
    ;