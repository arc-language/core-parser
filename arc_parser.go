// Code generated from ArcParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // ArcParser

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type ArcParser struct {
	*antlr.BaseParser
}

var ArcParserParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func arcparserParserInit() {
	staticData := &ArcParserParserStaticData
	staticData.LiteralNames = []string{
		"", "'import'", "'namespace'", "'let'", "'const'", "'func'", "'struct'",
		"'return'", "'if'", "'else'", "'for'", "'break'", "'continue'", "'defer'",
		"'extern'", "'self'", "'int8'", "'int16'", "'int32'", "'int64'", "'uint8'",
		"'uint16'", "'uint32'", "'uint64'", "'usize'", "'isize'", "'float32'",
		"'float64'", "'byte'", "'bool'", "'char'", "'string'", "'void'", "'vector'",
		"'map'", "'alloca'", "'cast'", "'+'", "'-'", "'*'", "'/'", "'%'", "'=='",
		"'!='", "'<'", "'<='", "'>'", "'>='", "'&&'", "'||'", "'!'", "'&'",
		"'@'", "'='", "'->'", "'('", "')'", "'{'", "'}'", "'['", "']'", "','",
		"':'", "';'", "'.'", "'...'",
	}
	staticData.SymbolicNames = []string{
		"", "IMPORT", "NAMESPACE", "LET", "CONST", "FUNC", "STRUCT", "RETURN",
		"IF", "ELSE", "FOR", "BREAK", "CONTINUE", "DEFER", "EXTERN", "SELF",
		"INT8", "INT16", "INT32", "INT64", "UINT8", "UINT16", "UINT32", "UINT64",
		"USIZE", "ISIZE", "FLOAT32", "FLOAT64", "BYTE", "BOOL", "CHAR", "STRING",
		"VOID", "VECTOR", "MAP", "ALLOCA", "CAST", "PLUS", "MINUS", "STAR",
		"SLASH", "PERCENT", "EQ", "NE", "LT", "LE", "GT", "GE", "AND", "OR",
		"NOT", "AMP", "AT", "ASSIGN", "ARROW", "LPAREN", "RPAREN", "LBRACE",
		"RBRACE", "LBRACKET", "RBRACKET", "COMMA", "COLON", "SEMICOLON", "DOT",
		"ELLIPSIS", "BOOLEAN_LITERAL", "INTEGER_LITERAL", "FLOAT_LITERAL", "STRING_LITERAL",
		"CHAR_LITERAL", "IDENTIFIER", "WS", "LINE_COMMENT", "BLOCK_COMMENT",
	}
	staticData.RuleNames = []string{
		"compilationUnit", "importDecl", "importSpec", "namespaceDecl", "topLevelDecl",
		"externDecl", "externMember", "externFunctionDecl", "externParameterList",
		"functionDecl", "parameterList", "parameter", "structDecl", "structField",
		"variableDecl", "constDecl", "type", "primitiveType", "pointerType",
		"referenceType", "vectorType", "mapType", "block", "statement", "assignmentStmt",
		"leftHandSide", "expressionStmt", "returnStmt", "ifStmt", "forStmt",
		"breakStmt", "continueStmt", "deferStmt", "expression", "logicalOrExpression",
		"logicalAndExpression", "equalityExpression", "relationalExpression",
		"additiveExpression", "multiplicativeExpression", "unaryExpression",
		"postfixExpression", "postfixOp", "primaryExpression", "literal", "vectorLiteral",
		"mapLiteral", "mapEntry", "structLiteral", "fieldInit", "argumentList",
		"castExpression", "allocaExpression",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 74, 536, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 7, 36,
		2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7, 41, 2,
		42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46, 2, 47,
		7, 47, 2, 48, 7, 48, 2, 49, 7, 49, 2, 50, 7, 50, 2, 51, 7, 51, 2, 52, 7,
		52, 1, 0, 1, 0, 1, 0, 5, 0, 110, 8, 0, 10, 0, 12, 0, 113, 9, 0, 1, 0, 1,
		0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 5, 1, 122, 8, 1, 10, 1, 12, 1, 125, 9,
		1, 1, 1, 3, 1, 128, 8, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1,
		4, 1, 4, 1, 4, 3, 4, 140, 8, 4, 1, 5, 1, 5, 1, 5, 1, 5, 5, 5, 146, 8, 5,
		10, 5, 12, 5, 149, 9, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7,
		3, 7, 159, 8, 7, 1, 7, 1, 7, 3, 7, 163, 8, 7, 1, 8, 1, 8, 1, 8, 5, 8, 168,
		8, 8, 10, 8, 12, 8, 171, 9, 8, 1, 8, 1, 8, 3, 8, 175, 8, 8, 1, 8, 3, 8,
		178, 8, 8, 1, 9, 1, 9, 1, 9, 1, 9, 3, 9, 184, 8, 9, 1, 9, 1, 9, 3, 9, 188,
		8, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 5, 10, 195, 8, 10, 10, 10, 12, 10,
		198, 9, 10, 1, 10, 1, 10, 3, 10, 202, 8, 10, 1, 10, 3, 10, 205, 8, 10,
		1, 11, 3, 11, 208, 8, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1,
		12, 1, 12, 5, 12, 218, 8, 12, 10, 12, 12, 12, 221, 9, 12, 1, 12, 1, 12,
		1, 13, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 14, 3, 14, 233, 8,
		14, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 15, 3, 15, 242, 8, 15,
		1, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 3, 16, 253,
		8, 16, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 19, 1, 19, 1, 19, 1, 20, 1,
		20, 1, 20, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21,
		1, 22, 1, 22, 5, 22, 277, 8, 22, 10, 22, 12, 22, 280, 9, 22, 1, 22, 1,
		22, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23,
		1, 23, 3, 23, 295, 8, 23, 1, 24, 1, 24, 1, 24, 1, 24, 1, 25, 1, 25, 1,
		25, 1, 25, 1, 25, 1, 25, 1, 25, 3, 25, 308, 8, 25, 1, 26, 1, 26, 1, 27,
		1, 27, 3, 27, 314, 8, 27, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1,
		28, 1, 28, 5, 28, 324, 8, 28, 10, 28, 12, 28, 327, 9, 28, 1, 28, 1, 28,
		3, 28, 331, 8, 28, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1,
		29, 1, 29, 3, 29, 342, 8, 29, 1, 29, 1, 29, 3, 29, 346, 8, 29, 1, 29, 1,
		29, 1, 29, 3, 29, 351, 8, 29, 1, 29, 3, 29, 354, 8, 29, 1, 30, 1, 30, 1,
		31, 1, 31, 1, 32, 1, 32, 1, 32, 3, 32, 363, 8, 32, 1, 33, 1, 33, 1, 34,
		1, 34, 1, 34, 5, 34, 370, 8, 34, 10, 34, 12, 34, 373, 9, 34, 1, 35, 1,
		35, 1, 35, 5, 35, 378, 8, 35, 10, 35, 12, 35, 381, 9, 35, 1, 36, 1, 36,
		1, 36, 5, 36, 386, 8, 36, 10, 36, 12, 36, 389, 9, 36, 1, 37, 1, 37, 1,
		37, 5, 37, 394, 8, 37, 10, 37, 12, 37, 397, 9, 37, 1, 38, 1, 38, 1, 38,
		5, 38, 402, 8, 38, 10, 38, 12, 38, 405, 9, 38, 1, 39, 1, 39, 1, 39, 5,
		39, 410, 8, 39, 10, 39, 12, 39, 413, 9, 39, 1, 40, 1, 40, 1, 40, 3, 40,
		418, 8, 40, 1, 41, 1, 41, 5, 41, 422, 8, 41, 10, 41, 12, 41, 425, 9, 41,
		1, 42, 1, 42, 1, 42, 1, 42, 1, 42, 1, 42, 3, 42, 433, 8, 42, 1, 42, 1,
		42, 1, 42, 3, 42, 438, 8, 42, 1, 42, 3, 42, 441, 8, 42, 1, 43, 1, 43, 1,
		43, 1, 43, 1, 43, 1, 43, 1, 43, 1, 43, 1, 43, 3, 43, 452, 8, 43, 1, 44,
		1, 44, 1, 44, 1, 44, 1, 44, 1, 44, 1, 44, 3, 44, 461, 8, 44, 1, 45, 1,
		45, 1, 45, 1, 45, 5, 45, 467, 8, 45, 10, 45, 12, 45, 470, 9, 45, 3, 45,
		472, 8, 45, 1, 45, 1, 45, 1, 46, 1, 46, 1, 46, 1, 46, 5, 46, 480, 8, 46,
		10, 46, 12, 46, 483, 9, 46, 3, 46, 485, 8, 46, 1, 46, 1, 46, 1, 47, 1,
		47, 1, 47, 1, 47, 1, 48, 1, 48, 1, 48, 1, 48, 1, 48, 5, 48, 498, 8, 48,
		10, 48, 12, 48, 501, 9, 48, 3, 48, 503, 8, 48, 1, 48, 1, 48, 1, 49, 1,
		49, 1, 49, 1, 49, 1, 50, 1, 50, 1, 50, 5, 50, 514, 8, 50, 10, 50, 12, 50,
		517, 9, 50, 1, 51, 1, 51, 1, 51, 1, 51, 1, 51, 1, 51, 1, 51, 1, 51, 1,
		52, 1, 52, 1, 52, 1, 52, 1, 52, 3, 52, 532, 8, 52, 1, 52, 1, 52, 1, 52,
		0, 0, 53, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32,
		34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68,
		70, 72, 74, 76, 78, 80, 82, 84, 86, 88, 90, 92, 94, 96, 98, 100, 102, 104,
		0, 6, 1, 0, 16, 32, 1, 0, 42, 43, 1, 0, 44, 47, 1, 0, 37, 38, 1, 0, 39,
		41, 2, 0, 38, 39, 50, 51, 566, 0, 111, 1, 0, 0, 0, 2, 127, 1, 0, 0, 0,
		4, 129, 1, 0, 0, 0, 6, 131, 1, 0, 0, 0, 8, 139, 1, 0, 0, 0, 10, 141, 1,
		0, 0, 0, 12, 152, 1, 0, 0, 0, 14, 154, 1, 0, 0, 0, 16, 177, 1, 0, 0, 0,
		18, 179, 1, 0, 0, 0, 20, 204, 1, 0, 0, 0, 22, 207, 1, 0, 0, 0, 24, 213,
		1, 0, 0, 0, 26, 224, 1, 0, 0, 0, 28, 228, 1, 0, 0, 0, 30, 237, 1, 0, 0,
		0, 32, 252, 1, 0, 0, 0, 34, 254, 1, 0, 0, 0, 36, 256, 1, 0, 0, 0, 38, 259,
		1, 0, 0, 0, 40, 262, 1, 0, 0, 0, 42, 267, 1, 0, 0, 0, 44, 274, 1, 0, 0,
		0, 46, 294, 1, 0, 0, 0, 48, 296, 1, 0, 0, 0, 50, 307, 1, 0, 0, 0, 52, 309,
		1, 0, 0, 0, 54, 311, 1, 0, 0, 0, 56, 315, 1, 0, 0, 0, 58, 353, 1, 0, 0,
		0, 60, 355, 1, 0, 0, 0, 62, 357, 1, 0, 0, 0, 64, 359, 1, 0, 0, 0, 66, 364,
		1, 0, 0, 0, 68, 366, 1, 0, 0, 0, 70, 374, 1, 0, 0, 0, 72, 382, 1, 0, 0,
		0, 74, 390, 1, 0, 0, 0, 76, 398, 1, 0, 0, 0, 78, 406, 1, 0, 0, 0, 80, 417,
		1, 0, 0, 0, 82, 419, 1, 0, 0, 0, 84, 440, 1, 0, 0, 0, 86, 451, 1, 0, 0,
		0, 88, 460, 1, 0, 0, 0, 90, 462, 1, 0, 0, 0, 92, 475, 1, 0, 0, 0, 94, 488,
		1, 0, 0, 0, 96, 492, 1, 0, 0, 0, 98, 506, 1, 0, 0, 0, 100, 510, 1, 0, 0,
		0, 102, 518, 1, 0, 0, 0, 104, 526, 1, 0, 0, 0, 106, 110, 3, 2, 1, 0, 107,
		110, 3, 6, 3, 0, 108, 110, 3, 8, 4, 0, 109, 106, 1, 0, 0, 0, 109, 107,
		1, 0, 0, 0, 109, 108, 1, 0, 0, 0, 110, 113, 1, 0, 0, 0, 111, 109, 1, 0,
		0, 0, 111, 112, 1, 0, 0, 0, 112, 114, 1, 0, 0, 0, 113, 111, 1, 0, 0, 0,
		114, 115, 5, 0, 0, 1, 115, 1, 1, 0, 0, 0, 116, 117, 5, 1, 0, 0, 117, 128,
		5, 69, 0, 0, 118, 119, 5, 1, 0, 0, 119, 123, 5, 55, 0, 0, 120, 122, 3,
		4, 2, 0, 121, 120, 1, 0, 0, 0, 122, 125, 1, 0, 0, 0, 123, 121, 1, 0, 0,
		0, 123, 124, 1, 0, 0, 0, 124, 126, 1, 0, 0, 0, 125, 123, 1, 0, 0, 0, 126,
		128, 5, 56, 0, 0, 127, 116, 1, 0, 0, 0, 127, 118, 1, 0, 0, 0, 128, 3, 1,
		0, 0, 0, 129, 130, 5, 69, 0, 0, 130, 5, 1, 0, 0, 0, 131, 132, 5, 2, 0,
		0, 132, 133, 5, 71, 0, 0, 133, 7, 1, 0, 0, 0, 134, 140, 3, 18, 9, 0, 135,
		140, 3, 24, 12, 0, 136, 140, 3, 28, 14, 0, 137, 140, 3, 30, 15, 0, 138,
		140, 3, 10, 5, 0, 139, 134, 1, 0, 0, 0, 139, 135, 1, 0, 0, 0, 139, 136,
		1, 0, 0, 0, 139, 137, 1, 0, 0, 0, 139, 138, 1, 0, 0, 0, 140, 9, 1, 0, 0,
		0, 141, 142, 5, 14, 0, 0, 142, 143, 5, 71, 0, 0, 143, 147, 5, 57, 0, 0,
		144, 146, 3, 12, 6, 0, 145, 144, 1, 0, 0, 0, 146, 149, 1, 0, 0, 0, 147,
		145, 1, 0, 0, 0, 147, 148, 1, 0, 0, 0, 148, 150, 1, 0, 0, 0, 149, 147,
		1, 0, 0, 0, 150, 151, 5, 58, 0, 0, 151, 11, 1, 0, 0, 0, 152, 153, 3, 14,
		7, 0, 153, 13, 1, 0, 0, 0, 154, 155, 5, 5, 0, 0, 155, 156, 5, 71, 0, 0,
		156, 158, 5, 55, 0, 0, 157, 159, 3, 16, 8, 0, 158, 157, 1, 0, 0, 0, 158,
		159, 1, 0, 0, 0, 159, 160, 1, 0, 0, 0, 160, 162, 5, 56, 0, 0, 161, 163,
		3, 32, 16, 0, 162, 161, 1, 0, 0, 0, 162, 163, 1, 0, 0, 0, 163, 15, 1, 0,
		0, 0, 164, 169, 3, 32, 16, 0, 165, 166, 5, 61, 0, 0, 166, 168, 3, 32, 16,
		0, 167, 165, 1, 0, 0, 0, 168, 171, 1, 0, 0, 0, 169, 167, 1, 0, 0, 0, 169,
		170, 1, 0, 0, 0, 170, 174, 1, 0, 0, 0, 171, 169, 1, 0, 0, 0, 172, 173,
		5, 61, 0, 0, 173, 175, 5, 65, 0, 0, 174, 172, 1, 0, 0, 0, 174, 175, 1,
		0, 0, 0, 175, 178, 1, 0, 0, 0, 176, 178, 5, 65, 0, 0, 177, 164, 1, 0, 0,
		0, 177, 176, 1, 0, 0, 0, 178, 17, 1, 0, 0, 0, 179, 180, 5, 5, 0, 0, 180,
		181, 5, 71, 0, 0, 181, 183, 5, 55, 0, 0, 182, 184, 3, 20, 10, 0, 183, 182,
		1, 0, 0, 0, 183, 184, 1, 0, 0, 0, 184, 185, 1, 0, 0, 0, 185, 187, 5, 56,
		0, 0, 186, 188, 3, 32, 16, 0, 187, 186, 1, 0, 0, 0, 187, 188, 1, 0, 0,
		0, 188, 189, 1, 0, 0, 0, 189, 190, 3, 44, 22, 0, 190, 19, 1, 0, 0, 0, 191,
		196, 3, 22, 11, 0, 192, 193, 5, 61, 0, 0, 193, 195, 3, 22, 11, 0, 194,
		192, 1, 0, 0, 0, 195, 198, 1, 0, 0, 0, 196, 194, 1, 0, 0, 0, 196, 197,
		1, 0, 0, 0, 197, 201, 1, 0, 0, 0, 198, 196, 1, 0, 0, 0, 199, 200, 5, 61,
		0, 0, 200, 202, 5, 65, 0, 0, 201, 199, 1, 0, 0, 0, 201, 202, 1, 0, 0, 0,
		202, 205, 1, 0, 0, 0, 203, 205, 5, 65, 0, 0, 204, 191, 1, 0, 0, 0, 204,
		203, 1, 0, 0, 0, 205, 21, 1, 0, 0, 0, 206, 208, 5, 15, 0, 0, 207, 206,
		1, 0, 0, 0, 207, 208, 1, 0, 0, 0, 208, 209, 1, 0, 0, 0, 209, 210, 5, 71,
		0, 0, 210, 211, 5, 62, 0, 0, 211, 212, 3, 32, 16, 0, 212, 23, 1, 0, 0,
		0, 213, 214, 5, 6, 0, 0, 214, 215, 5, 71, 0, 0, 215, 219, 5, 57, 0, 0,
		216, 218, 3, 26, 13, 0, 217, 216, 1, 0, 0, 0, 218, 221, 1, 0, 0, 0, 219,
		217, 1, 0, 0, 0, 219, 220, 1, 0, 0, 0, 220, 222, 1, 0, 0, 0, 221, 219,
		1, 0, 0, 0, 222, 223, 5, 58, 0, 0, 223, 25, 1, 0, 0, 0, 224, 225, 5, 71,
		0, 0, 225, 226, 5, 62, 0, 0, 226, 227, 3, 32, 16, 0, 227, 27, 1, 0, 0,
		0, 228, 229, 5, 3, 0, 0, 229, 232, 5, 71, 0, 0, 230, 231, 5, 62, 0, 0,
		231, 233, 3, 32, 16, 0, 232, 230, 1, 0, 0, 0, 232, 233, 1, 0, 0, 0, 233,
		234, 1, 0, 0, 0, 234, 235, 5, 53, 0, 0, 235, 236, 3, 66, 33, 0, 236, 29,
		1, 0, 0, 0, 237, 238, 5, 4, 0, 0, 238, 241, 5, 71, 0, 0, 239, 240, 5, 62,
		0, 0, 240, 242, 3, 32, 16, 0, 241, 239, 1, 0, 0, 0, 241, 242, 1, 0, 0,
		0, 242, 243, 1, 0, 0, 0, 243, 244, 5, 53, 0, 0, 244, 245, 3, 66, 33, 0,
		245, 31, 1, 0, 0, 0, 246, 253, 3, 34, 17, 0, 247, 253, 3, 36, 18, 0, 248,
		253, 3, 38, 19, 0, 249, 253, 3, 40, 20, 0, 250, 253, 3, 42, 21, 0, 251,
		253, 5, 71, 0, 0, 252, 246, 1, 0, 0, 0, 252, 247, 1, 0, 0, 0, 252, 248,
		1, 0, 0, 0, 252, 249, 1, 0, 0, 0, 252, 250, 1, 0, 0, 0, 252, 251, 1, 0,
		0, 0, 253, 33, 1, 0, 0, 0, 254, 255, 7, 0, 0, 0, 255, 35, 1, 0, 0, 0, 256,
		257, 5, 39, 0, 0, 257, 258, 3, 32, 16, 0, 258, 37, 1, 0, 0, 0, 259, 260,
		5, 51, 0, 0, 260, 261, 3, 32, 16, 0, 261, 39, 1, 0, 0, 0, 262, 263, 5,
		33, 0, 0, 263, 264, 5, 44, 0, 0, 264, 265, 3, 32, 16, 0, 265, 266, 5, 46,
		0, 0, 266, 41, 1, 0, 0, 0, 267, 268, 5, 34, 0, 0, 268, 269, 5, 44, 0, 0,
		269, 270, 3, 32, 16, 0, 270, 271, 5, 61, 0, 0, 271, 272, 3, 32, 16, 0,
		272, 273, 5, 46, 0, 0, 273, 43, 1, 0, 0, 0, 274, 278, 5, 57, 0, 0, 275,
		277, 3, 46, 23, 0, 276, 275, 1, 0, 0, 0, 277, 280, 1, 0, 0, 0, 278, 276,
		1, 0, 0, 0, 278, 279, 1, 0, 0, 0, 279, 281, 1, 0, 0, 0, 280, 278, 1, 0,
		0, 0, 281, 282, 5, 58, 0, 0, 282, 45, 1, 0, 0, 0, 283, 295, 3, 28, 14,
		0, 284, 295, 3, 30, 15, 0, 285, 295, 3, 48, 24, 0, 286, 295, 3, 52, 26,
		0, 287, 295, 3, 54, 27, 0, 288, 295, 3, 56, 28, 0, 289, 295, 3, 58, 29,
		0, 290, 295, 3, 60, 30, 0, 291, 295, 3, 62, 31, 0, 292, 295, 3, 64, 32,
		0, 293, 295, 3, 44, 22, 0, 294, 283, 1, 0, 0, 0, 294, 284, 1, 0, 0, 0,
		294, 285, 1, 0, 0, 0, 294, 286, 1, 0, 0, 0, 294, 287, 1, 0, 0, 0, 294,
		288, 1, 0, 0, 0, 294, 289, 1, 0, 0, 0, 294, 290, 1, 0, 0, 0, 294, 291,
		1, 0, 0, 0, 294, 292, 1, 0, 0, 0, 294, 293, 1, 0, 0, 0, 295, 47, 1, 0,
		0, 0, 296, 297, 3, 50, 25, 0, 297, 298, 5, 53, 0, 0, 298, 299, 3, 66, 33,
		0, 299, 49, 1, 0, 0, 0, 300, 308, 5, 71, 0, 0, 301, 302, 5, 39, 0, 0, 302,
		308, 3, 66, 33, 0, 303, 304, 3, 66, 33, 0, 304, 305, 5, 64, 0, 0, 305,
		306, 5, 71, 0, 0, 306, 308, 1, 0, 0, 0, 307, 300, 1, 0, 0, 0, 307, 301,
		1, 0, 0, 0, 307, 303, 1, 0, 0, 0, 308, 51, 1, 0, 0, 0, 309, 310, 3, 66,
		33, 0, 310, 53, 1, 0, 0, 0, 311, 313, 5, 7, 0, 0, 312, 314, 3, 66, 33,
		0, 313, 312, 1, 0, 0, 0, 313, 314, 1, 0, 0, 0, 314, 55, 1, 0, 0, 0, 315,
		316, 5, 8, 0, 0, 316, 317, 3, 66, 33, 0, 317, 325, 3, 44, 22, 0, 318, 319,
		5, 9, 0, 0, 319, 320, 5, 8, 0, 0, 320, 321, 3, 66, 33, 0, 321, 322, 3,
		44, 22, 0, 322, 324, 1, 0, 0, 0, 323, 318, 1, 0, 0, 0, 324, 327, 1, 0,
		0, 0, 325, 323, 1, 0, 0, 0, 325, 326, 1, 0, 0, 0, 326, 330, 1, 0, 0, 0,
		327, 325, 1, 0, 0, 0, 328, 329, 5, 9, 0, 0, 329, 331, 3, 44, 22, 0, 330,
		328, 1, 0, 0, 0, 330, 331, 1, 0, 0, 0, 331, 57, 1, 0, 0, 0, 332, 333, 5,
		10, 0, 0, 333, 354, 3, 44, 22, 0, 334, 335, 5, 10, 0, 0, 335, 336, 3, 66,
		33, 0, 336, 337, 3, 44, 22, 0, 337, 354, 1, 0, 0, 0, 338, 341, 5, 10, 0,
		0, 339, 342, 3, 28, 14, 0, 340, 342, 3, 48, 24, 0, 341, 339, 1, 0, 0, 0,
		341, 340, 1, 0, 0, 0, 341, 342, 1, 0, 0, 0, 342, 343, 1, 0, 0, 0, 343,
		345, 5, 63, 0, 0, 344, 346, 3, 66, 33, 0, 345, 344, 1, 0, 0, 0, 345, 346,
		1, 0, 0, 0, 346, 347, 1, 0, 0, 0, 347, 350, 5, 63, 0, 0, 348, 351, 3, 48,
		24, 0, 349, 351, 3, 66, 33, 0, 350, 348, 1, 0, 0, 0, 350, 349, 1, 0, 0,
		0, 350, 351, 1, 0, 0, 0, 351, 352, 1, 0, 0, 0, 352, 354, 3, 44, 22, 0,
		353, 332, 1, 0, 0, 0, 353, 334, 1, 0, 0, 0, 353, 338, 1, 0, 0, 0, 354,
		59, 1, 0, 0, 0, 355, 356, 5, 11, 0, 0, 356, 61, 1, 0, 0, 0, 357, 358, 5,
		12, 0, 0, 358, 63, 1, 0, 0, 0, 359, 362, 5, 13, 0, 0, 360, 363, 3, 48,
		24, 0, 361, 363, 3, 66, 33, 0, 362, 360, 1, 0, 0, 0, 362, 361, 1, 0, 0,
		0, 363, 65, 1, 0, 0, 0, 364, 365, 3, 68, 34, 0, 365, 67, 1, 0, 0, 0, 366,
		371, 3, 70, 35, 0, 367, 368, 5, 49, 0, 0, 368, 370, 3, 70, 35, 0, 369,
		367, 1, 0, 0, 0, 370, 373, 1, 0, 0, 0, 371, 369, 1, 0, 0, 0, 371, 372,
		1, 0, 0, 0, 372, 69, 1, 0, 0, 0, 373, 371, 1, 0, 0, 0, 374, 379, 3, 72,
		36, 0, 375, 376, 5, 48, 0, 0, 376, 378, 3, 72, 36, 0, 377, 375, 1, 0, 0,
		0, 378, 381, 1, 0, 0, 0, 379, 377, 1, 0, 0, 0, 379, 380, 1, 0, 0, 0, 380,
		71, 1, 0, 0, 0, 381, 379, 1, 0, 0, 0, 382, 387, 3, 74, 37, 0, 383, 384,
		7, 1, 0, 0, 384, 386, 3, 74, 37, 0, 385, 383, 1, 0, 0, 0, 386, 389, 1,
		0, 0, 0, 387, 385, 1, 0, 0, 0, 387, 388, 1, 0, 0, 0, 388, 73, 1, 0, 0,
		0, 389, 387, 1, 0, 0, 0, 390, 395, 3, 76, 38, 0, 391, 392, 7, 2, 0, 0,
		392, 394, 3, 76, 38, 0, 393, 391, 1, 0, 0, 0, 394, 397, 1, 0, 0, 0, 395,
		393, 1, 0, 0, 0, 395, 396, 1, 0, 0, 0, 396, 75, 1, 0, 0, 0, 397, 395, 1,
		0, 0, 0, 398, 403, 3, 78, 39, 0, 399, 400, 7, 3, 0, 0, 400, 402, 3, 78,
		39, 0, 401, 399, 1, 0, 0, 0, 402, 405, 1, 0, 0, 0, 403, 401, 1, 0, 0, 0,
		403, 404, 1, 0, 0, 0, 404, 77, 1, 0, 0, 0, 405, 403, 1, 0, 0, 0, 406, 411,
		3, 80, 40, 0, 407, 408, 7, 4, 0, 0, 408, 410, 3, 80, 40, 0, 409, 407, 1,
		0, 0, 0, 410, 413, 1, 0, 0, 0, 411, 409, 1, 0, 0, 0, 411, 412, 1, 0, 0,
		0, 412, 79, 1, 0, 0, 0, 413, 411, 1, 0, 0, 0, 414, 415, 7, 5, 0, 0, 415,
		418, 3, 80, 40, 0, 416, 418, 3, 82, 41, 0, 417, 414, 1, 0, 0, 0, 417, 416,
		1, 0, 0, 0, 418, 81, 1, 0, 0, 0, 419, 423, 3, 86, 43, 0, 420, 422, 3, 84,
		42, 0, 421, 420, 1, 0, 0, 0, 422, 425, 1, 0, 0, 0, 423, 421, 1, 0, 0, 0,
		423, 424, 1, 0, 0, 0, 424, 83, 1, 0, 0, 0, 425, 423, 1, 0, 0, 0, 426, 427,
		5, 64, 0, 0, 427, 441, 5, 71, 0, 0, 428, 429, 5, 64, 0, 0, 429, 430, 5,
		71, 0, 0, 430, 432, 5, 55, 0, 0, 431, 433, 3, 100, 50, 0, 432, 431, 1,
		0, 0, 0, 432, 433, 1, 0, 0, 0, 433, 434, 1, 0, 0, 0, 434, 441, 5, 56, 0,
		0, 435, 437, 5, 55, 0, 0, 436, 438, 3, 100, 50, 0, 437, 436, 1, 0, 0, 0,
		437, 438, 1, 0, 0, 0, 438, 439, 1, 0, 0, 0, 439, 441, 5, 56, 0, 0, 440,
		426, 1, 0, 0, 0, 440, 428, 1, 0, 0, 0, 440, 435, 1, 0, 0, 0, 441, 85, 1,
		0, 0, 0, 442, 452, 3, 88, 44, 0, 443, 452, 5, 71, 0, 0, 444, 445, 5, 55,
		0, 0, 445, 446, 3, 66, 33, 0, 446, 447, 5, 56, 0, 0, 447, 452, 1, 0, 0,
		0, 448, 452, 3, 102, 51, 0, 449, 452, 3, 104, 52, 0, 450, 452, 3, 96, 48,
		0, 451, 442, 1, 0, 0, 0, 451, 443, 1, 0, 0, 0, 451, 444, 1, 0, 0, 0, 451,
		448, 1, 0, 0, 0, 451, 449, 1, 0, 0, 0, 451, 450, 1, 0, 0, 0, 452, 87, 1,
		0, 0, 0, 453, 461, 5, 67, 0, 0, 454, 461, 5, 68, 0, 0, 455, 461, 5, 69,
		0, 0, 456, 461, 5, 70, 0, 0, 457, 461, 5, 66, 0, 0, 458, 461, 3, 90, 45,
		0, 459, 461, 3, 92, 46, 0, 460, 453, 1, 0, 0, 0, 460, 454, 1, 0, 0, 0,
		460, 455, 1, 0, 0, 0, 460, 456, 1, 0, 0, 0, 460, 457, 1, 0, 0, 0, 460,
		458, 1, 0, 0, 0, 460, 459, 1, 0, 0, 0, 461, 89, 1, 0, 0, 0, 462, 471, 5,
		57, 0, 0, 463, 468, 3, 66, 33, 0, 464, 465, 5, 61, 0, 0, 465, 467, 3, 66,
		33, 0, 466, 464, 1, 0, 0, 0, 467, 470, 1, 0, 0, 0, 468, 466, 1, 0, 0, 0,
		468, 469, 1, 0, 0, 0, 469, 472, 1, 0, 0, 0, 470, 468, 1, 0, 0, 0, 471,
		463, 1, 0, 0, 0, 471, 472, 1, 0, 0, 0, 472, 473, 1, 0, 0, 0, 473, 474,
		5, 58, 0, 0, 474, 91, 1, 0, 0, 0, 475, 484, 5, 57, 0, 0, 476, 481, 3, 94,
		47, 0, 477, 478, 5, 61, 0, 0, 478, 480, 3, 94, 47, 0, 479, 477, 1, 0, 0,
		0, 480, 483, 1, 0, 0, 0, 481, 479, 1, 0, 0, 0, 481, 482, 1, 0, 0, 0, 482,
		485, 1, 0, 0, 0, 483, 481, 1, 0, 0, 0, 484, 476, 1, 0, 0, 0, 484, 485,
		1, 0, 0, 0, 485, 486, 1, 0, 0, 0, 486, 487, 5, 58, 0, 0, 487, 93, 1, 0,
		0, 0, 488, 489, 3, 66, 33, 0, 489, 490, 5, 62, 0, 0, 490, 491, 3, 66, 33,
		0, 491, 95, 1, 0, 0, 0, 492, 493, 5, 71, 0, 0, 493, 502, 5, 57, 0, 0, 494,
		499, 3, 98, 49, 0, 495, 496, 5, 61, 0, 0, 496, 498, 3, 98, 49, 0, 497,
		495, 1, 0, 0, 0, 498, 501, 1, 0, 0, 0, 499, 497, 1, 0, 0, 0, 499, 500,
		1, 0, 0, 0, 500, 503, 1, 0, 0, 0, 501, 499, 1, 0, 0, 0, 502, 494, 1, 0,
		0, 0, 502, 503, 1, 0, 0, 0, 503, 504, 1, 0, 0, 0, 504, 505, 5, 58, 0, 0,
		505, 97, 1, 0, 0, 0, 506, 507, 5, 71, 0, 0, 507, 508, 5, 62, 0, 0, 508,
		509, 3, 66, 33, 0, 509, 99, 1, 0, 0, 0, 510, 515, 3, 66, 33, 0, 511, 512,
		5, 61, 0, 0, 512, 514, 3, 66, 33, 0, 513, 511, 1, 0, 0, 0, 514, 517, 1,
		0, 0, 0, 515, 513, 1, 0, 0, 0, 515, 516, 1, 0, 0, 0, 516, 101, 1, 0, 0,
		0, 517, 515, 1, 0, 0, 0, 518, 519, 5, 36, 0, 0, 519, 520, 5, 44, 0, 0,
		520, 521, 3, 32, 16, 0, 521, 522, 5, 46, 0, 0, 522, 523, 5, 55, 0, 0, 523,
		524, 3, 66, 33, 0, 524, 525, 5, 56, 0, 0, 525, 103, 1, 0, 0, 0, 526, 527,
		5, 35, 0, 0, 527, 528, 5, 55, 0, 0, 528, 531, 3, 32, 16, 0, 529, 530, 5,
		61, 0, 0, 530, 532, 3, 66, 33, 0, 531, 529, 1, 0, 0, 0, 531, 532, 1, 0,
		0, 0, 532, 533, 1, 0, 0, 0, 533, 534, 5, 56, 0, 0, 534, 105, 1, 0, 0, 0,
		53, 109, 111, 123, 127, 139, 147, 158, 162, 169, 174, 177, 183, 187, 196,
		201, 204, 207, 219, 232, 241, 252, 278, 294, 307, 313, 325, 330, 341, 345,
		350, 353, 362, 371, 379, 387, 395, 403, 411, 417, 423, 432, 437, 440, 451,
		460, 468, 471, 481, 484, 499, 502, 515, 531,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// ArcParserInit initializes any static state used to implement ArcParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewArcParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func ArcParserInit() {
	staticData := &ArcParserParserStaticData
	staticData.once.Do(arcparserParserInit)
}

// NewArcParser produces a new parser instance for the optional input antlr.TokenStream.
func NewArcParser(input antlr.TokenStream) *ArcParser {
	ArcParserInit()
	this := new(ArcParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &ArcParserParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "ArcParser.g4"

	return this
}

// ArcParser tokens.
const (
	ArcParserEOF             = antlr.TokenEOF
	ArcParserIMPORT          = 1
	ArcParserNAMESPACE       = 2
	ArcParserLET             = 3
	ArcParserCONST           = 4
	ArcParserFUNC            = 5
	ArcParserSTRUCT          = 6
	ArcParserRETURN          = 7
	ArcParserIF              = 8
	ArcParserELSE            = 9
	ArcParserFOR             = 10
	ArcParserBREAK           = 11
	ArcParserCONTINUE        = 12
	ArcParserDEFER           = 13
	ArcParserEXTERN          = 14
	ArcParserSELF            = 15
	ArcParserINT8            = 16
	ArcParserINT16           = 17
	ArcParserINT32           = 18
	ArcParserINT64           = 19
	ArcParserUINT8           = 20
	ArcParserUINT16          = 21
	ArcParserUINT32          = 22
	ArcParserUINT64          = 23
	ArcParserUSIZE           = 24
	ArcParserISIZE           = 25
	ArcParserFLOAT32         = 26
	ArcParserFLOAT64         = 27
	ArcParserBYTE            = 28
	ArcParserBOOL            = 29
	ArcParserCHAR            = 30
	ArcParserSTRING          = 31
	ArcParserVOID            = 32
	ArcParserVECTOR          = 33
	ArcParserMAP             = 34
	ArcParserALLOCA          = 35
	ArcParserCAST            = 36
	ArcParserPLUS            = 37
	ArcParserMINUS           = 38
	ArcParserSTAR            = 39
	ArcParserSLASH           = 40
	ArcParserPERCENT         = 41
	ArcParserEQ              = 42
	ArcParserNE              = 43
	ArcParserLT              = 44
	ArcParserLE              = 45
	ArcParserGT              = 46
	ArcParserGE              = 47
	ArcParserAND             = 48
	ArcParserOR              = 49
	ArcParserNOT             = 50
	ArcParserAMP             = 51
	ArcParserAT              = 52
	ArcParserASSIGN          = 53
	ArcParserARROW           = 54
	ArcParserLPAREN          = 55
	ArcParserRPAREN          = 56
	ArcParserLBRACE          = 57
	ArcParserRBRACE          = 58
	ArcParserLBRACKET        = 59
	ArcParserRBRACKET        = 60
	ArcParserCOMMA           = 61
	ArcParserCOLON           = 62
	ArcParserSEMICOLON       = 63
	ArcParserDOT             = 64
	ArcParserELLIPSIS        = 65
	ArcParserBOOLEAN_LITERAL = 66
	ArcParserINTEGER_LITERAL = 67
	ArcParserFLOAT_LITERAL   = 68
	ArcParserSTRING_LITERAL  = 69
	ArcParserCHAR_LITERAL    = 70
	ArcParserIDENTIFIER      = 71
	ArcParserWS              = 72
	ArcParserLINE_COMMENT    = 73
	ArcParserBLOCK_COMMENT   = 74
)

// ArcParser rules.
const (
	ArcParserRULE_compilationUnit          = 0
	ArcParserRULE_importDecl               = 1
	ArcParserRULE_importSpec               = 2
	ArcParserRULE_namespaceDecl            = 3
	ArcParserRULE_topLevelDecl             = 4
	ArcParserRULE_externDecl               = 5
	ArcParserRULE_externMember             = 6
	ArcParserRULE_externFunctionDecl       = 7
	ArcParserRULE_externParameterList      = 8
	ArcParserRULE_functionDecl             = 9
	ArcParserRULE_parameterList            = 10
	ArcParserRULE_parameter                = 11
	ArcParserRULE_structDecl               = 12
	ArcParserRULE_structField              = 13
	ArcParserRULE_variableDecl             = 14
	ArcParserRULE_constDecl                = 15
	ArcParserRULE_type                     = 16
	ArcParserRULE_primitiveType            = 17
	ArcParserRULE_pointerType              = 18
	ArcParserRULE_referenceType            = 19
	ArcParserRULE_vectorType               = 20
	ArcParserRULE_mapType                  = 21
	ArcParserRULE_block                    = 22
	ArcParserRULE_statement                = 23
	ArcParserRULE_assignmentStmt           = 24
	ArcParserRULE_leftHandSide             = 25
	ArcParserRULE_expressionStmt           = 26
	ArcParserRULE_returnStmt               = 27
	ArcParserRULE_ifStmt                   = 28
	ArcParserRULE_forStmt                  = 29
	ArcParserRULE_breakStmt                = 30
	ArcParserRULE_continueStmt             = 31
	ArcParserRULE_deferStmt                = 32
	ArcParserRULE_expression               = 33
	ArcParserRULE_logicalOrExpression      = 34
	ArcParserRULE_logicalAndExpression     = 35
	ArcParserRULE_equalityExpression       = 36
	ArcParserRULE_relationalExpression     = 37
	ArcParserRULE_additiveExpression       = 38
	ArcParserRULE_multiplicativeExpression = 39
	ArcParserRULE_unaryExpression          = 40
	ArcParserRULE_postfixExpression        = 41
	ArcParserRULE_postfixOp                = 42
	ArcParserRULE_primaryExpression        = 43
	ArcParserRULE_literal                  = 44
	ArcParserRULE_vectorLiteral            = 45
	ArcParserRULE_mapLiteral               = 46
	ArcParserRULE_mapEntry                 = 47
	ArcParserRULE_structLiteral            = 48
	ArcParserRULE_fieldInit                = 49
	ArcParserRULE_argumentList             = 50
	ArcParserRULE_castExpression           = 51
	ArcParserRULE_allocaExpression         = 52
)

// ICompilationUnitContext is an interface to support dynamic dispatch.
type ICompilationUnitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EOF() antlr.TerminalNode
	AllImportDecl() []IImportDeclContext
	ImportDecl(i int) IImportDeclContext
	AllNamespaceDecl() []INamespaceDeclContext
	NamespaceDecl(i int) INamespaceDeclContext
	AllTopLevelDecl() []ITopLevelDeclContext
	TopLevelDecl(i int) ITopLevelDeclContext

	// IsCompilationUnitContext differentiates from other interfaces.
	IsCompilationUnitContext()
}

type CompilationUnitContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCompilationUnitContext() *CompilationUnitContext {
	var p = new(CompilationUnitContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ArcParserRULE_compilationUnit
	return p
}

func InitEmptyCompilationUnitContext(p *CompilationUnitContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ArcParserRULE_compilationUnit
}

func (*CompilationUnitContext) IsCompilationUnitContext() {}

func NewCompilationUnitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CompilationUnitContext {
	var p = new(CompilationUnitContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ArcParserRULE_compilationUnit

	return p
}

func (s *CompilationUnitContext) GetParser() antlr.Parser { return s.parser }

func (s *CompilationUnitContext) EOF() antlr.TerminalNode {
	return s.GetToken(ArcParserEOF, 0)
}

func (s *CompilationUnitContext) AllImportDecl() []IImportDeclContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IImportDeclContext); ok {
			len++
		}
	}

	tst := make([]IImportDeclContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IImportDeclContext); ok {
			tst[i] = t.(IImportDeclContext)
			i++
		}
	}

	return tst
}

func (s *CompilationUnitContext) ImportDecl(i int) IImportDeclContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IImportDeclContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IImportDeclContext)
}

func (s *CompilationUnitContext) AllNamespaceDecl() []INamespaceDeclContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(INamespaceDeclContext); ok {
			len++
		}
	}

	tst := make([]INamespaceDeclContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(INamespaceDeclContext); ok {
			tst[i] = t.(INamespaceDeclContext)
			i++
		}
	}

	return tst
}

func (s *CompilationUnitContext) NamespaceDecl(i int) INamespaceDeclContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INamespaceDeclContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(INamespaceDeclContext)
}

func (s *CompilationUnitContext) AllTopLevelDecl() []ITopLevelDeclContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITopLevelDeclContext); ok {
			len++
		}
	}

	tst := make([]ITopLevelDeclContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITopLevelDeclContext); ok {
			tst[i] = t.(ITopLevelDeclContext)
			i++
		}
	}

	return tst
}

func (s *CompilationUnitContext) TopLevelDecl(i int) ITopLevelDeclContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITopLevelDeclContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITopLevelDeclContext)
}

func (s *CompilationUnitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CompilationUnitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CompilationUnitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ArcParserListener); ok {
		listenerT.EnterCompilationUnit(s)
	}
}

func (s *CompilationUnitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ArcParserListener); ok {
		listenerT.ExitCompilationUnit(s)
	}
}

func (s *CompilationUnitContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ArcParserVisitor:
		return t.VisitCompilationUnit(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ArcParser) CompilationUnit() (localctx ICompilationUnitContext) {
	localctx = NewCompilationUnitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ArcParserRULE_compilationUnit)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(111)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&16510) != 0 {
		p.SetState(109)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case ArcParserIMPORT:
			{
				p.SetState(106)
				p.ImportDecl()
			}

		case ArcParserNAMESPACE:
			{
				p.SetState(107)
				p.NamespaceDecl()
			}

		case ArcParserLET, ArcParserCONST, ArcParserFUNC, ArcParserSTRUCT, ArcParserEXTERN:
			{
				p.SetState(108)
				p.TopLevelDecl()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(113)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(114)
		p.Match(ArcParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IImportDeclContext is an interface to support dynamic dispatch.
type IImportDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IMPORT() antlr.TerminalNode
	STRING_LITERAL() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	AllImportSpec() []IImportSpecContext
	ImportSpec(i int) IImportSpecContext

	// IsImportDeclContext differentiates from other interfaces.
	IsImportDeclContext()
}

type ImportDeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyImportDeclContext() *ImportDeclContext {
	var p = new(ImportDeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ArcParserRULE_importDecl
	return p
}

func InitEmptyImportDeclContext(p *ImportDeclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ArcParserRULE_importDecl
}

func (*ImportDeclContext) IsImportDeclContext() {}

func NewImportDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ImportDeclContext {
	var p = new(ImportDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ArcParserRULE_importDecl

	return p
}

func (s *ImportDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *ImportDeclContext) IMPORT() antlr.TerminalNode {
	return s.GetToken(ArcParserIMPORT, 0)
}

func (s *ImportDeclContext) STRING_LITERAL() antlr.TerminalNode {
	return s.GetToken(ArcParserSTRING_LITERAL, 0)
}

func (s *ImportDeclContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(ArcParserLPAREN, 0)
}

func (s *ImportDeclContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(ArcParserRPAREN, 0)
}

func (s *ImportDeclContext) AllImportSpec() []IImportSpecContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IImportSpecContext); ok {
			len++
		}
	}

	tst := make([]IImportSpecContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IImportSpecContext); ok {
			tst[i] = t.(IImportSpecContext)
			i++
		}
	}

	return tst
}

func (s *ImportDeclContext) ImportSpec(i int) IImportSpecContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IImportSpecContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IImportSpecContext)
}

func (s *ImportDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ImportDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ImportDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ArcParserListener); ok {
		listenerT.EnterImportDecl(s)
	}
}

func (s *ImportDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ArcParserListener); ok {
		listenerT.ExitImportDecl(s)
	}
}

func (s *ImportDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ArcParserVisitor:
		return t.VisitImportDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ArcParser) ImportDecl() (localctx IImportDeclContext) {
	localctx = NewImportDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ArcParserRULE_importDecl)
	var _la int

	p.SetState(127)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(116)
			p.Match(ArcParserIMPORT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(117)
			p.Match(ArcParserSTRING_LITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(118)
			p.Match(ArcParserIMPORT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(119)
			p.Match(ArcParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(123)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == ArcParserSTRING_LITERAL {
			{
				p.SetState(120)
				p.ImportSpec()
			}

			p.SetState(125)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(126)
			p.Match(ArcParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IImportSpecContext is an interface to support dynamic dispatch.
type IImportSpecContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STRING_LITERAL() antlr.TerminalNode

	// IsImportSpecContext differentiates from other interfaces.
	IsImportSpecContext()
}

type ImportSpecContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyImportSpecContext() *ImportSpecContext {
	var p = new(ImportSpecContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ArcParserRULE_importSpec
	return p
}

func InitEmptyImportSpecContext(p *ImportSpecContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ArcParserRULE_importSpec
}

func (*ImportSpecContext) IsImportSpecContext() {}

func NewImportSpecContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ImportSpecContext {
	var p = new(ImportSpecContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ArcParserRULE_importSpec

	return p
}

func (s *ImportSpecContext) GetParser() antlr.Parser { return s.parser }

func (s *ImportSpecContext) STRING_LITERAL() antlr.TerminalNode {
	return s.GetToken(ArcParserSTRING_LITERAL, 0)
}

func (s *ImportSpecContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ImportSpecContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ImportSpecContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ArcParserListener); ok {
		listenerT.EnterImportSpec(s)
	}
}

func (s *ImportSpecContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ArcParserListener); ok {
		listenerT.ExitImportSpec(s)
	}
}

func (s *ImportSpecContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ArcParserVisitor:
		return t.VisitImportSpec(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ArcParser) ImportSpec() (localctx IImportSpecContext) {
	localctx = NewImportSpecContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, ArcParserRULE_importSpec)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(129)
		p.Match(ArcParserSTRING_LITERAL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// INamespaceDeclContext is an interface to support dynamic dispatch.
type INamespaceDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NAMESPACE() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode

	// IsNamespaceDeclContext differentiates from other interfaces.
	IsNamespaceDeclContext()
}

type NamespaceDeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNamespaceDeclContext() *NamespaceDeclContext {
	var p = new(NamespaceDeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ArcParserRULE_namespaceDecl
	return p
}

func InitEmptyNamespaceDeclContext(p *NamespaceDeclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ArcParserRULE_namespaceDecl
}

func (*NamespaceDeclContext) IsNamespaceDeclContext() {}

func NewNamespaceDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NamespaceDeclContext {
	var p = new(NamespaceDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ArcParserRULE_namespaceDecl

	return p
}

func (s *NamespaceDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *NamespaceDeclContext) NAMESPACE() antlr.TerminalNode {
	return s.GetToken(ArcParserNAMESPACE, 0)
}

func (s *NamespaceDeclContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ArcParserIDENTIFIER, 0)
}

func (s *NamespaceDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NamespaceDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NamespaceDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ArcParserListener); ok {
		listenerT.EnterNamespaceDecl(s)
	}
}

func (s *NamespaceDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ArcParserListener); ok {
		listenerT.ExitNamespaceDecl(s)
	}
}

func (s *NamespaceDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ArcParserVisitor:
		return t.VisitNamespaceDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ArcParser) NamespaceDecl() (localctx INamespaceDeclContext) {
	localctx = NewNamespaceDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, ArcParserRULE_namespaceDecl)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(131)
		p.Match(ArcParserNAMESPACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(132)
		p.Match(ArcParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITopLevelDeclContext is an interface to support dynamic dispatch.
type ITopLevelDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FunctionDecl() IFunctionDeclContext
	StructDecl() IStructDeclContext
	VariableDecl() IVariableDeclContext
	ConstDecl() IConstDeclContext
	ExternDecl() IExternDeclContext

	// IsTopLevelDeclContext differentiates from other interfaces.
	IsTopLevelDeclContext()
}

type TopLevelDeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTopLevelDeclContext() *TopLevelDeclContext {
	var p = new(TopLevelDeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ArcParserRULE_topLevelDecl
	return p
}

func InitEmptyTopLevelDeclContext(p *TopLevelDeclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ArcParserRULE_topLevelDecl
}

func (*TopLevelDeclContext) IsTopLevelDeclContext() {}

func NewTopLevelDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TopLevelDeclContext {
	var p = new(TopLevelDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ArcParserRULE_topLevelDecl

	return p
}

func (s *TopLevelDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *TopLevelDeclContext) FunctionDecl() IFunctionDeclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionDeclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionDeclContext)
}

func (s *TopLevelDeclContext) StructDecl() IStructDeclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStructDeclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStructDeclContext)
}

func (s *TopLevelDeclContext) VariableDecl() IVariableDeclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariableDeclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVariableDeclContext)
}

func (s *TopLevelDeclContext) ConstDecl() IConstDeclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConstDeclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConstDeclContext)
}

func (s *TopLevelDeclContext) ExternDecl() IExternDeclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExternDeclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExternDeclContext)
}

func (s *TopLevelDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TopLevelDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TopLevelDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ArcParserListener); ok {
		listenerT.EnterTopLevelDecl(s)
	}
}

func (s *TopLevelDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ArcParserListener); ok {
		listenerT.ExitTopLevelDecl(s)
	}
}

func (s *TopLevelDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ArcParserVisitor:
		return t.VisitTopLevelDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ArcParser) TopLevelDecl() (localctx ITopLevelDeclContext) {
	localctx = NewTopLevelDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, ArcParserRULE_topLevelDecl)
	p.SetState(139)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ArcParserFUNC:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(134)
			p.FunctionDecl()
		}

	case ArcParserSTRUCT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(135)
			p.StructDecl()
		}

	case ArcParserLET:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(136)
			p.VariableDecl()
		}

	case ArcParserCONST:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(137)
			p.ConstDecl()
		}

	case ArcParserEXTERN:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(138)
			p.ExternDecl()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExternDeclContext is an interface to support dynamic dispatch.
type IExternDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EXTERN() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	AllExternMember() []IExternMemberContext
	ExternMember(i int) IExternMemberContext

	// IsExternDeclContext differentiates from other interfaces.
	IsExternDeclContext()
}

type ExternDeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExternDeclContext() *ExternDeclContext {
	var p = new(ExternDeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ArcParserRULE_externDecl
	return p
}

func InitEmptyExternDeclContext(p *ExternDeclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ArcParserRULE_externDecl
}

func (*ExternDeclContext) IsExternDeclContext() {}

func NewExternDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExternDeclContext {
	var p = new(ExternDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ArcParserRULE_externDecl

	return p
}

func (s *ExternDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *ExternDeclContext) EXTERN() antlr.TerminalNode {
	return s.GetToken(ArcParserEXTERN, 0)
}

func (s *ExternDeclContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ArcParserIDENTIFIER, 0)
}

func (s *ExternDeclContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(ArcParserLBRACE, 0)
}

func (s *ExternDeclContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(ArcParserRBRACE, 0)
}

func (s *ExternDeclContext) AllExternMember() []IExternMemberContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExternMemberContext); ok {
			len++
		}
	}

	tst := make([]IExternMemberContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExternMemberContext); ok {
			tst[i] = t.(IExternMemberContext)
			i++
		}
	}

	return tst
}

func (s *ExternDeclContext) ExternMember(i int) IExternMemberContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExternMemberContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExternMemberContext)
}

func (s *ExternDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExternDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExternDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ArcParserListener); ok {
		listenerT.EnterExternDecl(s)
	}
}

func (s *ExternDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ArcParserListener); ok {
		listenerT.ExitExternDecl(s)
	}
}

func (s *ExternDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ArcParserVisitor:
		return t.VisitExternDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ArcParser) ExternDecl() (localctx IExternDeclContext) {
	localctx = NewExternDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, ArcParserRULE_externDecl)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(141)
		p.Match(ArcParserEXTERN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(142)
		p.Match(ArcParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(143)
		p.Match(ArcParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(147)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ArcParserFUNC {
		{
			p.SetState(144)
			p.ExternMember()
		}

		p.SetState(149)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(150)
		p.Match(ArcParserRBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExternMemberContext is an interface to support dynamic dispatch.
type IExternMemberContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ExternFunctionDecl() IExternFunctionDeclContext

	// IsExternMemberContext differentiates from other interfaces.
	IsExternMemberContext()
}

type ExternMemberContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExternMemberContext() *ExternMemberContext {
	var p = new(ExternMemberContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ArcParserRULE_externMember
	return p
}

func InitEmptyExternMemberContext(p *ExternMemberContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ArcParserRULE_externMember
}

func (*ExternMemberContext) IsExternMemberContext() {}

func NewExternMemberContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExternMemberContext {
	var p = new(ExternMemberContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ArcParserRULE_externMember

	return p
}

func (s *ExternMemberContext) GetParser() antlr.Parser { return s.parser }

func (s *ExternMemberContext) ExternFunctionDecl() IExternFunctionDeclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExternFunctionDeclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExternFunctionDeclContext)
}

func (s *ExternMemberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExternMemberContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExternMemberContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ArcParserListener); ok {
		listenerT.EnterExternMember(s)
	}
}

func (s *ExternMemberContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ArcParserListener); ok {
		listenerT.ExitExternMember(s)
	}
}

func (s *ExternMemberContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ArcParserVisitor:
		return t.VisitExternMember(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ArcParser) ExternMember() (localctx IExternMemberContext) {
	localctx = NewExternMemberContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, ArcParserRULE_externMember)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(152)
		p.ExternFunctionDecl()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExternFunctionDeclContext is an interface to support dynamic dispatch.
type IExternFunctionDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FUNC() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	ExternParameterList() IExternParameterListContext
	Type_() ITypeContext

	// IsExternFunctionDeclContext differentiates from other interfaces.
	IsExternFunctionDeclContext()
}

type ExternFunctionDeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExternFunctionDeclContext() *ExternFunctionDeclContext {
	var p = new(ExternFunctionDeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ArcParserRULE_externFunctionDecl
	return p
}

func InitEmptyExternFunctionDeclContext(p *ExternFunctionDeclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ArcParserRULE_externFunctionDecl
}

func (*ExternFunctionDeclContext) IsExternFunctionDeclContext() {}

func NewExternFunctionDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExternFunctionDeclContext {
	var p = new(ExternFunctionDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ArcParserRULE_externFunctionDecl

	return p
}

func (s *ExternFunctionDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *ExternFunctionDeclContext) FUNC() antlr.TerminalNode {
	return s.GetToken(ArcParserFUNC, 0)
}

func (s *ExternFunctionDeclContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ArcParserIDENTIFIER, 0)
}

func (s *ExternFunctionDeclContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(ArcParserLPAREN, 0)
}

func (s *ExternFunctionDeclContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(ArcParserRPAREN, 0)
}

func (s *ExternFunctionDeclContext) ExternParameterList() IExternParameterListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExternParameterListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExternParameterListContext)
}

func (s *ExternFunctionDeclContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *ExternFunctionDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExternFunctionDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExternFunctionDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ArcParserListener); ok {
		listenerT.EnterExternFunctionDecl(s)
	}
}

func (s *ExternFunctionDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ArcParserListener); ok {
		listenerT.ExitExternFunctionDecl(s)
	}
}

func (s *ExternFunctionDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ArcParserVisitor:
		return t.VisitExternFunctionDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ArcParser) ExternFunctionDecl() (localctx IExternFunctionDeclContext) {
	localctx = NewExternFunctionDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, ArcParserRULE_externFunctionDecl)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(154)
		p.Match(ArcParserFUNC)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(155)
		p.Match(ArcParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(156)
		p.Match(ArcParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(158)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64((_la-16)) & ^0x3f) == 0 && ((int64(1)<<(_la-16))&36591781341036543) != 0 {
		{
			p.SetState(157)
			p.ExternParameterList()
		}

	}
	{
		p.SetState(160)
		p.Match(ArcParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(162)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64((_la-16)) & ^0x3f) == 0 && ((int64(1)<<(_la-16))&36028831387615231) != 0 {
		{
			p.SetState(161)
			p.Type_()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExternParameterListContext is an interface to support dynamic dispatch.
type IExternParameterListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllType_() []ITypeContext
	Type_(i int) ITypeContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode
	ELLIPSIS() antlr.TerminalNode

	// IsExternParameterListContext differentiates from other interfaces.
	IsExternParameterListContext()
}

type ExternParameterListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExternParameterListContext() *ExternParameterListContext {
	var p = new(ExternParameterListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ArcParserRULE_externParameterList
	return p
}

func InitEmptyExternParameterListContext(p *ExternParameterListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ArcParserRULE_externParameterList
}

func (*ExternParameterListContext) IsExternParameterListContext() {}

func NewExternParameterListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExternParameterListContext {
	var p = new(ExternParameterListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ArcParserRULE_externParameterList

	return p
}

func (s *ExternParameterListContext) GetParser() antlr.Parser { return s.parser }

func (s *ExternParameterListContext) AllType_() []ITypeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITypeContext); ok {
			len++
		}
	}

	tst := make([]ITypeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITypeContext); ok {
			tst[i] = t.(ITypeContext)
			i++
		}
	}

	return tst
}

func (s *ExternParameterListContext) Type_(i int) ITypeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *ExternParameterListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(ArcParserCOMMA)
}

func (s *ExternParameterListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(ArcParserCOMMA, i)
}

func (s *ExternParameterListContext) ELLIPSIS() antlr.TerminalNode {
	return s.GetToken(ArcParserELLIPSIS, 0)
}

func (s *ExternParameterListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExternParameterListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExternParameterListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ArcParserListener); ok {
		listenerT.EnterExternParameterList(s)
	}
}

func (s *ExternParameterListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ArcParserListener); ok {
		listenerT.ExitExternParameterList(s)
	}
}

func (s *ExternParameterListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ArcParserVisitor:
		return t.VisitExternParameterList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ArcParser) ExternParameterList() (localctx IExternParameterListContext) {
	localctx = NewExternParameterListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, ArcParserRULE_externParameterList)
	var _la int

	var _alt int

	p.SetState(177)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ArcParserINT8, ArcParserINT16, ArcParserINT32, ArcParserINT64, ArcParserUINT8, ArcParserUINT16, ArcParserUINT32, ArcParserUINT64, ArcParserUSIZE, ArcParserISIZE, ArcParserFLOAT32, ArcParserFLOAT64, ArcParserBYTE, ArcParserBOOL, ArcParserCHAR, ArcParserSTRING, ArcParserVOID, ArcParserVECTOR, ArcParserMAP, ArcParserSTAR, ArcParserAMP, ArcParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(164)
			p.Type_()
		}
		p.SetState(169)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(165)
					p.Match(ArcParserCOMMA)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(166)
					p.Type_()
				}

			}
			p.SetState(171)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}
		p.SetState(174)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == ArcParserCOMMA {
			{
				p.SetState(172)
				p.Match(ArcParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(173)
				p.Match(ArcParserELLIPSIS)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case ArcParserELLIPSIS:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(176)
			p.Match(ArcParserELLIPSIS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFunctionDeclContext is an interface to support dynamic dispatch.
type IFunctionDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FUNC() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	Block() IBlockContext
	ParameterList() IParameterListContext
	Type_() ITypeContext

	// IsFunctionDeclContext differentiates from other interfaces.
	IsFunctionDeclContext()
}

type FunctionDeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionDeclContext() *FunctionDeclContext {
	var p = new(FunctionDeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ArcParserRULE_functionDecl
	return p
}

func InitEmptyFunctionDeclContext(p *FunctionDeclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ArcParserRULE_functionDecl
}

func (*FunctionDeclContext) IsFunctionDeclContext() {}

func NewFunctionDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionDeclContext {
	var p = new(FunctionDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ArcParserRULE_functionDecl

	return p
}

func (s *FunctionDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionDeclContext) FUNC() antlr.TerminalNode {
	return s.GetToken(ArcParserFUNC, 0)
}

func (s *FunctionDeclContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ArcParserIDENTIFIER, 0)
}

func (s *FunctionDeclContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(ArcParserLPAREN, 0)
}

func (s *FunctionDeclContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(ArcParserRPAREN, 0)
}

func (s *FunctionDeclContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *FunctionDeclContext) ParameterList() IParameterListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParameterListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParameterListContext)
}

func (s *FunctionDeclContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *FunctionDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ArcParserListener); ok {
		listenerT.EnterFunctionDecl(s)
	}
}

func (s *FunctionDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ArcParserListener); ok {
		listenerT.ExitFunctionDecl(s)
	}
}

func (s *FunctionDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ArcParserVisitor:
		return t.VisitFunctionDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ArcParser) FunctionDecl() (localctx IFunctionDeclContext) {
	localctx = NewFunctionDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, ArcParserRULE_functionDecl)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(179)
		p.Match(ArcParserFUNC)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(180)
		p.Match(ArcParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(181)
		p.Match(ArcParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(183)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64((_la-15)) & ^0x3f) == 0 && ((int64(1)<<(_la-15))&73183493944770561) != 0 {
		{
			p.SetState(182)
			p.ParameterList()
		}

	}
	{
		p.SetState(185)
		p.Match(ArcParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(187)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64((_la-16)) & ^0x3f) == 0 && ((int64(1)<<(_la-16))&36028831387615231) != 0 {
		{
			p.SetState(186)
			p.Type_()
		}

	}
	{
		p.SetState(189)
		p.Block()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IParameterListContext is an interface to support dynamic dispatch.
type IParameterListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllParameter() []IParameterContext
	Parameter(i int) IParameterContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode
	ELLIPSIS() antlr.TerminalNode

	// IsParameterListContext differentiates from other interfaces.
	IsParameterListContext()
}

type ParameterListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParameterListContext() *ParameterListContext {
	var p = new(ParameterListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ArcParserRULE_parameterList
	return p
}

func InitEmptyParameterListContext(p *ParameterListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ArcParserRULE_parameterList
}

func (*ParameterListContext) IsParameterListContext() {}

func NewParameterListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParameterListContext {
	var p = new(ParameterListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ArcParserRULE_parameterList

	return p
}

func (s *ParameterListContext) GetParser() antlr.Parser { return s.parser }

func (s *ParameterListContext) AllParameter() []IParameterContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IParameterContext); ok {
			len++
		}
	}

	tst := make([]IParameterContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IParameterContext); ok {
			tst[i] = t.(IParameterContext)
			i++
		}
	}

	return tst
}

func (s *ParameterListContext) Parameter(i int) IParameterContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParameterContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParameterContext)
}

func (s *ParameterListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(ArcParserCOMMA)
}

func (s *ParameterListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(ArcParserCOMMA, i)
}

func (s *ParameterListContext) ELLIPSIS() antlr.TerminalNode {
	return s.GetToken(ArcParserELLIPSIS, 0)
}

func (s *ParameterListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParameterListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParameterListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ArcParserListener); ok {
		listenerT.EnterParameterList(s)
	}
}

func (s *ParameterListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ArcParserListener); ok {
		listenerT.ExitParameterList(s)
	}
}

func (s *ParameterListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ArcParserVisitor:
		return t.VisitParameterList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ArcParser) ParameterList() (localctx IParameterListContext) {
	localctx = NewParameterListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, ArcParserRULE_parameterList)
	var _la int

	var _alt int

	p.SetState(204)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ArcParserSELF, ArcParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(191)
			p.Parameter()
		}
		p.SetState(196)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 13, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(192)
					p.Match(ArcParserCOMMA)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(193)
					p.Parameter()
				}

			}
			p.SetState(198)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 13, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}
		p.SetState(201)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == ArcParserCOMMA {
			{
				p.SetState(199)
				p.Match(ArcParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(200)
				p.Match(ArcParserELLIPSIS)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case ArcParserELLIPSIS:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(203)
			p.Match(ArcParserELLIPSIS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IParameterContext is an interface to support dynamic dispatch.
type IParameterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	COLON() antlr.TerminalNode
	Type_() ITypeContext
	SELF() antlr.TerminalNode

	// IsParameterContext differentiates from other interfaces.
	IsParameterContext()
}

type ParameterContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParameterContext() *ParameterContext {
	var p = new(ParameterContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ArcParserRULE_parameter
	return p
}

func InitEmptyParameterContext(p *ParameterContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ArcParserRULE_parameter
}

func (*ParameterContext) IsParameterContext() {}

func NewParameterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParameterContext {
	var p = new(ParameterContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ArcParserRULE_parameter

	return p
}

func (s *ParameterContext) GetParser() antlr.Parser { return s.parser }

func (s *ParameterContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ArcParserIDENTIFIER, 0)
}

func (s *ParameterContext) COLON() antlr.TerminalNode {
	return s.GetToken(ArcParserCOLON, 0)
}

func (s *ParameterContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *ParameterContext) SELF() antlr.TerminalNode {
	return s.GetToken(ArcParserSELF, 0)
}

func (s *ParameterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParameterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParameterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ArcParserListener); ok {
		listenerT.EnterParameter(s)
	}
}

func (s *ParameterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ArcParserListener); ok {
		listenerT.ExitParameter(s)
	}
}

func (s *ParameterContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ArcParserVisitor:
		return t.VisitParameter(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ArcParser) Parameter() (localctx IParameterContext) {
	localctx = NewParameterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, ArcParserRULE_parameter)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(207)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ArcParserSELF {
		{
			p.SetState(206)
			p.Match(ArcParserSELF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(209)
		p.Match(ArcParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(210)
		p.Match(ArcParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(211)
		p.Type_()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStructDeclContext is an interface to support dynamic dispatch.
type IStructDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STRUCT() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	AllStructField() []IStructFieldContext
	StructField(i int) IStructFieldContext

	// IsStructDeclContext differentiates from other interfaces.
	IsStructDeclContext()
}

type StructDeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStructDeclContext() *StructDeclContext {
	var p = new(StructDeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ArcParserRULE_structDecl
	return p
}

func InitEmptyStructDeclContext(p *StructDeclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ArcParserRULE_structDecl
}

func (*StructDeclContext) IsStructDeclContext() {}

func NewStructDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructDeclContext {
	var p = new(StructDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ArcParserRULE_structDecl

	return p
}

func (s *StructDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *StructDeclContext) STRUCT() antlr.TerminalNode {
	return s.GetToken(ArcParserSTRUCT, 0)
}

func (s *StructDeclContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ArcParserIDENTIFIER, 0)
}

func (s *StructDeclContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(ArcParserLBRACE, 0)
}

func (s *StructDeclContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(ArcParserRBRACE, 0)
}

func (s *StructDeclContext) AllStructField() []IStructFieldContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStructFieldContext); ok {
			len++
		}
	}

	tst := make([]IStructFieldContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStructFieldContext); ok {
			tst[i] = t.(IStructFieldContext)
			i++
		}
	}

	return tst
}

func (s *StructDeclContext) StructField(i int) IStructFieldContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStructFieldContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStructFieldContext)
}

func (s *StructDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StructDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ArcParserListener); ok {
		listenerT.EnterStructDecl(s)
	}
}

func (s *StructDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ArcParserListener); ok {
		listenerT.ExitStructDecl(s)
	}
}

func (s *StructDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ArcParserVisitor:
		return t.VisitStructDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ArcParser) StructDecl() (localctx IStructDeclContext) {
	localctx = NewStructDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, ArcParserRULE_structDecl)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(213)
		p.Match(ArcParserSTRUCT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(214)
		p.Match(ArcParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(215)
		p.Match(ArcParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(219)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ArcParserIDENTIFIER {
		{
			p.SetState(216)
			p.StructField()
		}

		p.SetState(221)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(222)
		p.Match(ArcParserRBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStructFieldContext is an interface to support dynamic dispatch.
type IStructFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	COLON() antlr.TerminalNode
	Type_() ITypeContext

	// IsStructFieldContext differentiates from other interfaces.
	IsStructFieldContext()
}

type StructFieldContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStructFieldContext() *StructFieldContext {
	var p = new(StructFieldContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ArcParserRULE_structField
	return p
}

func InitEmptyStructFieldContext(p *StructFieldContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ArcParserRULE_structField
}

func (*StructFieldContext) IsStructFieldContext() {}

func NewStructFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructFieldContext {
	var p = new(StructFieldContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ArcParserRULE_structField

	return p
}

func (s *StructFieldContext) GetParser() antlr.Parser { return s.parser }

func (s *StructFieldContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ArcParserIDENTIFIER, 0)
}

func (s *StructFieldContext) COLON() antlr.TerminalNode {
	return s.GetToken(ArcParserCOLON, 0)
}

func (s *StructFieldContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *StructFieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructFieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StructFieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ArcParserListener); ok {
		listenerT.EnterStructField(s)
	}
}

func (s *StructFieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ArcParserListener); ok {
		listenerT.ExitStructField(s)
	}
}

func (s *StructFieldContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ArcParserVisitor:
		return t.VisitStructField(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ArcParser) StructField() (localctx IStructFieldContext) {
	localctx = NewStructFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, ArcParserRULE_structField)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(224)
		p.Match(ArcParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(225)
		p.Match(ArcParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(226)
		p.Type_()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVariableDeclContext is an interface to support dynamic dispatch.
type IVariableDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LET() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode
	ASSIGN() antlr.TerminalNode
	Expression() IExpressionContext
	COLON() antlr.TerminalNode
	Type_() ITypeContext

	// IsVariableDeclContext differentiates from other interfaces.
	IsVariableDeclContext()
}

type VariableDeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableDeclContext() *VariableDeclContext {
	var p = new(VariableDeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ArcParserRULE_variableDecl
	return p
}

func InitEmptyVariableDeclContext(p *VariableDeclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ArcParserRULE_variableDecl
}

func (*VariableDeclContext) IsVariableDeclContext() {}

func NewVariableDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableDeclContext {
	var p = new(VariableDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ArcParserRULE_variableDecl

	return p
}

func (s *VariableDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableDeclContext) LET() antlr.TerminalNode {
	return s.GetToken(ArcParserLET, 0)
}

func (s *VariableDeclContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ArcParserIDENTIFIER, 0)
}

func (s *VariableDeclContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(ArcParserASSIGN, 0)
}

func (s *VariableDeclContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *VariableDeclContext) COLON() antlr.TerminalNode {
	return s.GetToken(ArcParserCOLON, 0)
}

func (s *VariableDeclContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *VariableDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ArcParserListener); ok {
		listenerT.EnterVariableDecl(s)
	}
}

func (s *VariableDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ArcParserListener); ok {
		listenerT.ExitVariableDecl(s)
	}
}

func (s *VariableDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ArcParserVisitor:
		return t.VisitVariableDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ArcParser) VariableDecl() (localctx IVariableDeclContext) {
	localctx = NewVariableDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, ArcParserRULE_variableDecl)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(228)
		p.Match(ArcParserLET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(229)
		p.Match(ArcParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(232)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ArcParserCOLON {
		{
			p.SetState(230)
			p.Match(ArcParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(231)
			p.Type_()
		}

	}
	{
		p.SetState(234)
		p.Match(ArcParserASSIGN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(235)
		p.Expression()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IConstDeclContext is an interface to support dynamic dispatch.
type IConstDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CONST() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode
	ASSIGN() antlr.TerminalNode
	Expression() IExpressionContext
	COLON() antlr.TerminalNode
	Type_() ITypeContext

	// IsConstDeclContext differentiates from other interfaces.
	IsConstDeclContext()
}

type ConstDeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConstDeclContext() *ConstDeclContext {
	var p = new(ConstDeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ArcParserRULE_constDecl
	return p
}

func InitEmptyConstDeclContext(p *ConstDeclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ArcParserRULE_constDecl
}

func (*ConstDeclContext) IsConstDeclContext() {}

func NewConstDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstDeclContext {
	var p = new(ConstDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ArcParserRULE_constDecl

	return p
}

func (s *ConstDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstDeclContext) CONST() antlr.TerminalNode {
	return s.GetToken(ArcParserCONST, 0)
}

func (s *ConstDeclContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ArcParserIDENTIFIER, 0)
}

func (s *ConstDeclContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(ArcParserASSIGN, 0)
}

func (s *ConstDeclContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ConstDeclContext) COLON() antlr.TerminalNode {
	return s.GetToken(ArcParserCOLON, 0)
}

func (s *ConstDeclContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *ConstDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConstDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ArcParserListener); ok {
		listenerT.EnterConstDecl(s)
	}
}

func (s *ConstDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ArcParserListener); ok {
		listenerT.ExitConstDecl(s)
	}
}

func (s *ConstDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ArcParserVisitor:
		return t.VisitConstDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ArcParser) ConstDecl() (localctx IConstDeclContext) {
	localctx = NewConstDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, ArcParserRULE_constDecl)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(237)
		p.Match(ArcParserCONST)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(238)
		p.Match(ArcParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(241)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ArcParserCOLON {
		{
			p.SetState(239)
			p.Match(ArcParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(240)
			p.Type_()
		}

	}
	{
		p.SetState(243)
		p.Match(ArcParserASSIGN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(244)
		p.Expression()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITypeContext is an interface to support dynamic dispatch.
type ITypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PrimitiveType() IPrimitiveTypeContext
	PointerType() IPointerTypeContext
	ReferenceType() IReferenceTypeContext
	VectorType() IVectorTypeContext
	MapType() IMapTypeContext
	IDENTIFIER() antlr.TerminalNode

	// IsTypeContext differentiates from other interfaces.
	IsTypeContext()
}

type TypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeContext() *TypeContext {
	var p = new(TypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ArcParserRULE_type
	return p
}

func InitEmptyTypeContext(p *TypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ArcParserRULE_type
}

func (*TypeContext) IsTypeContext() {}

func NewTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeContext {
	var p = new(TypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ArcParserRULE_type

	return p
}

func (s *TypeContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeContext) PrimitiveType() IPrimitiveTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimitiveTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimitiveTypeContext)
}

func (s *TypeContext) PointerType() IPointerTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPointerTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPointerTypeContext)
}

func (s *TypeContext) ReferenceType() IReferenceTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReferenceTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReferenceTypeContext)
}

func (s *TypeContext) VectorType() IVectorTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVectorTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVectorTypeContext)
}

func (s *TypeContext) MapType() IMapTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMapTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMapTypeContext)
}

func (s *TypeContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ArcParserIDENTIFIER, 0)
}

func (s *TypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ArcParserListener); ok {
		listenerT.EnterType(s)
	}
}

func (s *TypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ArcParserListener); ok {
		listenerT.ExitType(s)
	}
}

func (s *TypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ArcParserVisitor:
		return t.VisitType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ArcParser) Type_() (localctx ITypeContext) {
	localctx = NewTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, ArcParserRULE_type)
	p.SetState(252)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ArcParserINT8, ArcParserINT16, ArcParserINT32, ArcParserINT64, ArcParserUINT8, ArcParserUINT16, ArcParserUINT32, ArcParserUINT64, ArcParserUSIZE, ArcParserISIZE, ArcParserFLOAT32, ArcParserFLOAT64, ArcParserBYTE, ArcParserBOOL, ArcParserCHAR, ArcParserSTRING, ArcParserVOID:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(246)
			p.PrimitiveType()
		}

	case ArcParserSTAR:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(247)
			p.PointerType()
		}

	case ArcParserAMP:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(248)
			p.ReferenceType()
		}

	case ArcParserVECTOR:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(249)
			p.VectorType()
		}

	case ArcParserMAP:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(250)
			p.MapType()
		}

	case ArcParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(251)
			p.Match(ArcParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPrimitiveTypeContext is an interface to support dynamic dispatch.
type IPrimitiveTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INT8() antlr.TerminalNode
	INT16() antlr.TerminalNode
	INT32() antlr.TerminalNode
	INT64() antlr.TerminalNode
	UINT8() antlr.TerminalNode
	UINT16() antlr.TerminalNode
	UINT32() antlr.TerminalNode
	UINT64() antlr.TerminalNode
	USIZE() antlr.TerminalNode
	ISIZE() antlr.TerminalNode
	FLOAT32() antlr.TerminalNode
	FLOAT64() antlr.TerminalNode
	BYTE() antlr.TerminalNode
	BOOL() antlr.TerminalNode
	CHAR() antlr.TerminalNode
	STRING() antlr.TerminalNode
	VOID() antlr.TerminalNode

	// IsPrimitiveTypeContext differentiates from other interfaces.
	IsPrimitiveTypeContext()
}

type PrimitiveTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrimitiveTypeContext() *PrimitiveTypeContext {
	var p = new(PrimitiveTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ArcParserRULE_primitiveType
	return p
}

func InitEmptyPrimitiveTypeContext(p *PrimitiveTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ArcParserRULE_primitiveType
}

func (*PrimitiveTypeContext) IsPrimitiveTypeContext() {}

func NewPrimitiveTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimitiveTypeContext {
	var p = new(PrimitiveTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ArcParserRULE_primitiveType

	return p
}

func (s *PrimitiveTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *PrimitiveTypeContext) INT8() antlr.TerminalNode {
	return s.GetToken(ArcParserINT8, 0)
}

func (s *PrimitiveTypeContext) INT16() antlr.TerminalNode {
	return s.GetToken(ArcParserINT16, 0)
}

func (s *PrimitiveTypeContext) INT32() antlr.TerminalNode {
	return s.GetToken(ArcParserINT32, 0)
}

func (s *PrimitiveTypeContext) INT64() antlr.TerminalNode {
	return s.GetToken(ArcParserINT64, 0)
}

func (s *PrimitiveTypeContext) UINT8() antlr.TerminalNode {
	return s.GetToken(ArcParserUINT8, 0)
}

func (s *PrimitiveTypeContext) UINT16() antlr.TerminalNode {
	return s.GetToken(ArcParserUINT16, 0)
}

func (s *PrimitiveTypeContext) UINT32() antlr.TerminalNode {
	return s.GetToken(ArcParserUINT32, 0)
}

func (s *PrimitiveTypeContext) UINT64() antlr.TerminalNode {
	return s.GetToken(ArcParserUINT64, 0)
}

func (s *PrimitiveTypeContext) USIZE() antlr.TerminalNode {
	return s.GetToken(ArcParserUSIZE, 0)
}

func (s *PrimitiveTypeContext) ISIZE() antlr.TerminalNode {
	return s.GetToken(ArcParserISIZE, 0)
}

func (s *PrimitiveTypeContext) FLOAT32() antlr.TerminalNode {
	return s.GetToken(ArcParserFLOAT32, 0)
}

func (s *PrimitiveTypeContext) FLOAT64() antlr.TerminalNode {
	return s.GetToken(ArcParserFLOAT64, 0)
}

func (s *PrimitiveTypeContext) BYTE() antlr.TerminalNode {
	return s.GetToken(ArcParserBYTE, 0)
}

func (s *PrimitiveTypeContext) BOOL() antlr.TerminalNode {
	return s.GetToken(ArcParserBOOL, 0)
}

func (s *PrimitiveTypeContext) CHAR() antlr.TerminalNode {
	return s.GetToken(ArcParserCHAR, 0)
}

func (s *PrimitiveTypeContext) STRING() antlr.TerminalNode {
	return s.GetToken(ArcParserSTRING, 0)
}

func (s *PrimitiveTypeContext) VOID() antlr.TerminalNode {
	return s.GetToken(ArcParserVOID, 0)
}

func (s *PrimitiveTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimitiveTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrimitiveTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ArcParserListener); ok {
		listenerT.EnterPrimitiveType(s)
	}
}

func (s *PrimitiveTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ArcParserListener); ok {
		listenerT.ExitPrimitiveType(s)
	}
}

func (s *PrimitiveTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ArcParserVisitor:
		return t.VisitPrimitiveType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ArcParser) PrimitiveType() (localctx IPrimitiveTypeContext) {
	localctx = NewPrimitiveTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, ArcParserRULE_primitiveType)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(254)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&8589869056) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPointerTypeContext is an interface to support dynamic dispatch.
type IPointerTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STAR() antlr.TerminalNode
	Type_() ITypeContext

	// IsPointerTypeContext differentiates from other interfaces.
	IsPointerTypeContext()
}

type PointerTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPointerTypeContext() *PointerTypeContext {
	var p = new(PointerTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ArcParserRULE_pointerType
	return p
}

func InitEmptyPointerTypeContext(p *PointerTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ArcParserRULE_pointerType
}

func (*PointerTypeContext) IsPointerTypeContext() {}

func NewPointerTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PointerTypeContext {
	var p = new(PointerTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ArcParserRULE_pointerType

	return p
}

func (s *PointerTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *PointerTypeContext) STAR() antlr.TerminalNode {
	return s.GetToken(ArcParserSTAR, 0)
}

func (s *PointerTypeContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *PointerTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PointerTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PointerTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ArcParserListener); ok {
		listenerT.EnterPointerType(s)
	}
}

func (s *PointerTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ArcParserListener); ok {
		listenerT.ExitPointerType(s)
	}
}

func (s *PointerTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ArcParserVisitor:
		return t.VisitPointerType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ArcParser) PointerType() (localctx IPointerTypeContext) {
	localctx = NewPointerTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, ArcParserRULE_pointerType)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(256)
		p.Match(ArcParserSTAR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(257)
		p.Type_()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IReferenceTypeContext is an interface to support dynamic dispatch.
type IReferenceTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AMP() antlr.TerminalNode
	Type_() ITypeContext

	// IsReferenceTypeContext differentiates from other interfaces.
	IsReferenceTypeContext()
}

type ReferenceTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReferenceTypeContext() *ReferenceTypeContext {
	var p = new(ReferenceTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ArcParserRULE_referenceType
	return p
}

func InitEmptyReferenceTypeContext(p *ReferenceTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ArcParserRULE_referenceType
}

func (*ReferenceTypeContext) IsReferenceTypeContext() {}

func NewReferenceTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReferenceTypeContext {
	var p = new(ReferenceTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ArcParserRULE_referenceType

	return p
}

func (s *ReferenceTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *ReferenceTypeContext) AMP() antlr.TerminalNode {
	return s.GetToken(ArcParserAMP, 0)
}

func (s *ReferenceTypeContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *ReferenceTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReferenceTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReferenceTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ArcParserListener); ok {
		listenerT.EnterReferenceType(s)
	}
}

func (s *ReferenceTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ArcParserListener); ok {
		listenerT.ExitReferenceType(s)
	}
}

func (s *ReferenceTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ArcParserVisitor:
		return t.VisitReferenceType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ArcParser) ReferenceType() (localctx IReferenceTypeContext) {
	localctx = NewReferenceTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, ArcParserRULE_referenceType)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(259)
		p.Match(ArcParserAMP)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(260)
		p.Type_()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVectorTypeContext is an interface to support dynamic dispatch.
type IVectorTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	VECTOR() antlr.TerminalNode
	LT() antlr.TerminalNode
	Type_() ITypeContext
	GT() antlr.TerminalNode

	// IsVectorTypeContext differentiates from other interfaces.
	IsVectorTypeContext()
}

type VectorTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVectorTypeContext() *VectorTypeContext {
	var p = new(VectorTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ArcParserRULE_vectorType
	return p
}

func InitEmptyVectorTypeContext(p *VectorTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ArcParserRULE_vectorType
}

func (*VectorTypeContext) IsVectorTypeContext() {}

func NewVectorTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VectorTypeContext {
	var p = new(VectorTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ArcParserRULE_vectorType

	return p
}

func (s *VectorTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *VectorTypeContext) VECTOR() antlr.TerminalNode {
	return s.GetToken(ArcParserVECTOR, 0)
}

func (s *VectorTypeContext) LT() antlr.TerminalNode {
	return s.GetToken(ArcParserLT, 0)
}

func (s *VectorTypeContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *VectorTypeContext) GT() antlr.TerminalNode {
	return s.GetToken(ArcParserGT, 0)
}

func (s *VectorTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VectorTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VectorTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ArcParserListener); ok {
		listenerT.EnterVectorType(s)
	}
}

func (s *VectorTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ArcParserListener); ok {
		listenerT.ExitVectorType(s)
	}
}

func (s *VectorTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ArcParserVisitor:
		return t.VisitVectorType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ArcParser) VectorType() (localctx IVectorTypeContext) {
	localctx = NewVectorTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, ArcParserRULE_vectorType)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(262)
		p.Match(ArcParserVECTOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(263)
		p.Match(ArcParserLT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(264)
		p.Type_()
	}
	{
		p.SetState(265)
		p.Match(ArcParserGT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMapTypeContext is an interface to support dynamic dispatch.
type IMapTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	MAP() antlr.TerminalNode
	LT() antlr.TerminalNode
	AllType_() []ITypeContext
	Type_(i int) ITypeContext
	COMMA() antlr.TerminalNode
	GT() antlr.TerminalNode

	// IsMapTypeContext differentiates from other interfaces.
	IsMapTypeContext()
}

type MapTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMapTypeContext() *MapTypeContext {
	var p = new(MapTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ArcParserRULE_mapType
	return p
}

func InitEmptyMapTypeContext(p *MapTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ArcParserRULE_mapType
}

func (*MapTypeContext) IsMapTypeContext() {}

func NewMapTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MapTypeContext {
	var p = new(MapTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ArcParserRULE_mapType

	return p
}

func (s *MapTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *MapTypeContext) MAP() antlr.TerminalNode {
	return s.GetToken(ArcParserMAP, 0)
}

func (s *MapTypeContext) LT() antlr.TerminalNode {
	return s.GetToken(ArcParserLT, 0)
}

func (s *MapTypeContext) AllType_() []ITypeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITypeContext); ok {
			len++
		}
	}

	tst := make([]ITypeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITypeContext); ok {
			tst[i] = t.(ITypeContext)
			i++
		}
	}

	return tst
}

func (s *MapTypeContext) Type_(i int) ITypeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *MapTypeContext) COMMA() antlr.TerminalNode {
	return s.GetToken(ArcParserCOMMA, 0)
}

func (s *MapTypeContext) GT() antlr.TerminalNode {
	return s.GetToken(ArcParserGT, 0)
}

func (s *MapTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MapTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MapTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ArcParserListener); ok {
		listenerT.EnterMapType(s)
	}
}

func (s *MapTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ArcParserListener); ok {
		listenerT.ExitMapType(s)
	}
}

func (s *MapTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ArcParserVisitor:
		return t.VisitMapType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ArcParser) MapType() (localctx IMapTypeContext) {
	localctx = NewMapTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, ArcParserRULE_mapType)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(267)
		p.Match(ArcParserMAP)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(268)
		p.Match(ArcParserLT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(269)
		p.Type_()
	}
	{
		p.SetState(270)
		p.Match(ArcParserCOMMA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(271)
		p.Type_()
	}
	{
		p.SetState(272)
		p.Match(ArcParserGT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBlockContext is an interface to support dynamic dispatch.
type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext

	// IsBlockContext differentiates from other interfaces.
	IsBlockContext()
}

type BlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockContext() *BlockContext {
	var p = new(BlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ArcParserRULE_block
	return p
}

func InitEmptyBlockContext(p *BlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ArcParserRULE_block
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ArcParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(ArcParserLBRACE, 0)
}

func (s *BlockContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(ArcParserRBRACE, 0)
}

func (s *BlockContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *BlockContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ArcParserListener); ok {
		listenerT.EnterBlock(s)
	}
}

func (s *BlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ArcParserListener); ok {
		listenerT.ExitBlock(s)
	}
}

func (s *BlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ArcParserVisitor:
		return t.VisitBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ArcParser) Block() (localctx IBlockContext) {
	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, ArcParserRULE_block)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(274)
		p.Match(ArcParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(278)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&183522612528299416) != 0) || ((int64((_la-66)) & ^0x3f) == 0 && ((int64(1)<<(_la-66))&63) != 0) {
		{
			p.SetState(275)
			p.Statement()
		}

		p.SetState(280)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(281)
		p.Match(ArcParserRBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	VariableDecl() IVariableDeclContext
	ConstDecl() IConstDeclContext
	AssignmentStmt() IAssignmentStmtContext
	ExpressionStmt() IExpressionStmtContext
	ReturnStmt() IReturnStmtContext
	IfStmt() IIfStmtContext
	ForStmt() IForStmtContext
	BreakStmt() IBreakStmtContext
	ContinueStmt() IContinueStmtContext
	DeferStmt() IDeferStmtContext
	Block() IBlockContext

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ArcParserRULE_statement
	return p
}

func InitEmptyStatementContext(p *StatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ArcParserRULE_statement
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ArcParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) VariableDecl() IVariableDeclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariableDeclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVariableDeclContext)
}

func (s *StatementContext) ConstDecl() IConstDeclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConstDeclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConstDeclContext)
}

func (s *StatementContext) AssignmentStmt() IAssignmentStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignmentStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignmentStmtContext)
}

func (s *StatementContext) ExpressionStmt() IExpressionStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionStmtContext)
}

func (s *StatementContext) ReturnStmt() IReturnStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReturnStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReturnStmtContext)
}

func (s *StatementContext) IfStmt() IIfStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfStmtContext)
}

func (s *StatementContext) ForStmt() IForStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IForStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IForStmtContext)
}

func (s *StatementContext) BreakStmt() IBreakStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBreakStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBreakStmtContext)
}

func (s *StatementContext) ContinueStmt() IContinueStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IContinueStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IContinueStmtContext)
}

func (s *StatementContext) DeferStmt() IDeferStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeferStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeferStmtContext)
}

func (s *StatementContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ArcParserListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ArcParserListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (s *StatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ArcParserVisitor:
		return t.VisitStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ArcParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, ArcParserRULE_statement)
	p.SetState(294)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 22, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(283)
			p.VariableDecl()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(284)
			p.ConstDecl()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(285)
			p.AssignmentStmt()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(286)
			p.ExpressionStmt()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(287)
			p.ReturnStmt()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(288)
			p.IfStmt()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(289)
			p.ForStmt()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(290)
			p.BreakStmt()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(291)
			p.ContinueStmt()
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(292)
			p.DeferStmt()
		}

	case 11:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(293)
			p.Block()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAssignmentStmtContext is an interface to support dynamic dispatch.
type IAssignmentStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LeftHandSide() ILeftHandSideContext
	ASSIGN() antlr.TerminalNode
	Expression() IExpressionContext

	// IsAssignmentStmtContext differentiates from other interfaces.
	IsAssignmentStmtContext()
}

type AssignmentStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignmentStmtContext() *AssignmentStmtContext {
	var p = new(AssignmentStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ArcParserRULE_assignmentStmt
	return p
}

func InitEmptyAssignmentStmtContext(p *AssignmentStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ArcParserRULE_assignmentStmt
}

func (*AssignmentStmtContext) IsAssignmentStmtContext() {}

func NewAssignmentStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignmentStmtContext {
	var p = new(AssignmentStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ArcParserRULE_assignmentStmt

	return p
}

func (s *AssignmentStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignmentStmtContext) LeftHandSide() ILeftHandSideContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILeftHandSideContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILeftHandSideContext)
}

func (s *AssignmentStmtContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(ArcParserASSIGN, 0)
}

func (s *AssignmentStmtContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AssignmentStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignmentStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ArcParserListener); ok {
		listenerT.EnterAssignmentStmt(s)
	}
}

func (s *AssignmentStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ArcParserListener); ok {
		listenerT.ExitAssignmentStmt(s)
	}
}

func (s *AssignmentStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ArcParserVisitor:
		return t.VisitAssignmentStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ArcParser) AssignmentStmt() (localctx IAssignmentStmtContext) {
	localctx = NewAssignmentStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, ArcParserRULE_assignmentStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(296)
		p.LeftHandSide()
	}
	{
		p.SetState(297)
		p.Match(ArcParserASSIGN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(298)
		p.Expression()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILeftHandSideContext is an interface to support dynamic dispatch.
type ILeftHandSideContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	STAR() antlr.TerminalNode
	Expression() IExpressionContext
	DOT() antlr.TerminalNode

	// IsLeftHandSideContext differentiates from other interfaces.
	IsLeftHandSideContext()
}

type LeftHandSideContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLeftHandSideContext() *LeftHandSideContext {
	var p = new(LeftHandSideContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ArcParserRULE_leftHandSide
	return p
}

func InitEmptyLeftHandSideContext(p *LeftHandSideContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ArcParserRULE_leftHandSide
}

func (*LeftHandSideContext) IsLeftHandSideContext() {}

func NewLeftHandSideContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LeftHandSideContext {
	var p = new(LeftHandSideContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ArcParserRULE_leftHandSide

	return p
}

func (s *LeftHandSideContext) GetParser() antlr.Parser { return s.parser }

func (s *LeftHandSideContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ArcParserIDENTIFIER, 0)
}

func (s *LeftHandSideContext) STAR() antlr.TerminalNode {
	return s.GetToken(ArcParserSTAR, 0)
}

func (s *LeftHandSideContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *LeftHandSideContext) DOT() antlr.TerminalNode {
	return s.GetToken(ArcParserDOT, 0)
}

func (s *LeftHandSideContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LeftHandSideContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LeftHandSideContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ArcParserListener); ok {
		listenerT.EnterLeftHandSide(s)
	}
}

func (s *LeftHandSideContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ArcParserListener); ok {
		listenerT.ExitLeftHandSide(s)
	}
}

func (s *LeftHandSideContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ArcParserVisitor:
		return t.VisitLeftHandSide(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ArcParser) LeftHandSide() (localctx ILeftHandSideContext) {
	localctx = NewLeftHandSideContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, ArcParserRULE_leftHandSide)
	p.SetState(307)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 23, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(300)
			p.Match(ArcParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(301)
			p.Match(ArcParserSTAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(302)
			p.Expression()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(303)
			p.Expression()
		}
		{
			p.SetState(304)
			p.Match(ArcParserDOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(305)
			p.Match(ArcParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExpressionStmtContext is an interface to support dynamic dispatch.
type IExpressionStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression() IExpressionContext

	// IsExpressionStmtContext differentiates from other interfaces.
	IsExpressionStmtContext()
}

type ExpressionStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionStmtContext() *ExpressionStmtContext {
	var p = new(ExpressionStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ArcParserRULE_expressionStmt
	return p
}

func InitEmptyExpressionStmtContext(p *ExpressionStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ArcParserRULE_expressionStmt
}

func (*ExpressionStmtContext) IsExpressionStmtContext() {}

func NewExpressionStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionStmtContext {
	var p = new(ExpressionStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ArcParserRULE_expressionStmt

	return p
}

func (s *ExpressionStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionStmtContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpressionStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ArcParserListener); ok {
		listenerT.EnterExpressionStmt(s)
	}
}

func (s *ExpressionStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ArcParserListener); ok {
		listenerT.ExitExpressionStmt(s)
	}
}

func (s *ExpressionStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ArcParserVisitor:
		return t.VisitExpressionStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ArcParser) ExpressionStmt() (localctx IExpressionStmtContext) {
	localctx = NewExpressionStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, ArcParserRULE_expressionStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(309)
		p.Expression()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IReturnStmtContext is an interface to support dynamic dispatch.
type IReturnStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	RETURN() antlr.TerminalNode
	Expression() IExpressionContext

	// IsReturnStmtContext differentiates from other interfaces.
	IsReturnStmtContext()
}

type ReturnStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReturnStmtContext() *ReturnStmtContext {
	var p = new(ReturnStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ArcParserRULE_returnStmt
	return p
}

func InitEmptyReturnStmtContext(p *ReturnStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ArcParserRULE_returnStmt
}

func (*ReturnStmtContext) IsReturnStmtContext() {}

func NewReturnStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReturnStmtContext {
	var p = new(ReturnStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ArcParserRULE_returnStmt

	return p
}

func (s *ReturnStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *ReturnStmtContext) RETURN() antlr.TerminalNode {
	return s.GetToken(ArcParserRETURN, 0)
}

func (s *ReturnStmtContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ReturnStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturnStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReturnStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ArcParserListener); ok {
		listenerT.EnterReturnStmt(s)
	}
}

func (s *ReturnStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ArcParserListener); ok {
		listenerT.ExitReturnStmt(s)
	}
}

func (s *ReturnStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ArcParserVisitor:
		return t.VisitReturnStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ArcParser) ReturnStmt() (localctx IReturnStmtContext) {
	localctx = NewReturnStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, ArcParserRULE_returnStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(311)
		p.Match(ArcParserRETURN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(313)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 24, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(312)
			p.Expression()
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIfStmtContext is an interface to support dynamic dispatch.
type IIfStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllIF() []antlr.TerminalNode
	IF(i int) antlr.TerminalNode
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	AllBlock() []IBlockContext
	Block(i int) IBlockContext
	AllELSE() []antlr.TerminalNode
	ELSE(i int) antlr.TerminalNode

	// IsIfStmtContext differentiates from other interfaces.
	IsIfStmtContext()
}

type IfStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfStmtContext() *IfStmtContext {
	var p = new(IfStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ArcParserRULE_ifStmt
	return p
}

func InitEmptyIfStmtContext(p *IfStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ArcParserRULE_ifStmt
}

func (*IfStmtContext) IsIfStmtContext() {}

func NewIfStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfStmtContext {
	var p = new(IfStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ArcParserRULE_ifStmt

	return p
}

func (s *IfStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *IfStmtContext) AllIF() []antlr.TerminalNode {
	return s.GetTokens(ArcParserIF)
}

func (s *IfStmtContext) IF(i int) antlr.TerminalNode {
	return s.GetToken(ArcParserIF, i)
}

func (s *IfStmtContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *IfStmtContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *IfStmtContext) AllBlock() []IBlockContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBlockContext); ok {
			len++
		}
	}

	tst := make([]IBlockContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBlockContext); ok {
			tst[i] = t.(IBlockContext)
			i++
		}
	}

	return tst
}

func (s *IfStmtContext) Block(i int) IBlockContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *IfStmtContext) AllELSE() []antlr.TerminalNode {
	return s.GetTokens(ArcParserELSE)
}

func (s *IfStmtContext) ELSE(i int) antlr.TerminalNode {
	return s.GetToken(ArcParserELSE, i)
}

func (s *IfStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ArcParserListener); ok {
		listenerT.EnterIfStmt(s)
	}
}

func (s *IfStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ArcParserListener); ok {
		listenerT.ExitIfStmt(s)
	}
}

func (s *IfStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ArcParserVisitor:
		return t.VisitIfStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ArcParser) IfStmt() (localctx IIfStmtContext) {
	localctx = NewIfStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, ArcParserRULE_ifStmt)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(315)
		p.Match(ArcParserIF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(316)
		p.Expression()
	}
	{
		p.SetState(317)
		p.Block()
	}
	p.SetState(325)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 25, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(318)
				p.Match(ArcParserELSE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(319)
				p.Match(ArcParserIF)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(320)
				p.Expression()
			}
			{
				p.SetState(321)
				p.Block()
			}

		}
		p.SetState(327)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 25, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(330)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ArcParserELSE {
		{
			p.SetState(328)
			p.Match(ArcParserELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(329)
			p.Block()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IForStmtContext is an interface to support dynamic dispatch.
type IForStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FOR() antlr.TerminalNode
	Block() IBlockContext
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	AllSEMICOLON() []antlr.TerminalNode
	SEMICOLON(i int) antlr.TerminalNode
	VariableDecl() IVariableDeclContext
	AllAssignmentStmt() []IAssignmentStmtContext
	AssignmentStmt(i int) IAssignmentStmtContext

	// IsForStmtContext differentiates from other interfaces.
	IsForStmtContext()
}

type ForStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyForStmtContext() *ForStmtContext {
	var p = new(ForStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ArcParserRULE_forStmt
	return p
}

func InitEmptyForStmtContext(p *ForStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ArcParserRULE_forStmt
}

func (*ForStmtContext) IsForStmtContext() {}

func NewForStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ForStmtContext {
	var p = new(ForStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ArcParserRULE_forStmt

	return p
}

func (s *ForStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *ForStmtContext) FOR() antlr.TerminalNode {
	return s.GetToken(ArcParserFOR, 0)
}

func (s *ForStmtContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *ForStmtContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ForStmtContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ForStmtContext) AllSEMICOLON() []antlr.TerminalNode {
	return s.GetTokens(ArcParserSEMICOLON)
}

func (s *ForStmtContext) SEMICOLON(i int) antlr.TerminalNode {
	return s.GetToken(ArcParserSEMICOLON, i)
}

func (s *ForStmtContext) VariableDecl() IVariableDeclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariableDeclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVariableDeclContext)
}

func (s *ForStmtContext) AllAssignmentStmt() []IAssignmentStmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAssignmentStmtContext); ok {
			len++
		}
	}

	tst := make([]IAssignmentStmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAssignmentStmtContext); ok {
			tst[i] = t.(IAssignmentStmtContext)
			i++
		}
	}

	return tst
}

func (s *ForStmtContext) AssignmentStmt(i int) IAssignmentStmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignmentStmtContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignmentStmtContext)
}

func (s *ForStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ForStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ArcParserListener); ok {
		listenerT.EnterForStmt(s)
	}
}

func (s *ForStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ArcParserListener); ok {
		listenerT.ExitForStmt(s)
	}
}

func (s *ForStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ArcParserVisitor:
		return t.VisitForStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ArcParser) ForStmt() (localctx IForStmtContext) {
	localctx = NewForStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, ArcParserRULE_forStmt)
	var _la int

	p.SetState(353)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 30, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(332)
			p.Match(ArcParserFOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(333)
			p.Block()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(334)
			p.Match(ArcParserFOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(335)
			p.Expression()
		}
		{
			p.SetState(336)
			p.Block()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(338)
			p.Match(ArcParserFOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(341)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		switch p.GetTokenStream().LA(1) {
		case ArcParserLET:
			{
				p.SetState(339)
				p.VariableDecl()
			}

		case ArcParserALLOCA, ArcParserCAST, ArcParserMINUS, ArcParserSTAR, ArcParserNOT, ArcParserAMP, ArcParserLPAREN, ArcParserLBRACE, ArcParserBOOLEAN_LITERAL, ArcParserINTEGER_LITERAL, ArcParserFLOAT_LITERAL, ArcParserSTRING_LITERAL, ArcParserCHAR_LITERAL, ArcParserIDENTIFIER:
			{
				p.SetState(340)
				p.AssignmentStmt()
			}

		case ArcParserSEMICOLON:

		default:
		}
		{
			p.SetState(343)
			p.Match(ArcParserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(345)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64((_la-35)) & ^0x3f) == 0 && ((int64(1)<<(_la-35))&135296811035) != 0 {
			{
				p.SetState(344)
				p.Expression()
			}

		}
		{
			p.SetState(347)
			p.Match(ArcParserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(350)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 29, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(348)
				p.AssignmentStmt()
			}

		} else if p.HasError() { // JIM
			goto errorExit
		} else if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 29, p.GetParserRuleContext()) == 2 {
			{
				p.SetState(349)
				p.Expression()
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}
		{
			p.SetState(352)
			p.Block()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBreakStmtContext is an interface to support dynamic dispatch.
type IBreakStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	BREAK() antlr.TerminalNode

	// IsBreakStmtContext differentiates from other interfaces.
	IsBreakStmtContext()
}

type BreakStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBreakStmtContext() *BreakStmtContext {
	var p = new(BreakStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ArcParserRULE_breakStmt
	return p
}

func InitEmptyBreakStmtContext(p *BreakStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ArcParserRULE_breakStmt
}

func (*BreakStmtContext) IsBreakStmtContext() {}

func NewBreakStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BreakStmtContext {
	var p = new(BreakStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ArcParserRULE_breakStmt

	return p
}

func (s *BreakStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *BreakStmtContext) BREAK() antlr.TerminalNode {
	return s.GetToken(ArcParserBREAK, 0)
}

func (s *BreakStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BreakStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BreakStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ArcParserListener); ok {
		listenerT.EnterBreakStmt(s)
	}
}

func (s *BreakStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ArcParserListener); ok {
		listenerT.ExitBreakStmt(s)
	}
}

func (s *BreakStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ArcParserVisitor:
		return t.VisitBreakStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ArcParser) BreakStmt() (localctx IBreakStmtContext) {
	localctx = NewBreakStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, ArcParserRULE_breakStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(355)
		p.Match(ArcParserBREAK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IContinueStmtContext is an interface to support dynamic dispatch.
type IContinueStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CONTINUE() antlr.TerminalNode

	// IsContinueStmtContext differentiates from other interfaces.
	IsContinueStmtContext()
}

type ContinueStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyContinueStmtContext() *ContinueStmtContext {
	var p = new(ContinueStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ArcParserRULE_continueStmt
	return p
}

func InitEmptyContinueStmtContext(p *ContinueStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ArcParserRULE_continueStmt
}

func (*ContinueStmtContext) IsContinueStmtContext() {}

func NewContinueStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ContinueStmtContext {
	var p = new(ContinueStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ArcParserRULE_continueStmt

	return p
}

func (s *ContinueStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *ContinueStmtContext) CONTINUE() antlr.TerminalNode {
	return s.GetToken(ArcParserCONTINUE, 0)
}

func (s *ContinueStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ContinueStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ContinueStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ArcParserListener); ok {
		listenerT.EnterContinueStmt(s)
	}
}

func (s *ContinueStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ArcParserListener); ok {
		listenerT.ExitContinueStmt(s)
	}
}

func (s *ContinueStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ArcParserVisitor:
		return t.VisitContinueStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ArcParser) ContinueStmt() (localctx IContinueStmtContext) {
	localctx = NewContinueStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, ArcParserRULE_continueStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(357)
		p.Match(ArcParserCONTINUE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDeferStmtContext is an interface to support dynamic dispatch.
type IDeferStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DEFER() antlr.TerminalNode
	AssignmentStmt() IAssignmentStmtContext
	Expression() IExpressionContext

	// IsDeferStmtContext differentiates from other interfaces.
	IsDeferStmtContext()
}

type DeferStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeferStmtContext() *DeferStmtContext {
	var p = new(DeferStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ArcParserRULE_deferStmt
	return p
}

func InitEmptyDeferStmtContext(p *DeferStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ArcParserRULE_deferStmt
}

func (*DeferStmtContext) IsDeferStmtContext() {}

func NewDeferStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeferStmtContext {
	var p = new(DeferStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ArcParserRULE_deferStmt

	return p
}

func (s *DeferStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *DeferStmtContext) DEFER() antlr.TerminalNode {
	return s.GetToken(ArcParserDEFER, 0)
}

func (s *DeferStmtContext) AssignmentStmt() IAssignmentStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignmentStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignmentStmtContext)
}

func (s *DeferStmtContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *DeferStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeferStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeferStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ArcParserListener); ok {
		listenerT.EnterDeferStmt(s)
	}
}

func (s *DeferStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ArcParserListener); ok {
		listenerT.ExitDeferStmt(s)
	}
}

func (s *DeferStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ArcParserVisitor:
		return t.VisitDeferStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ArcParser) DeferStmt() (localctx IDeferStmtContext) {
	localctx = NewDeferStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, ArcParserRULE_deferStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(359)
		p.Match(ArcParserDEFER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(362)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 31, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(360)
			p.AssignmentStmt()
		}

	case 2:
		{
			p.SetState(361)
			p.Expression()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LogicalOrExpression() ILogicalOrExpressionContext

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ArcParserRULE_expression
	return p
}

func InitEmptyExpressionContext(p *ExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ArcParserRULE_expression
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ArcParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) LogicalOrExpression() ILogicalOrExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILogicalOrExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILogicalOrExpressionContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ArcParserListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ArcParserListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (s *ExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ArcParserVisitor:
		return t.VisitExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ArcParser) Expression() (localctx IExpressionContext) {
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, ArcParserRULE_expression)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(364)
		p.LogicalOrExpression()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILogicalOrExpressionContext is an interface to support dynamic dispatch.
type ILogicalOrExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllLogicalAndExpression() []ILogicalAndExpressionContext
	LogicalAndExpression(i int) ILogicalAndExpressionContext
	AllOR() []antlr.TerminalNode
	OR(i int) antlr.TerminalNode

	// IsLogicalOrExpressionContext differentiates from other interfaces.
	IsLogicalOrExpressionContext()
}

type LogicalOrExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLogicalOrExpressionContext() *LogicalOrExpressionContext {
	var p = new(LogicalOrExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ArcParserRULE_logicalOrExpression
	return p
}

func InitEmptyLogicalOrExpressionContext(p *LogicalOrExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ArcParserRULE_logicalOrExpression
}

func (*LogicalOrExpressionContext) IsLogicalOrExpressionContext() {}

func NewLogicalOrExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LogicalOrExpressionContext {
	var p = new(LogicalOrExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ArcParserRULE_logicalOrExpression

	return p
}

func (s *LogicalOrExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *LogicalOrExpressionContext) AllLogicalAndExpression() []ILogicalAndExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ILogicalAndExpressionContext); ok {
			len++
		}
	}

	tst := make([]ILogicalAndExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ILogicalAndExpressionContext); ok {
			tst[i] = t.(ILogicalAndExpressionContext)
			i++
		}
	}

	return tst
}

func (s *LogicalOrExpressionContext) LogicalAndExpression(i int) ILogicalAndExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILogicalAndExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILogicalAndExpressionContext)
}

func (s *LogicalOrExpressionContext) AllOR() []antlr.TerminalNode {
	return s.GetTokens(ArcParserOR)
}

func (s *LogicalOrExpressionContext) OR(i int) antlr.TerminalNode {
	return s.GetToken(ArcParserOR, i)
}

func (s *LogicalOrExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LogicalOrExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LogicalOrExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ArcParserListener); ok {
		listenerT.EnterLogicalOrExpression(s)
	}
}

func (s *LogicalOrExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ArcParserListener); ok {
		listenerT.ExitLogicalOrExpression(s)
	}
}

func (s *LogicalOrExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ArcParserVisitor:
		return t.VisitLogicalOrExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ArcParser) LogicalOrExpression() (localctx ILogicalOrExpressionContext) {
	localctx = NewLogicalOrExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, ArcParserRULE_logicalOrExpression)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(366)
		p.LogicalAndExpression()
	}
	p.SetState(371)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ArcParserOR {
		{
			p.SetState(367)
			p.Match(ArcParserOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(368)
			p.LogicalAndExpression()
		}

		p.SetState(373)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILogicalAndExpressionContext is an interface to support dynamic dispatch.
type ILogicalAndExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllEqualityExpression() []IEqualityExpressionContext
	EqualityExpression(i int) IEqualityExpressionContext
	AllAND() []antlr.TerminalNode
	AND(i int) antlr.TerminalNode

	// IsLogicalAndExpressionContext differentiates from other interfaces.
	IsLogicalAndExpressionContext()
}

type LogicalAndExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLogicalAndExpressionContext() *LogicalAndExpressionContext {
	var p = new(LogicalAndExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ArcParserRULE_logicalAndExpression
	return p
}

func InitEmptyLogicalAndExpressionContext(p *LogicalAndExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ArcParserRULE_logicalAndExpression
}

func (*LogicalAndExpressionContext) IsLogicalAndExpressionContext() {}

func NewLogicalAndExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LogicalAndExpressionContext {
	var p = new(LogicalAndExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ArcParserRULE_logicalAndExpression

	return p
}

func (s *LogicalAndExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *LogicalAndExpressionContext) AllEqualityExpression() []IEqualityExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IEqualityExpressionContext); ok {
			len++
		}
	}

	tst := make([]IEqualityExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IEqualityExpressionContext); ok {
			tst[i] = t.(IEqualityExpressionContext)
			i++
		}
	}

	return tst
}

func (s *LogicalAndExpressionContext) EqualityExpression(i int) IEqualityExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEqualityExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEqualityExpressionContext)
}

func (s *LogicalAndExpressionContext) AllAND() []antlr.TerminalNode {
	return s.GetTokens(ArcParserAND)
}

func (s *LogicalAndExpressionContext) AND(i int) antlr.TerminalNode {
	return s.GetToken(ArcParserAND, i)
}

func (s *LogicalAndExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LogicalAndExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LogicalAndExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ArcParserListener); ok {
		listenerT.EnterLogicalAndExpression(s)
	}
}

func (s *LogicalAndExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ArcParserListener); ok {
		listenerT.ExitLogicalAndExpression(s)
	}
}

func (s *LogicalAndExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ArcParserVisitor:
		return t.VisitLogicalAndExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ArcParser) LogicalAndExpression() (localctx ILogicalAndExpressionContext) {
	localctx = NewLogicalAndExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, ArcParserRULE_logicalAndExpression)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(374)
		p.EqualityExpression()
	}
	p.SetState(379)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ArcParserAND {
		{
			p.SetState(375)
			p.Match(ArcParserAND)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(376)
			p.EqualityExpression()
		}

		p.SetState(381)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IEqualityExpressionContext is an interface to support dynamic dispatch.
type IEqualityExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllRelationalExpression() []IRelationalExpressionContext
	RelationalExpression(i int) IRelationalExpressionContext
	AllEQ() []antlr.TerminalNode
	EQ(i int) antlr.TerminalNode
	AllNE() []antlr.TerminalNode
	NE(i int) antlr.TerminalNode

	// IsEqualityExpressionContext differentiates from other interfaces.
	IsEqualityExpressionContext()
}

type EqualityExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEqualityExpressionContext() *EqualityExpressionContext {
	var p = new(EqualityExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ArcParserRULE_equalityExpression
	return p
}

func InitEmptyEqualityExpressionContext(p *EqualityExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ArcParserRULE_equalityExpression
}

func (*EqualityExpressionContext) IsEqualityExpressionContext() {}

func NewEqualityExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EqualityExpressionContext {
	var p = new(EqualityExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ArcParserRULE_equalityExpression

	return p
}

func (s *EqualityExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *EqualityExpressionContext) AllRelationalExpression() []IRelationalExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IRelationalExpressionContext); ok {
			len++
		}
	}

	tst := make([]IRelationalExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IRelationalExpressionContext); ok {
			tst[i] = t.(IRelationalExpressionContext)
			i++
		}
	}

	return tst
}

func (s *EqualityExpressionContext) RelationalExpression(i int) IRelationalExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRelationalExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRelationalExpressionContext)
}

func (s *EqualityExpressionContext) AllEQ() []antlr.TerminalNode {
	return s.GetTokens(ArcParserEQ)
}

func (s *EqualityExpressionContext) EQ(i int) antlr.TerminalNode {
	return s.GetToken(ArcParserEQ, i)
}

func (s *EqualityExpressionContext) AllNE() []antlr.TerminalNode {
	return s.GetTokens(ArcParserNE)
}

func (s *EqualityExpressionContext) NE(i int) antlr.TerminalNode {
	return s.GetToken(ArcParserNE, i)
}

func (s *EqualityExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EqualityExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EqualityExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ArcParserListener); ok {
		listenerT.EnterEqualityExpression(s)
	}
}

func (s *EqualityExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ArcParserListener); ok {
		listenerT.ExitEqualityExpression(s)
	}
}

func (s *EqualityExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ArcParserVisitor:
		return t.VisitEqualityExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ArcParser) EqualityExpression() (localctx IEqualityExpressionContext) {
	localctx = NewEqualityExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, ArcParserRULE_equalityExpression)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(382)
		p.RelationalExpression()
	}
	p.SetState(387)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ArcParserEQ || _la == ArcParserNE {
		{
			p.SetState(383)
			_la = p.GetTokenStream().LA(1)

			if !(_la == ArcParserEQ || _la == ArcParserNE) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(384)
			p.RelationalExpression()
		}

		p.SetState(389)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRelationalExpressionContext is an interface to support dynamic dispatch.
type IRelationalExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllAdditiveExpression() []IAdditiveExpressionContext
	AdditiveExpression(i int) IAdditiveExpressionContext
	AllLT() []antlr.TerminalNode
	LT(i int) antlr.TerminalNode
	AllLE() []antlr.TerminalNode
	LE(i int) antlr.TerminalNode
	AllGT() []antlr.TerminalNode
	GT(i int) antlr.TerminalNode
	AllGE() []antlr.TerminalNode
	GE(i int) antlr.TerminalNode

	// IsRelationalExpressionContext differentiates from other interfaces.
	IsRelationalExpressionContext()
}

type RelationalExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRelationalExpressionContext() *RelationalExpressionContext {
	var p = new(RelationalExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ArcParserRULE_relationalExpression
	return p
}

func InitEmptyRelationalExpressionContext(p *RelationalExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ArcParserRULE_relationalExpression
}

func (*RelationalExpressionContext) IsRelationalExpressionContext() {}

func NewRelationalExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RelationalExpressionContext {
	var p = new(RelationalExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ArcParserRULE_relationalExpression

	return p
}

func (s *RelationalExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *RelationalExpressionContext) AllAdditiveExpression() []IAdditiveExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAdditiveExpressionContext); ok {
			len++
		}
	}

	tst := make([]IAdditiveExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAdditiveExpressionContext); ok {
			tst[i] = t.(IAdditiveExpressionContext)
			i++
		}
	}

	return tst
}

func (s *RelationalExpressionContext) AdditiveExpression(i int) IAdditiveExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAdditiveExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAdditiveExpressionContext)
}

func (s *RelationalExpressionContext) AllLT() []antlr.TerminalNode {
	return s.GetTokens(ArcParserLT)
}

func (s *RelationalExpressionContext) LT(i int) antlr.TerminalNode {
	return s.GetToken(ArcParserLT, i)
}

func (s *RelationalExpressionContext) AllLE() []antlr.TerminalNode {
	return s.GetTokens(ArcParserLE)
}

func (s *RelationalExpressionContext) LE(i int) antlr.TerminalNode {
	return s.GetToken(ArcParserLE, i)
}

func (s *RelationalExpressionContext) AllGT() []antlr.TerminalNode {
	return s.GetTokens(ArcParserGT)
}

func (s *RelationalExpressionContext) GT(i int) antlr.TerminalNode {
	return s.GetToken(ArcParserGT, i)
}

func (s *RelationalExpressionContext) AllGE() []antlr.TerminalNode {
	return s.GetTokens(ArcParserGE)
}

func (s *RelationalExpressionContext) GE(i int) antlr.TerminalNode {
	return s.GetToken(ArcParserGE, i)
}

func (s *RelationalExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RelationalExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RelationalExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ArcParserListener); ok {
		listenerT.EnterRelationalExpression(s)
	}
}

func (s *RelationalExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ArcParserListener); ok {
		listenerT.ExitRelationalExpression(s)
	}
}

func (s *RelationalExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ArcParserVisitor:
		return t.VisitRelationalExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ArcParser) RelationalExpression() (localctx IRelationalExpressionContext) {
	localctx = NewRelationalExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, ArcParserRULE_relationalExpression)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(390)
		p.AdditiveExpression()
	}
	p.SetState(395)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&263882790666240) != 0 {
		{
			p.SetState(391)
			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&263882790666240) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(392)
			p.AdditiveExpression()
		}

		p.SetState(397)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAdditiveExpressionContext is an interface to support dynamic dispatch.
type IAdditiveExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllMultiplicativeExpression() []IMultiplicativeExpressionContext
	MultiplicativeExpression(i int) IMultiplicativeExpressionContext
	AllPLUS() []antlr.TerminalNode
	PLUS(i int) antlr.TerminalNode
	AllMINUS() []antlr.TerminalNode
	MINUS(i int) antlr.TerminalNode

	// IsAdditiveExpressionContext differentiates from other interfaces.
	IsAdditiveExpressionContext()
}

type AdditiveExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAdditiveExpressionContext() *AdditiveExpressionContext {
	var p = new(AdditiveExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ArcParserRULE_additiveExpression
	return p
}

func InitEmptyAdditiveExpressionContext(p *AdditiveExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ArcParserRULE_additiveExpression
}

func (*AdditiveExpressionContext) IsAdditiveExpressionContext() {}

func NewAdditiveExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AdditiveExpressionContext {
	var p = new(AdditiveExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ArcParserRULE_additiveExpression

	return p
}

func (s *AdditiveExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *AdditiveExpressionContext) AllMultiplicativeExpression() []IMultiplicativeExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IMultiplicativeExpressionContext); ok {
			len++
		}
	}

	tst := make([]IMultiplicativeExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IMultiplicativeExpressionContext); ok {
			tst[i] = t.(IMultiplicativeExpressionContext)
			i++
		}
	}

	return tst
}

func (s *AdditiveExpressionContext) MultiplicativeExpression(i int) IMultiplicativeExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMultiplicativeExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMultiplicativeExpressionContext)
}

func (s *AdditiveExpressionContext) AllPLUS() []antlr.TerminalNode {
	return s.GetTokens(ArcParserPLUS)
}

func (s *AdditiveExpressionContext) PLUS(i int) antlr.TerminalNode {
	return s.GetToken(ArcParserPLUS, i)
}

func (s *AdditiveExpressionContext) AllMINUS() []antlr.TerminalNode {
	return s.GetTokens(ArcParserMINUS)
}

func (s *AdditiveExpressionContext) MINUS(i int) antlr.TerminalNode {
	return s.GetToken(ArcParserMINUS, i)
}

func (s *AdditiveExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AdditiveExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AdditiveExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ArcParserListener); ok {
		listenerT.EnterAdditiveExpression(s)
	}
}

func (s *AdditiveExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ArcParserListener); ok {
		listenerT.ExitAdditiveExpression(s)
	}
}

func (s *AdditiveExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ArcParserVisitor:
		return t.VisitAdditiveExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ArcParser) AdditiveExpression() (localctx IAdditiveExpressionContext) {
	localctx = NewAdditiveExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, ArcParserRULE_additiveExpression)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(398)
		p.MultiplicativeExpression()
	}
	p.SetState(403)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 36, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(399)
				_la = p.GetTokenStream().LA(1)

				if !(_la == ArcParserPLUS || _la == ArcParserMINUS) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}
			{
				p.SetState(400)
				p.MultiplicativeExpression()
			}

		}
		p.SetState(405)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 36, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMultiplicativeExpressionContext is an interface to support dynamic dispatch.
type IMultiplicativeExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllUnaryExpression() []IUnaryExpressionContext
	UnaryExpression(i int) IUnaryExpressionContext
	AllSTAR() []antlr.TerminalNode
	STAR(i int) antlr.TerminalNode
	AllSLASH() []antlr.TerminalNode
	SLASH(i int) antlr.TerminalNode
	AllPERCENT() []antlr.TerminalNode
	PERCENT(i int) antlr.TerminalNode

	// IsMultiplicativeExpressionContext differentiates from other interfaces.
	IsMultiplicativeExpressionContext()
}

type MultiplicativeExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMultiplicativeExpressionContext() *MultiplicativeExpressionContext {
	var p = new(MultiplicativeExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ArcParserRULE_multiplicativeExpression
	return p
}

func InitEmptyMultiplicativeExpressionContext(p *MultiplicativeExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ArcParserRULE_multiplicativeExpression
}

func (*MultiplicativeExpressionContext) IsMultiplicativeExpressionContext() {}

func NewMultiplicativeExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MultiplicativeExpressionContext {
	var p = new(MultiplicativeExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ArcParserRULE_multiplicativeExpression

	return p
}

func (s *MultiplicativeExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *MultiplicativeExpressionContext) AllUnaryExpression() []IUnaryExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IUnaryExpressionContext); ok {
			len++
		}
	}

	tst := make([]IUnaryExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IUnaryExpressionContext); ok {
			tst[i] = t.(IUnaryExpressionContext)
			i++
		}
	}

	return tst
}

func (s *MultiplicativeExpressionContext) UnaryExpression(i int) IUnaryExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnaryExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnaryExpressionContext)
}

func (s *MultiplicativeExpressionContext) AllSTAR() []antlr.TerminalNode {
	return s.GetTokens(ArcParserSTAR)
}

func (s *MultiplicativeExpressionContext) STAR(i int) antlr.TerminalNode {
	return s.GetToken(ArcParserSTAR, i)
}

func (s *MultiplicativeExpressionContext) AllSLASH() []antlr.TerminalNode {
	return s.GetTokens(ArcParserSLASH)
}

func (s *MultiplicativeExpressionContext) SLASH(i int) antlr.TerminalNode {
	return s.GetToken(ArcParserSLASH, i)
}

func (s *MultiplicativeExpressionContext) AllPERCENT() []antlr.TerminalNode {
	return s.GetTokens(ArcParserPERCENT)
}

func (s *MultiplicativeExpressionContext) PERCENT(i int) antlr.TerminalNode {
	return s.GetToken(ArcParserPERCENT, i)
}

func (s *MultiplicativeExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultiplicativeExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MultiplicativeExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ArcParserListener); ok {
		listenerT.EnterMultiplicativeExpression(s)
	}
}

func (s *MultiplicativeExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ArcParserListener); ok {
		listenerT.ExitMultiplicativeExpression(s)
	}
}

func (s *MultiplicativeExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ArcParserVisitor:
		return t.VisitMultiplicativeExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ArcParser) MultiplicativeExpression() (localctx IMultiplicativeExpressionContext) {
	localctx = NewMultiplicativeExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, ArcParserRULE_multiplicativeExpression)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(406)
		p.UnaryExpression()
	}
	p.SetState(411)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 37, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(407)
				_la = p.GetTokenStream().LA(1)

				if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&3848290697216) != 0) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}
			{
				p.SetState(408)
				p.UnaryExpression()
			}

		}
		p.SetState(413)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 37, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IUnaryExpressionContext is an interface to support dynamic dispatch.
type IUnaryExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	UnaryExpression() IUnaryExpressionContext
	MINUS() antlr.TerminalNode
	NOT() antlr.TerminalNode
	STAR() antlr.TerminalNode
	AMP() antlr.TerminalNode
	PostfixExpression() IPostfixExpressionContext

	// IsUnaryExpressionContext differentiates from other interfaces.
	IsUnaryExpressionContext()
}

type UnaryExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnaryExpressionContext() *UnaryExpressionContext {
	var p = new(UnaryExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ArcParserRULE_unaryExpression
	return p
}

func InitEmptyUnaryExpressionContext(p *UnaryExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ArcParserRULE_unaryExpression
}

func (*UnaryExpressionContext) IsUnaryExpressionContext() {}

func NewUnaryExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnaryExpressionContext {
	var p = new(UnaryExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ArcParserRULE_unaryExpression

	return p
}

func (s *UnaryExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *UnaryExpressionContext) UnaryExpression() IUnaryExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnaryExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnaryExpressionContext)
}

func (s *UnaryExpressionContext) MINUS() antlr.TerminalNode {
	return s.GetToken(ArcParserMINUS, 0)
}

func (s *UnaryExpressionContext) NOT() antlr.TerminalNode {
	return s.GetToken(ArcParserNOT, 0)
}

func (s *UnaryExpressionContext) STAR() antlr.TerminalNode {
	return s.GetToken(ArcParserSTAR, 0)
}

func (s *UnaryExpressionContext) AMP() antlr.TerminalNode {
	return s.GetToken(ArcParserAMP, 0)
}

func (s *UnaryExpressionContext) PostfixExpression() IPostfixExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPostfixExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPostfixExpressionContext)
}

func (s *UnaryExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnaryExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ArcParserListener); ok {
		listenerT.EnterUnaryExpression(s)
	}
}

func (s *UnaryExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ArcParserListener); ok {
		listenerT.ExitUnaryExpression(s)
	}
}

func (s *UnaryExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ArcParserVisitor:
		return t.VisitUnaryExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ArcParser) UnaryExpression() (localctx IUnaryExpressionContext) {
	localctx = NewUnaryExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, ArcParserRULE_unaryExpression)
	var _la int

	p.SetState(417)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ArcParserMINUS, ArcParserSTAR, ArcParserNOT, ArcParserAMP:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(414)
			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&3378524354248704) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(415)
			p.UnaryExpression()
		}

	case ArcParserALLOCA, ArcParserCAST, ArcParserLPAREN, ArcParserLBRACE, ArcParserBOOLEAN_LITERAL, ArcParserINTEGER_LITERAL, ArcParserFLOAT_LITERAL, ArcParserSTRING_LITERAL, ArcParserCHAR_LITERAL, ArcParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(416)
			p.PostfixExpression()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPostfixExpressionContext is an interface to support dynamic dispatch.
type IPostfixExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PrimaryExpression() IPrimaryExpressionContext
	AllPostfixOp() []IPostfixOpContext
	PostfixOp(i int) IPostfixOpContext

	// IsPostfixExpressionContext differentiates from other interfaces.
	IsPostfixExpressionContext()
}

type PostfixExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPostfixExpressionContext() *PostfixExpressionContext {
	var p = new(PostfixExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ArcParserRULE_postfixExpression
	return p
}

func InitEmptyPostfixExpressionContext(p *PostfixExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ArcParserRULE_postfixExpression
}

func (*PostfixExpressionContext) IsPostfixExpressionContext() {}

func NewPostfixExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PostfixExpressionContext {
	var p = new(PostfixExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ArcParserRULE_postfixExpression

	return p
}

func (s *PostfixExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *PostfixExpressionContext) PrimaryExpression() IPrimaryExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimaryExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimaryExpressionContext)
}

func (s *PostfixExpressionContext) AllPostfixOp() []IPostfixOpContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IPostfixOpContext); ok {
			len++
		}
	}

	tst := make([]IPostfixOpContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IPostfixOpContext); ok {
			tst[i] = t.(IPostfixOpContext)
			i++
		}
	}

	return tst
}

func (s *PostfixExpressionContext) PostfixOp(i int) IPostfixOpContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPostfixOpContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPostfixOpContext)
}

func (s *PostfixExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PostfixExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PostfixExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ArcParserListener); ok {
		listenerT.EnterPostfixExpression(s)
	}
}

func (s *PostfixExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ArcParserListener); ok {
		listenerT.ExitPostfixExpression(s)
	}
}

func (s *PostfixExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ArcParserVisitor:
		return t.VisitPostfixExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ArcParser) PostfixExpression() (localctx IPostfixExpressionContext) {
	localctx = NewPostfixExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 82, ArcParserRULE_postfixExpression)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(419)
		p.PrimaryExpression()
	}
	p.SetState(423)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 39, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(420)
				p.PostfixOp()
			}

		}
		p.SetState(425)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 39, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPostfixOpContext is an interface to support dynamic dispatch.
type IPostfixOpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DOT() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	ArgumentList() IArgumentListContext

	// IsPostfixOpContext differentiates from other interfaces.
	IsPostfixOpContext()
}

type PostfixOpContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPostfixOpContext() *PostfixOpContext {
	var p = new(PostfixOpContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ArcParserRULE_postfixOp
	return p
}

func InitEmptyPostfixOpContext(p *PostfixOpContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ArcParserRULE_postfixOp
}

func (*PostfixOpContext) IsPostfixOpContext() {}

func NewPostfixOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PostfixOpContext {
	var p = new(PostfixOpContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ArcParserRULE_postfixOp

	return p
}

func (s *PostfixOpContext) GetParser() antlr.Parser { return s.parser }

func (s *PostfixOpContext) DOT() antlr.TerminalNode {
	return s.GetToken(ArcParserDOT, 0)
}

func (s *PostfixOpContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ArcParserIDENTIFIER, 0)
}

func (s *PostfixOpContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(ArcParserLPAREN, 0)
}

func (s *PostfixOpContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(ArcParserRPAREN, 0)
}

func (s *PostfixOpContext) ArgumentList() IArgumentListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgumentListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgumentListContext)
}

func (s *PostfixOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PostfixOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PostfixOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ArcParserListener); ok {
		listenerT.EnterPostfixOp(s)
	}
}

func (s *PostfixOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ArcParserListener); ok {
		listenerT.ExitPostfixOp(s)
	}
}

func (s *PostfixOpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ArcParserVisitor:
		return t.VisitPostfixOp(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ArcParser) PostfixOp() (localctx IPostfixOpContext) {
	localctx = NewPostfixOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 84, ArcParserRULE_postfixOp)
	var _la int

	p.SetState(440)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 42, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(426)
			p.Match(ArcParserDOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(427)
			p.Match(ArcParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(428)
			p.Match(ArcParserDOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(429)
			p.Match(ArcParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(430)
			p.Match(ArcParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(432)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64((_la-35)) & ^0x3f) == 0 && ((int64(1)<<(_la-35))&135296811035) != 0 {
			{
				p.SetState(431)
				p.ArgumentList()
			}

		}
		{
			p.SetState(434)
			p.Match(ArcParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(435)
			p.Match(ArcParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(437)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64((_la-35)) & ^0x3f) == 0 && ((int64(1)<<(_la-35))&135296811035) != 0 {
			{
				p.SetState(436)
				p.ArgumentList()
			}

		}
		{
			p.SetState(439)
			p.Match(ArcParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPrimaryExpressionContext is an interface to support dynamic dispatch.
type IPrimaryExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Literal() ILiteralContext
	IDENTIFIER() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	Expression() IExpressionContext
	RPAREN() antlr.TerminalNode
	CastExpression() ICastExpressionContext
	AllocaExpression() IAllocaExpressionContext
	StructLiteral() IStructLiteralContext

	// IsPrimaryExpressionContext differentiates from other interfaces.
	IsPrimaryExpressionContext()
}

type PrimaryExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrimaryExpressionContext() *PrimaryExpressionContext {
	var p = new(PrimaryExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ArcParserRULE_primaryExpression
	return p
}

func InitEmptyPrimaryExpressionContext(p *PrimaryExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ArcParserRULE_primaryExpression
}

func (*PrimaryExpressionContext) IsPrimaryExpressionContext() {}

func NewPrimaryExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimaryExpressionContext {
	var p = new(PrimaryExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ArcParserRULE_primaryExpression

	return p
}

func (s *PrimaryExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *PrimaryExpressionContext) Literal() ILiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *PrimaryExpressionContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ArcParserIDENTIFIER, 0)
}

func (s *PrimaryExpressionContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(ArcParserLPAREN, 0)
}

func (s *PrimaryExpressionContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *PrimaryExpressionContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(ArcParserRPAREN, 0)
}

func (s *PrimaryExpressionContext) CastExpression() ICastExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICastExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICastExpressionContext)
}

func (s *PrimaryExpressionContext) AllocaExpression() IAllocaExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAllocaExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAllocaExpressionContext)
}

func (s *PrimaryExpressionContext) StructLiteral() IStructLiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStructLiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStructLiteralContext)
}

func (s *PrimaryExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimaryExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrimaryExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ArcParserListener); ok {
		listenerT.EnterPrimaryExpression(s)
	}
}

func (s *PrimaryExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ArcParserListener); ok {
		listenerT.ExitPrimaryExpression(s)
	}
}

func (s *PrimaryExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ArcParserVisitor:
		return t.VisitPrimaryExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ArcParser) PrimaryExpression() (localctx IPrimaryExpressionContext) {
	localctx = NewPrimaryExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 86, ArcParserRULE_primaryExpression)
	p.SetState(451)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 43, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(442)
			p.Literal()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(443)
			p.Match(ArcParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(444)
			p.Match(ArcParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(445)
			p.Expression()
		}
		{
			p.SetState(446)
			p.Match(ArcParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(448)
			p.CastExpression()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(449)
			p.AllocaExpression()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(450)
			p.StructLiteral()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILiteralContext is an interface to support dynamic dispatch.
type ILiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INTEGER_LITERAL() antlr.TerminalNode
	FLOAT_LITERAL() antlr.TerminalNode
	STRING_LITERAL() antlr.TerminalNode
	CHAR_LITERAL() antlr.TerminalNode
	BOOLEAN_LITERAL() antlr.TerminalNode
	VectorLiteral() IVectorLiteralContext
	MapLiteral() IMapLiteralContext

	// IsLiteralContext differentiates from other interfaces.
	IsLiteralContext()
}

type LiteralContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteralContext() *LiteralContext {
	var p = new(LiteralContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ArcParserRULE_literal
	return p
}

func InitEmptyLiteralContext(p *LiteralContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ArcParserRULE_literal
}

func (*LiteralContext) IsLiteralContext() {}

func NewLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralContext {
	var p = new(LiteralContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ArcParserRULE_literal

	return p
}

func (s *LiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralContext) INTEGER_LITERAL() antlr.TerminalNode {
	return s.GetToken(ArcParserINTEGER_LITERAL, 0)
}

func (s *LiteralContext) FLOAT_LITERAL() antlr.TerminalNode {
	return s.GetToken(ArcParserFLOAT_LITERAL, 0)
}

func (s *LiteralContext) STRING_LITERAL() antlr.TerminalNode {
	return s.GetToken(ArcParserSTRING_LITERAL, 0)
}

func (s *LiteralContext) CHAR_LITERAL() antlr.TerminalNode {
	return s.GetToken(ArcParserCHAR_LITERAL, 0)
}

func (s *LiteralContext) BOOLEAN_LITERAL() antlr.TerminalNode {
	return s.GetToken(ArcParserBOOLEAN_LITERAL, 0)
}

func (s *LiteralContext) VectorLiteral() IVectorLiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVectorLiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVectorLiteralContext)
}

func (s *LiteralContext) MapLiteral() IMapLiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMapLiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMapLiteralContext)
}

func (s *LiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ArcParserListener); ok {
		listenerT.EnterLiteral(s)
	}
}

func (s *LiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ArcParserListener); ok {
		listenerT.ExitLiteral(s)
	}
}

func (s *LiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ArcParserVisitor:
		return t.VisitLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ArcParser) Literal() (localctx ILiteralContext) {
	localctx = NewLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 88, ArcParserRULE_literal)
	p.SetState(460)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 44, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(453)
			p.Match(ArcParserINTEGER_LITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(454)
			p.Match(ArcParserFLOAT_LITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(455)
			p.Match(ArcParserSTRING_LITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(456)
			p.Match(ArcParserCHAR_LITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(457)
			p.Match(ArcParserBOOLEAN_LITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(458)
			p.VectorLiteral()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(459)
			p.MapLiteral()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVectorLiteralContext is an interface to support dynamic dispatch.
type IVectorLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsVectorLiteralContext differentiates from other interfaces.
	IsVectorLiteralContext()
}

type VectorLiteralContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVectorLiteralContext() *VectorLiteralContext {
	var p = new(VectorLiteralContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ArcParserRULE_vectorLiteral
	return p
}

func InitEmptyVectorLiteralContext(p *VectorLiteralContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ArcParserRULE_vectorLiteral
}

func (*VectorLiteralContext) IsVectorLiteralContext() {}

func NewVectorLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VectorLiteralContext {
	var p = new(VectorLiteralContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ArcParserRULE_vectorLiteral

	return p
}

func (s *VectorLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *VectorLiteralContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(ArcParserLBRACE, 0)
}

func (s *VectorLiteralContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(ArcParserRBRACE, 0)
}

func (s *VectorLiteralContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *VectorLiteralContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *VectorLiteralContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(ArcParserCOMMA)
}

func (s *VectorLiteralContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(ArcParserCOMMA, i)
}

func (s *VectorLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VectorLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VectorLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ArcParserListener); ok {
		listenerT.EnterVectorLiteral(s)
	}
}

func (s *VectorLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ArcParserListener); ok {
		listenerT.ExitVectorLiteral(s)
	}
}

func (s *VectorLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ArcParserVisitor:
		return t.VisitVectorLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ArcParser) VectorLiteral() (localctx IVectorLiteralContext) {
	localctx = NewVectorLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 90, ArcParserRULE_vectorLiteral)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(462)
		p.Match(ArcParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(471)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64((_la-35)) & ^0x3f) == 0 && ((int64(1)<<(_la-35))&135296811035) != 0 {
		{
			p.SetState(463)
			p.Expression()
		}
		p.SetState(468)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == ArcParserCOMMA {
			{
				p.SetState(464)
				p.Match(ArcParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(465)
				p.Expression()
			}

			p.SetState(470)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(473)
		p.Match(ArcParserRBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMapLiteralContext is an interface to support dynamic dispatch.
type IMapLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	AllMapEntry() []IMapEntryContext
	MapEntry(i int) IMapEntryContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsMapLiteralContext differentiates from other interfaces.
	IsMapLiteralContext()
}

type MapLiteralContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMapLiteralContext() *MapLiteralContext {
	var p = new(MapLiteralContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ArcParserRULE_mapLiteral
	return p
}

func InitEmptyMapLiteralContext(p *MapLiteralContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ArcParserRULE_mapLiteral
}

func (*MapLiteralContext) IsMapLiteralContext() {}

func NewMapLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MapLiteralContext {
	var p = new(MapLiteralContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ArcParserRULE_mapLiteral

	return p
}

func (s *MapLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *MapLiteralContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(ArcParserLBRACE, 0)
}

func (s *MapLiteralContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(ArcParserRBRACE, 0)
}

func (s *MapLiteralContext) AllMapEntry() []IMapEntryContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IMapEntryContext); ok {
			len++
		}
	}

	tst := make([]IMapEntryContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IMapEntryContext); ok {
			tst[i] = t.(IMapEntryContext)
			i++
		}
	}

	return tst
}

func (s *MapLiteralContext) MapEntry(i int) IMapEntryContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMapEntryContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMapEntryContext)
}

func (s *MapLiteralContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(ArcParserCOMMA)
}

func (s *MapLiteralContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(ArcParserCOMMA, i)
}

func (s *MapLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MapLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MapLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ArcParserListener); ok {
		listenerT.EnterMapLiteral(s)
	}
}

func (s *MapLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ArcParserListener); ok {
		listenerT.ExitMapLiteral(s)
	}
}

func (s *MapLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ArcParserVisitor:
		return t.VisitMapLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ArcParser) MapLiteral() (localctx IMapLiteralContext) {
	localctx = NewMapLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 92, ArcParserRULE_mapLiteral)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(475)
		p.Match(ArcParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(484)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64((_la-35)) & ^0x3f) == 0 && ((int64(1)<<(_la-35))&135296811035) != 0 {
		{
			p.SetState(476)
			p.MapEntry()
		}
		p.SetState(481)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == ArcParserCOMMA {
			{
				p.SetState(477)
				p.Match(ArcParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(478)
				p.MapEntry()
			}

			p.SetState(483)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(486)
		p.Match(ArcParserRBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMapEntryContext is an interface to support dynamic dispatch.
type IMapEntryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	COLON() antlr.TerminalNode

	// IsMapEntryContext differentiates from other interfaces.
	IsMapEntryContext()
}

type MapEntryContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMapEntryContext() *MapEntryContext {
	var p = new(MapEntryContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ArcParserRULE_mapEntry
	return p
}

func InitEmptyMapEntryContext(p *MapEntryContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ArcParserRULE_mapEntry
}

func (*MapEntryContext) IsMapEntryContext() {}

func NewMapEntryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MapEntryContext {
	var p = new(MapEntryContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ArcParserRULE_mapEntry

	return p
}

func (s *MapEntryContext) GetParser() antlr.Parser { return s.parser }

func (s *MapEntryContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *MapEntryContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *MapEntryContext) COLON() antlr.TerminalNode {
	return s.GetToken(ArcParserCOLON, 0)
}

func (s *MapEntryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MapEntryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MapEntryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ArcParserListener); ok {
		listenerT.EnterMapEntry(s)
	}
}

func (s *MapEntryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ArcParserListener); ok {
		listenerT.ExitMapEntry(s)
	}
}

func (s *MapEntryContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ArcParserVisitor:
		return t.VisitMapEntry(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ArcParser) MapEntry() (localctx IMapEntryContext) {
	localctx = NewMapEntryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 94, ArcParserRULE_mapEntry)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(488)
		p.Expression()
	}
	{
		p.SetState(489)
		p.Match(ArcParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(490)
		p.Expression()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStructLiteralContext is an interface to support dynamic dispatch.
type IStructLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	AllFieldInit() []IFieldInitContext
	FieldInit(i int) IFieldInitContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsStructLiteralContext differentiates from other interfaces.
	IsStructLiteralContext()
}

type StructLiteralContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStructLiteralContext() *StructLiteralContext {
	var p = new(StructLiteralContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ArcParserRULE_structLiteral
	return p
}

func InitEmptyStructLiteralContext(p *StructLiteralContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ArcParserRULE_structLiteral
}

func (*StructLiteralContext) IsStructLiteralContext() {}

func NewStructLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructLiteralContext {
	var p = new(StructLiteralContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ArcParserRULE_structLiteral

	return p
}

func (s *StructLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *StructLiteralContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ArcParserIDENTIFIER, 0)
}

func (s *StructLiteralContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(ArcParserLBRACE, 0)
}

func (s *StructLiteralContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(ArcParserRBRACE, 0)
}

func (s *StructLiteralContext) AllFieldInit() []IFieldInitContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFieldInitContext); ok {
			len++
		}
	}

	tst := make([]IFieldInitContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFieldInitContext); ok {
			tst[i] = t.(IFieldInitContext)
			i++
		}
	}

	return tst
}

func (s *StructLiteralContext) FieldInit(i int) IFieldInitContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldInitContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldInitContext)
}

func (s *StructLiteralContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(ArcParserCOMMA)
}

func (s *StructLiteralContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(ArcParserCOMMA, i)
}

func (s *StructLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StructLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ArcParserListener); ok {
		listenerT.EnterStructLiteral(s)
	}
}

func (s *StructLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ArcParserListener); ok {
		listenerT.ExitStructLiteral(s)
	}
}

func (s *StructLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ArcParserVisitor:
		return t.VisitStructLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ArcParser) StructLiteral() (localctx IStructLiteralContext) {
	localctx = NewStructLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 96, ArcParserRULE_structLiteral)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(492)
		p.Match(ArcParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(493)
		p.Match(ArcParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(502)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ArcParserIDENTIFIER {
		{
			p.SetState(494)
			p.FieldInit()
		}
		p.SetState(499)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == ArcParserCOMMA {
			{
				p.SetState(495)
				p.Match(ArcParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(496)
				p.FieldInit()
			}

			p.SetState(501)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(504)
		p.Match(ArcParserRBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFieldInitContext is an interface to support dynamic dispatch.
type IFieldInitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	COLON() antlr.TerminalNode
	Expression() IExpressionContext

	// IsFieldInitContext differentiates from other interfaces.
	IsFieldInitContext()
}

type FieldInitContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldInitContext() *FieldInitContext {
	var p = new(FieldInitContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ArcParserRULE_fieldInit
	return p
}

func InitEmptyFieldInitContext(p *FieldInitContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ArcParserRULE_fieldInit
}

func (*FieldInitContext) IsFieldInitContext() {}

func NewFieldInitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldInitContext {
	var p = new(FieldInitContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ArcParserRULE_fieldInit

	return p
}

func (s *FieldInitContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldInitContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ArcParserIDENTIFIER, 0)
}

func (s *FieldInitContext) COLON() antlr.TerminalNode {
	return s.GetToken(ArcParserCOLON, 0)
}

func (s *FieldInitContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *FieldInitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldInitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldInitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ArcParserListener); ok {
		listenerT.EnterFieldInit(s)
	}
}

func (s *FieldInitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ArcParserListener); ok {
		listenerT.ExitFieldInit(s)
	}
}

func (s *FieldInitContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ArcParserVisitor:
		return t.VisitFieldInit(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ArcParser) FieldInit() (localctx IFieldInitContext) {
	localctx = NewFieldInitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 98, ArcParserRULE_fieldInit)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(506)
		p.Match(ArcParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(507)
		p.Match(ArcParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(508)
		p.Expression()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IArgumentListContext is an interface to support dynamic dispatch.
type IArgumentListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsArgumentListContext differentiates from other interfaces.
	IsArgumentListContext()
}

type ArgumentListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgumentListContext() *ArgumentListContext {
	var p = new(ArgumentListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ArcParserRULE_argumentList
	return p
}

func InitEmptyArgumentListContext(p *ArgumentListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ArcParserRULE_argumentList
}

func (*ArgumentListContext) IsArgumentListContext() {}

func NewArgumentListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgumentListContext {
	var p = new(ArgumentListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ArcParserRULE_argumentList

	return p
}

func (s *ArgumentListContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgumentListContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ArgumentListContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ArgumentListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(ArcParserCOMMA)
}

func (s *ArgumentListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(ArcParserCOMMA, i)
}

func (s *ArgumentListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgumentListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgumentListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ArcParserListener); ok {
		listenerT.EnterArgumentList(s)
	}
}

func (s *ArgumentListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ArcParserListener); ok {
		listenerT.ExitArgumentList(s)
	}
}

func (s *ArgumentListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ArcParserVisitor:
		return t.VisitArgumentList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ArcParser) ArgumentList() (localctx IArgumentListContext) {
	localctx = NewArgumentListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 100, ArcParserRULE_argumentList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(510)
		p.Expression()
	}
	p.SetState(515)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ArcParserCOMMA {
		{
			p.SetState(511)
			p.Match(ArcParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(512)
			p.Expression()
		}

		p.SetState(517)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICastExpressionContext is an interface to support dynamic dispatch.
type ICastExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CAST() antlr.TerminalNode
	LT() antlr.TerminalNode
	Type_() ITypeContext
	GT() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	Expression() IExpressionContext
	RPAREN() antlr.TerminalNode

	// IsCastExpressionContext differentiates from other interfaces.
	IsCastExpressionContext()
}

type CastExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCastExpressionContext() *CastExpressionContext {
	var p = new(CastExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ArcParserRULE_castExpression
	return p
}

func InitEmptyCastExpressionContext(p *CastExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ArcParserRULE_castExpression
}

func (*CastExpressionContext) IsCastExpressionContext() {}

func NewCastExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CastExpressionContext {
	var p = new(CastExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ArcParserRULE_castExpression

	return p
}

func (s *CastExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *CastExpressionContext) CAST() antlr.TerminalNode {
	return s.GetToken(ArcParserCAST, 0)
}

func (s *CastExpressionContext) LT() antlr.TerminalNode {
	return s.GetToken(ArcParserLT, 0)
}

func (s *CastExpressionContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *CastExpressionContext) GT() antlr.TerminalNode {
	return s.GetToken(ArcParserGT, 0)
}

func (s *CastExpressionContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(ArcParserLPAREN, 0)
}

func (s *CastExpressionContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *CastExpressionContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(ArcParserRPAREN, 0)
}

func (s *CastExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CastExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CastExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ArcParserListener); ok {
		listenerT.EnterCastExpression(s)
	}
}

func (s *CastExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ArcParserListener); ok {
		listenerT.ExitCastExpression(s)
	}
}

func (s *CastExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ArcParserVisitor:
		return t.VisitCastExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ArcParser) CastExpression() (localctx ICastExpressionContext) {
	localctx = NewCastExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 102, ArcParserRULE_castExpression)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(518)
		p.Match(ArcParserCAST)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(519)
		p.Match(ArcParserLT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(520)
		p.Type_()
	}
	{
		p.SetState(521)
		p.Match(ArcParserGT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(522)
		p.Match(ArcParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(523)
		p.Expression()
	}
	{
		p.SetState(524)
		p.Match(ArcParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAllocaExpressionContext is an interface to support dynamic dispatch.
type IAllocaExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ALLOCA() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	Type_() ITypeContext
	RPAREN() antlr.TerminalNode
	COMMA() antlr.TerminalNode
	Expression() IExpressionContext

	// IsAllocaExpressionContext differentiates from other interfaces.
	IsAllocaExpressionContext()
}

type AllocaExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAllocaExpressionContext() *AllocaExpressionContext {
	var p = new(AllocaExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ArcParserRULE_allocaExpression
	return p
}

func InitEmptyAllocaExpressionContext(p *AllocaExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ArcParserRULE_allocaExpression
}

func (*AllocaExpressionContext) IsAllocaExpressionContext() {}

func NewAllocaExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AllocaExpressionContext {
	var p = new(AllocaExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ArcParserRULE_allocaExpression

	return p
}

func (s *AllocaExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *AllocaExpressionContext) ALLOCA() antlr.TerminalNode {
	return s.GetToken(ArcParserALLOCA, 0)
}

func (s *AllocaExpressionContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(ArcParserLPAREN, 0)
}

func (s *AllocaExpressionContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *AllocaExpressionContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(ArcParserRPAREN, 0)
}

func (s *AllocaExpressionContext) COMMA() antlr.TerminalNode {
	return s.GetToken(ArcParserCOMMA, 0)
}

func (s *AllocaExpressionContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AllocaExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AllocaExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AllocaExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ArcParserListener); ok {
		listenerT.EnterAllocaExpression(s)
	}
}

func (s *AllocaExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ArcParserListener); ok {
		listenerT.ExitAllocaExpression(s)
	}
}

func (s *AllocaExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ArcParserVisitor:
		return t.VisitAllocaExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ArcParser) AllocaExpression() (localctx IAllocaExpressionContext) {
	localctx = NewAllocaExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 104, ArcParserRULE_allocaExpression)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(526)
		p.Match(ArcParserALLOCA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(527)
		p.Match(ArcParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(528)
		p.Type_()
	}
	p.SetState(531)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ArcParserCOMMA {
		{
			p.SetState(529)
			p.Match(ArcParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(530)
			p.Expression()
		}

	}
	{
		p.SetState(533)
		p.Match(ArcParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}
