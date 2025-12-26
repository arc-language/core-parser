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
    : IMPORT STRING_LITERAL
    | IMPORT LPAREN importSpec* RPAREN
    ;

importSpec
    : STRING_LITERAL
    ;

// Namespace Declaration
namespaceDecl
    : NAMESPACE IDENTIFIER
    ;

// Top Level Declarations
topLevelDecl
    : functionDecl
    | structDecl
    | classDecl
    | methodDecl
    | deinitDecl
    | variableDecl
    | constDecl
    | externDecl
    ;

// Extern Declaration
externDecl
    : EXTERN IDENTIFIER LBRACE externMember* RBRACE
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

// Struct Declaration (can contain inline methods)
structDecl
    : STRUCT IDENTIFIER LBRACE structMember* RBRACE
    ;

structMember
    : structField
    | functionDecl
    ;

structField
    : IDENTIFIER COLON type
    ;

// Class Declaration (can contain inline methods and deinit)
classDecl
    : CLASS IDENTIFIER LBRACE classMember* RBRACE
    ;

classMember
    : classField
    | functionDecl
    | deinitDecl
    ;

classField
    : IDENTIFIER COLON type
    ;

// Flat Method Declaration (outside struct/class body)
methodDecl
    : FUNC IDENTIFIER LPAREN SELF IDENTIFIER COLON type (COMMA parameter)* RPAREN type? block
    ;

// Deinit Declaration (can be inline or flat)
deinitDecl
    : DEINIT LPAREN SELF IDENTIFIER COLON type RPAREN block
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
    | IDENTIFIER
    ;

primitiveType
    : INT8 | INT16 | INT32 | INT64
    | UINT8 | UINT16 | UINT32 | UINT64
    | USIZE | ISIZE
    | FLOAT32 | FLOAT64
    | BYTE | BOOL | CHAR
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
    | forStmt
    | breakStmt
    | continueStmt
    | deferStmt
    | block
    ;

assignmentStmt
    : leftHandSide (ASSIGN | PLUS_ASSIGN | MINUS_ASSIGN) expression
    ;

leftHandSide
    : IDENTIFIER
    | STAR expression
    | expression DOT IDENTIFIER
    ;

expressionStmt
    : expression
    ;

returnStmt
    : RETURN expression?
    ;

// Control Flow
ifStmt
    : IF expression block (ELSE IF expression block)* (ELSE block)?
    ;

// Modern For Loop
// 1. Infinite: for { }
// 2. Condition: for x < 10 { }
// 3. Clause: for let i=0; i<10; i=i+1 { }
// 4. For-in: for item in collection { }
// 5. For-in (map): for key, value in map { }
// 6. Range: for i in 0..10 { }
forStmt
    : FOR block
    | FOR expression block
    | FOR (variableDecl | assignmentStmt)? SEMICOLON expression? SEMICOLON (assignmentStmt | expression)? block
    | FOR IDENTIFIER IN expression block
    | FOR IDENTIFIER COMMA IDENTIFIER IN expression block
    ;

breakStmt
    : BREAK
    ;

continueStmt
    : CONTINUE
    ;

deferStmt
    : DEFER (assignmentStmt | expression)
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
    : rangeExpression ((LT | LE | GT | GE) rangeExpression)*
    ;

rangeExpression
    : additiveExpression (RANGE additiveExpression)?
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