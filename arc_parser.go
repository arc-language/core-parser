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
		"'return'", "'if'", "'else'", "'defer'", "'extern'", "'self'", "'int8'",
		"'int16'", "'int32'", "'int64'", "'uint8'", "'uint16'", "'uint32'",
		"'uint64'", "'usize'", "'isize'", "'float32'", "'float64'", "'byte'",
		"'bool'", "'char'", "'string'", "'void'", "'vector'", "'map'", "'alloca'",
		"'cast'", "'+'", "'-'", "'*'", "'/'", "'%'", "'=='", "'!='", "'<'",
		"'<='", "'>'", "'>='", "'&&'", "'||'", "'!'", "'&'", "'@'", "'='", "'->'",
		"'('", "')'", "'{'", "'}'", "'['", "']'", "','", "':'", "';'", "'.'",
		"'...'",
	}
	staticData.SymbolicNames = []string{
		"", "IMPORT", "NAMESPACE", "LET", "CONST", "FUNC", "STRUCT", "RETURN",
		"IF", "ELSE", "DEFER", "EXTERN", "SELF", "INT8", "INT16", "INT32", "INT64",
		"UINT8", "UINT16", "UINT32", "UINT64", "USIZE", "ISIZE", "FLOAT32",
		"FLOAT64", "BYTE", "BOOL", "CHAR", "STRING", "VOID", "VECTOR", "MAP",
		"ALLOCA", "CAST", "PLUS", "MINUS", "STAR", "SLASH", "PERCENT", "EQ",
		"NE", "LT", "LE", "GT", "GE", "AND", "OR", "NOT", "AMP", "AT", "ASSIGN",
		"ARROW", "LPAREN", "RPAREN", "LBRACE", "RBRACE", "LBRACKET", "RBRACKET",
		"COMMA", "COLON", "SEMICOLON", "DOT", "ELLIPSIS", "BOOLEAN_LITERAL",
		"INTEGER_LITERAL", "FLOAT_LITERAL", "STRING_LITERAL", "CHAR_LITERAL",
		"IDENTIFIER", "WS", "LINE_COMMENT", "BLOCK_COMMENT",
	}
	staticData.RuleNames = []string{
		"compilationUnit", "importDecl", "importSpec", "namespaceDecl", "topLevelDecl",
		"externDecl", "externMember", "externFunctionDecl", "externParameterList",
		"functionDecl", "parameterList", "parameter", "structDecl", "structField",
		"variableDecl", "constDecl", "type", "primitiveType", "pointerType",
		"referenceType", "vectorType", "mapType", "block", "statement", "assignmentStmt",
		"leftHandSide", "expressionStmt", "returnStmt", "ifStmt", "deferStmt",
		"expression", "logicalOrExpression", "logicalAndExpression", "equalityExpression",
		"relationalExpression", "additiveExpression", "multiplicativeExpression",
		"unaryExpression", "postfixExpression", "postfixOp", "primaryExpression",
		"literal", "vectorLiteral", "mapLiteral", "mapEntry", "structLiteral",
		"fieldInit", "argumentList", "castExpression", "allocaExpression",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 71, 497, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 7, 36,
		2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7, 41, 2,
		42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46, 2, 47,
		7, 47, 2, 48, 7, 48, 2, 49, 7, 49, 1, 0, 1, 0, 1, 0, 5, 0, 104, 8, 0, 10,
		0, 12, 0, 107, 9, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 5, 1, 116,
		8, 1, 10, 1, 12, 1, 119, 9, 1, 1, 1, 3, 1, 122, 8, 1, 1, 2, 1, 2, 1, 3,
		1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 134, 8, 4, 1, 5, 1, 5,
		1, 5, 1, 5, 5, 5, 140, 8, 5, 10, 5, 12, 5, 143, 9, 5, 1, 5, 1, 5, 1, 6,
		1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 3, 7, 153, 8, 7, 1, 7, 1, 7, 3, 7, 157, 8,
		7, 1, 8, 1, 8, 1, 8, 5, 8, 162, 8, 8, 10, 8, 12, 8, 165, 9, 8, 1, 8, 1,
		8, 3, 8, 169, 8, 8, 1, 8, 3, 8, 172, 8, 8, 1, 9, 1, 9, 1, 9, 1, 9, 3, 9,
		178, 8, 9, 1, 9, 1, 9, 3, 9, 182, 8, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10,
		5, 10, 189, 8, 10, 10, 10, 12, 10, 192, 9, 10, 1, 10, 1, 10, 3, 10, 196,
		8, 10, 1, 10, 3, 10, 199, 8, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 12, 1,
		12, 1, 12, 1, 12, 5, 12, 209, 8, 12, 10, 12, 12, 12, 212, 9, 12, 1, 12,
		1, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 14, 3, 14, 224,
		8, 14, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 15, 3, 15, 233, 8,
		15, 1, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 3, 16,
		244, 8, 16, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 19, 1, 19, 1, 19, 1,
		20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21,
		1, 21, 1, 22, 1, 22, 5, 22, 268, 8, 22, 10, 22, 12, 22, 271, 9, 22, 1,
		22, 1, 22, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 3, 23,
		283, 8, 23, 1, 24, 1, 24, 1, 24, 1, 24, 1, 25, 1, 25, 1, 25, 1, 25, 1,
		25, 1, 25, 1, 25, 3, 25, 296, 8, 25, 1, 26, 1, 26, 1, 27, 1, 27, 3, 27,
		302, 8, 27, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 5,
		28, 312, 8, 28, 10, 28, 12, 28, 315, 9, 28, 1, 28, 1, 28, 3, 28, 319, 8,
		28, 1, 29, 1, 29, 1, 29, 3, 29, 324, 8, 29, 1, 30, 1, 30, 1, 31, 1, 31,
		1, 31, 5, 31, 331, 8, 31, 10, 31, 12, 31, 334, 9, 31, 1, 32, 1, 32, 1,
		32, 5, 32, 339, 8, 32, 10, 32, 12, 32, 342, 9, 32, 1, 33, 1, 33, 1, 33,
		5, 33, 347, 8, 33, 10, 33, 12, 33, 350, 9, 33, 1, 34, 1, 34, 1, 34, 5,
		34, 355, 8, 34, 10, 34, 12, 34, 358, 9, 34, 1, 35, 1, 35, 1, 35, 5, 35,
		363, 8, 35, 10, 35, 12, 35, 366, 9, 35, 1, 36, 1, 36, 1, 36, 5, 36, 371,
		8, 36, 10, 36, 12, 36, 374, 9, 36, 1, 37, 1, 37, 1, 37, 3, 37, 379, 8,
		37, 1, 38, 1, 38, 5, 38, 383, 8, 38, 10, 38, 12, 38, 386, 9, 38, 1, 39,
		1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 3, 39, 394, 8, 39, 1, 39, 1, 39, 1,
		39, 3, 39, 399, 8, 39, 1, 39, 3, 39, 402, 8, 39, 1, 40, 1, 40, 1, 40, 1,
		40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 3, 40, 413, 8, 40, 1, 41, 1, 41,
		1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 3, 41, 422, 8, 41, 1, 42, 1, 42, 1,
		42, 1, 42, 5, 42, 428, 8, 42, 10, 42, 12, 42, 431, 9, 42, 3, 42, 433, 8,
		42, 1, 42, 1, 42, 1, 43, 1, 43, 1, 43, 1, 43, 5, 43, 441, 8, 43, 10, 43,
		12, 43, 444, 9, 43, 3, 43, 446, 8, 43, 1, 43, 1, 43, 1, 44, 1, 44, 1, 44,
		1, 44, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 5, 45, 459, 8, 45, 10, 45, 12,
		45, 462, 9, 45, 3, 45, 464, 8, 45, 1, 45, 1, 45, 1, 46, 1, 46, 1, 46, 1,
		46, 1, 47, 1, 47, 1, 47, 5, 47, 475, 8, 47, 10, 47, 12, 47, 478, 9, 47,
		1, 48, 1, 48, 1, 48, 1, 48, 1, 48, 1, 48, 1, 48, 1, 48, 1, 49, 1, 49, 1,
		49, 1, 49, 1, 49, 3, 49, 493, 8, 49, 1, 49, 1, 49, 1, 49, 0, 0, 50, 0,
		2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38,
		40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 72, 74,
		76, 78, 80, 82, 84, 86, 88, 90, 92, 94, 96, 98, 0, 6, 1, 0, 13, 29, 1,
		0, 39, 40, 1, 0, 41, 44, 1, 0, 34, 35, 1, 0, 36, 38, 2, 0, 35, 36, 47,
		48, 519, 0, 105, 1, 0, 0, 0, 2, 121, 1, 0, 0, 0, 4, 123, 1, 0, 0, 0, 6,
		125, 1, 0, 0, 0, 8, 133, 1, 0, 0, 0, 10, 135, 1, 0, 0, 0, 12, 146, 1, 0,
		0, 0, 14, 148, 1, 0, 0, 0, 16, 171, 1, 0, 0, 0, 18, 173, 1, 0, 0, 0, 20,
		198, 1, 0, 0, 0, 22, 200, 1, 0, 0, 0, 24, 204, 1, 0, 0, 0, 26, 215, 1,
		0, 0, 0, 28, 219, 1, 0, 0, 0, 30, 228, 1, 0, 0, 0, 32, 243, 1, 0, 0, 0,
		34, 245, 1, 0, 0, 0, 36, 247, 1, 0, 0, 0, 38, 250, 1, 0, 0, 0, 40, 253,
		1, 0, 0, 0, 42, 258, 1, 0, 0, 0, 44, 265, 1, 0, 0, 0, 46, 282, 1, 0, 0,
		0, 48, 284, 1, 0, 0, 0, 50, 295, 1, 0, 0, 0, 52, 297, 1, 0, 0, 0, 54, 299,
		1, 0, 0, 0, 56, 303, 1, 0, 0, 0, 58, 320, 1, 0, 0, 0, 60, 325, 1, 0, 0,
		0, 62, 327, 1, 0, 0, 0, 64, 335, 1, 0, 0, 0, 66, 343, 1, 0, 0, 0, 68, 351,
		1, 0, 0, 0, 70, 359, 1, 0, 0, 0, 72, 367, 1, 0, 0, 0, 74, 378, 1, 0, 0,
		0, 76, 380, 1, 0, 0, 0, 78, 401, 1, 0, 0, 0, 80, 412, 1, 0, 0, 0, 82, 421,
		1, 0, 0, 0, 84, 423, 1, 0, 0, 0, 86, 436, 1, 0, 0, 0, 88, 449, 1, 0, 0,
		0, 90, 453, 1, 0, 0, 0, 92, 467, 1, 0, 0, 0, 94, 471, 1, 0, 0, 0, 96, 479,
		1, 0, 0, 0, 98, 487, 1, 0, 0, 0, 100, 104, 3, 2, 1, 0, 101, 104, 3, 6,
		3, 0, 102, 104, 3, 8, 4, 0, 103, 100, 1, 0, 0, 0, 103, 101, 1, 0, 0, 0,
		103, 102, 1, 0, 0, 0, 104, 107, 1, 0, 0, 0, 105, 103, 1, 0, 0, 0, 105,
		106, 1, 0, 0, 0, 106, 108, 1, 0, 0, 0, 107, 105, 1, 0, 0, 0, 108, 109,
		5, 0, 0, 1, 109, 1, 1, 0, 0, 0, 110, 111, 5, 1, 0, 0, 111, 122, 5, 66,
		0, 0, 112, 113, 5, 1, 0, 0, 113, 117, 5, 52, 0, 0, 114, 116, 3, 4, 2, 0,
		115, 114, 1, 0, 0, 0, 116, 119, 1, 0, 0, 0, 117, 115, 1, 0, 0, 0, 117,
		118, 1, 0, 0, 0, 118, 120, 1, 0, 0, 0, 119, 117, 1, 0, 0, 0, 120, 122,
		5, 53, 0, 0, 121, 110, 1, 0, 0, 0, 121, 112, 1, 0, 0, 0, 122, 3, 1, 0,
		0, 0, 123, 124, 5, 66, 0, 0, 124, 5, 1, 0, 0, 0, 125, 126, 5, 2, 0, 0,
		126, 127, 5, 68, 0, 0, 127, 7, 1, 0, 0, 0, 128, 134, 3, 18, 9, 0, 129,
		134, 3, 24, 12, 0, 130, 134, 3, 28, 14, 0, 131, 134, 3, 30, 15, 0, 132,
		134, 3, 10, 5, 0, 133, 128, 1, 0, 0, 0, 133, 129, 1, 0, 0, 0, 133, 130,
		1, 0, 0, 0, 133, 131, 1, 0, 0, 0, 133, 132, 1, 0, 0, 0, 134, 9, 1, 0, 0,
		0, 135, 136, 5, 11, 0, 0, 136, 137, 5, 68, 0, 0, 137, 141, 5, 54, 0, 0,
		138, 140, 3, 12, 6, 0, 139, 138, 1, 0, 0, 0, 140, 143, 1, 0, 0, 0, 141,
		139, 1, 0, 0, 0, 141, 142, 1, 0, 0, 0, 142, 144, 1, 0, 0, 0, 143, 141,
		1, 0, 0, 0, 144, 145, 5, 55, 0, 0, 145, 11, 1, 0, 0, 0, 146, 147, 3, 14,
		7, 0, 147, 13, 1, 0, 0, 0, 148, 149, 5, 5, 0, 0, 149, 150, 5, 68, 0, 0,
		150, 152, 5, 52, 0, 0, 151, 153, 3, 16, 8, 0, 152, 151, 1, 0, 0, 0, 152,
		153, 1, 0, 0, 0, 153, 154, 1, 0, 0, 0, 154, 156, 5, 53, 0, 0, 155, 157,
		3, 32, 16, 0, 156, 155, 1, 0, 0, 0, 156, 157, 1, 0, 0, 0, 157, 15, 1, 0,
		0, 0, 158, 163, 3, 32, 16, 0, 159, 160, 5, 58, 0, 0, 160, 162, 3, 32, 16,
		0, 161, 159, 1, 0, 0, 0, 162, 165, 1, 0, 0, 0, 163, 161, 1, 0, 0, 0, 163,
		164, 1, 0, 0, 0, 164, 168, 1, 0, 0, 0, 165, 163, 1, 0, 0, 0, 166, 167,
		5, 58, 0, 0, 167, 169, 5, 62, 0, 0, 168, 166, 1, 0, 0, 0, 168, 169, 1,
		0, 0, 0, 169, 172, 1, 0, 0, 0, 170, 172, 5, 62, 0, 0, 171, 158, 1, 0, 0,
		0, 171, 170, 1, 0, 0, 0, 172, 17, 1, 0, 0, 0, 173, 174, 5, 5, 0, 0, 174,
		175, 5, 68, 0, 0, 175, 177, 5, 52, 0, 0, 176, 178, 3, 20, 10, 0, 177, 176,
		1, 0, 0, 0, 177, 178, 1, 0, 0, 0, 178, 179, 1, 0, 0, 0, 179, 181, 5, 53,
		0, 0, 180, 182, 3, 32, 16, 0, 181, 180, 1, 0, 0, 0, 181, 182, 1, 0, 0,
		0, 182, 183, 1, 0, 0, 0, 183, 184, 3, 44, 22, 0, 184, 19, 1, 0, 0, 0, 185,
		190, 3, 22, 11, 0, 186, 187, 5, 58, 0, 0, 187, 189, 3, 22, 11, 0, 188,
		186, 1, 0, 0, 0, 189, 192, 1, 0, 0, 0, 190, 188, 1, 0, 0, 0, 190, 191,
		1, 0, 0, 0, 191, 195, 1, 0, 0, 0, 192, 190, 1, 0, 0, 0, 193, 194, 5, 58,
		0, 0, 194, 196, 5, 62, 0, 0, 195, 193, 1, 0, 0, 0, 195, 196, 1, 0, 0, 0,
		196, 199, 1, 0, 0, 0, 197, 199, 5, 62, 0, 0, 198, 185, 1, 0, 0, 0, 198,
		197, 1, 0, 0, 0, 199, 21, 1, 0, 0, 0, 200, 201, 5, 68, 0, 0, 201, 202,
		5, 59, 0, 0, 202, 203, 3, 32, 16, 0, 203, 23, 1, 0, 0, 0, 204, 205, 5,
		6, 0, 0, 205, 206, 5, 68, 0, 0, 206, 210, 5, 54, 0, 0, 207, 209, 3, 26,
		13, 0, 208, 207, 1, 0, 0, 0, 209, 212, 1, 0, 0, 0, 210, 208, 1, 0, 0, 0,
		210, 211, 1, 0, 0, 0, 211, 213, 1, 0, 0, 0, 212, 210, 1, 0, 0, 0, 213,
		214, 5, 55, 0, 0, 214, 25, 1, 0, 0, 0, 215, 216, 5, 68, 0, 0, 216, 217,
		5, 59, 0, 0, 217, 218, 3, 32, 16, 0, 218, 27, 1, 0, 0, 0, 219, 220, 5,
		3, 0, 0, 220, 223, 5, 68, 0, 0, 221, 222, 5, 59, 0, 0, 222, 224, 3, 32,
		16, 0, 223, 221, 1, 0, 0, 0, 223, 224, 1, 0, 0, 0, 224, 225, 1, 0, 0, 0,
		225, 226, 5, 50, 0, 0, 226, 227, 3, 60, 30, 0, 227, 29, 1, 0, 0, 0, 228,
		229, 5, 4, 0, 0, 229, 232, 5, 68, 0, 0, 230, 231, 5, 59, 0, 0, 231, 233,
		3, 32, 16, 0, 232, 230, 1, 0, 0, 0, 232, 233, 1, 0, 0, 0, 233, 234, 1,
		0, 0, 0, 234, 235, 5, 50, 0, 0, 235, 236, 3, 60, 30, 0, 236, 31, 1, 0,
		0, 0, 237, 244, 3, 34, 17, 0, 238, 244, 3, 36, 18, 0, 239, 244, 3, 38,
		19, 0, 240, 244, 3, 40, 20, 0, 241, 244, 3, 42, 21, 0, 242, 244, 5, 68,
		0, 0, 243, 237, 1, 0, 0, 0, 243, 238, 1, 0, 0, 0, 243, 239, 1, 0, 0, 0,
		243, 240, 1, 0, 0, 0, 243, 241, 1, 0, 0, 0, 243, 242, 1, 0, 0, 0, 244,
		33, 1, 0, 0, 0, 245, 246, 7, 0, 0, 0, 246, 35, 1, 0, 0, 0, 247, 248, 5,
		36, 0, 0, 248, 249, 3, 32, 16, 0, 249, 37, 1, 0, 0, 0, 250, 251, 5, 48,
		0, 0, 251, 252, 3, 32, 16, 0, 252, 39, 1, 0, 0, 0, 253, 254, 5, 30, 0,
		0, 254, 255, 5, 41, 0, 0, 255, 256, 3, 32, 16, 0, 256, 257, 5, 43, 0, 0,
		257, 41, 1, 0, 0, 0, 258, 259, 5, 31, 0, 0, 259, 260, 5, 41, 0, 0, 260,
		261, 3, 32, 16, 0, 261, 262, 5, 58, 0, 0, 262, 263, 3, 32, 16, 0, 263,
		264, 5, 43, 0, 0, 264, 43, 1, 0, 0, 0, 265, 269, 5, 54, 0, 0, 266, 268,
		3, 46, 23, 0, 267, 266, 1, 0, 0, 0, 268, 271, 1, 0, 0, 0, 269, 267, 1,
		0, 0, 0, 269, 270, 1, 0, 0, 0, 270, 272, 1, 0, 0, 0, 271, 269, 1, 0, 0,
		0, 272, 273, 5, 55, 0, 0, 273, 45, 1, 0, 0, 0, 274, 283, 3, 28, 14, 0,
		275, 283, 3, 30, 15, 0, 276, 283, 3, 48, 24, 0, 277, 283, 3, 52, 26, 0,
		278, 283, 3, 54, 27, 0, 279, 283, 3, 56, 28, 0, 280, 283, 3, 58, 29, 0,
		281, 283, 3, 44, 22, 0, 282, 274, 1, 0, 0, 0, 282, 275, 1, 0, 0, 0, 282,
		276, 1, 0, 0, 0, 282, 277, 1, 0, 0, 0, 282, 278, 1, 0, 0, 0, 282, 279,
		1, 0, 0, 0, 282, 280, 1, 0, 0, 0, 282, 281, 1, 0, 0, 0, 283, 47, 1, 0,
		0, 0, 284, 285, 3, 50, 25, 0, 285, 286, 5, 50, 0, 0, 286, 287, 3, 60, 30,
		0, 287, 49, 1, 0, 0, 0, 288, 296, 5, 68, 0, 0, 289, 290, 5, 36, 0, 0, 290,
		296, 3, 60, 30, 0, 291, 292, 3, 60, 30, 0, 292, 293, 5, 61, 0, 0, 293,
		294, 5, 68, 0, 0, 294, 296, 1, 0, 0, 0, 295, 288, 1, 0, 0, 0, 295, 289,
		1, 0, 0, 0, 295, 291, 1, 0, 0, 0, 296, 51, 1, 0, 0, 0, 297, 298, 3, 60,
		30, 0, 298, 53, 1, 0, 0, 0, 299, 301, 5, 7, 0, 0, 300, 302, 3, 60, 30,
		0, 301, 300, 1, 0, 0, 0, 301, 302, 1, 0, 0, 0, 302, 55, 1, 0, 0, 0, 303,
		304, 5, 8, 0, 0, 304, 305, 3, 60, 30, 0, 305, 313, 3, 44, 22, 0, 306, 307,
		5, 9, 0, 0, 307, 308, 5, 8, 0, 0, 308, 309, 3, 60, 30, 0, 309, 310, 3,
		44, 22, 0, 310, 312, 1, 0, 0, 0, 311, 306, 1, 0, 0, 0, 312, 315, 1, 0,
		0, 0, 313, 311, 1, 0, 0, 0, 313, 314, 1, 0, 0, 0, 314, 318, 1, 0, 0, 0,
		315, 313, 1, 0, 0, 0, 316, 317, 5, 9, 0, 0, 317, 319, 3, 44, 22, 0, 318,
		316, 1, 0, 0, 0, 318, 319, 1, 0, 0, 0, 319, 57, 1, 0, 0, 0, 320, 323, 5,
		10, 0, 0, 321, 324, 3, 48, 24, 0, 322, 324, 3, 60, 30, 0, 323, 321, 1,
		0, 0, 0, 323, 322, 1, 0, 0, 0, 324, 59, 1, 0, 0, 0, 325, 326, 3, 62, 31,
		0, 326, 61, 1, 0, 0, 0, 327, 332, 3, 64, 32, 0, 328, 329, 5, 46, 0, 0,
		329, 331, 3, 64, 32, 0, 330, 328, 1, 0, 0, 0, 331, 334, 1, 0, 0, 0, 332,
		330, 1, 0, 0, 0, 332, 333, 1, 0, 0, 0, 333, 63, 1, 0, 0, 0, 334, 332, 1,
		0, 0, 0, 335, 340, 3, 66, 33, 0, 336, 337, 5, 45, 0, 0, 337, 339, 3, 66,
		33, 0, 338, 336, 1, 0, 0, 0, 339, 342, 1, 0, 0, 0, 340, 338, 1, 0, 0, 0,
		340, 341, 1, 0, 0, 0, 341, 65, 1, 0, 0, 0, 342, 340, 1, 0, 0, 0, 343, 348,
		3, 68, 34, 0, 344, 345, 7, 1, 0, 0, 345, 347, 3, 68, 34, 0, 346, 344, 1,
		0, 0, 0, 347, 350, 1, 0, 0, 0, 348, 346, 1, 0, 0, 0, 348, 349, 1, 0, 0,
		0, 349, 67, 1, 0, 0, 0, 350, 348, 1, 0, 0, 0, 351, 356, 3, 70, 35, 0, 352,
		353, 7, 2, 0, 0, 353, 355, 3, 70, 35, 0, 354, 352, 1, 0, 0, 0, 355, 358,
		1, 0, 0, 0, 356, 354, 1, 0, 0, 0, 356, 357, 1, 0, 0, 0, 357, 69, 1, 0,
		0, 0, 358, 356, 1, 0, 0, 0, 359, 364, 3, 72, 36, 0, 360, 361, 7, 3, 0,
		0, 361, 363, 3, 72, 36, 0, 362, 360, 1, 0, 0, 0, 363, 366, 1, 0, 0, 0,
		364, 362, 1, 0, 0, 0, 364, 365, 1, 0, 0, 0, 365, 71, 1, 0, 0, 0, 366, 364,
		1, 0, 0, 0, 367, 372, 3, 74, 37, 0, 368, 369, 7, 4, 0, 0, 369, 371, 3,
		74, 37, 0, 370, 368, 1, 0, 0, 0, 371, 374, 1, 0, 0, 0, 372, 370, 1, 0,
		0, 0, 372, 373, 1, 0, 0, 0, 373, 73, 1, 0, 0, 0, 374, 372, 1, 0, 0, 0,
		375, 376, 7, 5, 0, 0, 376, 379, 3, 74, 37, 0, 377, 379, 3, 76, 38, 0, 378,
		375, 1, 0, 0, 0, 378, 377, 1, 0, 0, 0, 379, 75, 1, 0, 0, 0, 380, 384, 3,
		80, 40, 0, 381, 383, 3, 78, 39, 0, 382, 381, 1, 0, 0, 0, 383, 386, 1, 0,
		0, 0, 384, 382, 1, 0, 0, 0, 384, 385, 1, 0, 0, 0, 385, 77, 1, 0, 0, 0,
		386, 384, 1, 0, 0, 0, 387, 388, 5, 61, 0, 0, 388, 402, 5, 68, 0, 0, 389,
		390, 5, 61, 0, 0, 390, 391, 5, 68, 0, 0, 391, 393, 5, 52, 0, 0, 392, 394,
		3, 94, 47, 0, 393, 392, 1, 0, 0, 0, 393, 394, 1, 0, 0, 0, 394, 395, 1,
		0, 0, 0, 395, 402, 5, 53, 0, 0, 396, 398, 5, 52, 0, 0, 397, 399, 3, 94,
		47, 0, 398, 397, 1, 0, 0, 0, 398, 399, 1, 0, 0, 0, 399, 400, 1, 0, 0, 0,
		400, 402, 5, 53, 0, 0, 401, 387, 1, 0, 0, 0, 401, 389, 1, 0, 0, 0, 401,
		396, 1, 0, 0, 0, 402, 79, 1, 0, 0, 0, 403, 413, 3, 82, 41, 0, 404, 413,
		5, 68, 0, 0, 405, 406, 5, 52, 0, 0, 406, 407, 3, 60, 30, 0, 407, 408, 5,
		53, 0, 0, 408, 413, 1, 0, 0, 0, 409, 413, 3, 96, 48, 0, 410, 413, 3, 98,
		49, 0, 411, 413, 3, 90, 45, 0, 412, 403, 1, 0, 0, 0, 412, 404, 1, 0, 0,
		0, 412, 405, 1, 0, 0, 0, 412, 409, 1, 0, 0, 0, 412, 410, 1, 0, 0, 0, 412,
		411, 1, 0, 0, 0, 413, 81, 1, 0, 0, 0, 414, 422, 5, 64, 0, 0, 415, 422,
		5, 65, 0, 0, 416, 422, 5, 66, 0, 0, 417, 422, 5, 67, 0, 0, 418, 422, 5,
		63, 0, 0, 419, 422, 3, 84, 42, 0, 420, 422, 3, 86, 43, 0, 421, 414, 1,
		0, 0, 0, 421, 415, 1, 0, 0, 0, 421, 416, 1, 0, 0, 0, 421, 417, 1, 0, 0,
		0, 421, 418, 1, 0, 0, 0, 421, 419, 1, 0, 0, 0, 421, 420, 1, 0, 0, 0, 422,
		83, 1, 0, 0, 0, 423, 432, 5, 54, 0, 0, 424, 429, 3, 60, 30, 0, 425, 426,
		5, 58, 0, 0, 426, 428, 3, 60, 30, 0, 427, 425, 1, 0, 0, 0, 428, 431, 1,
		0, 0, 0, 429, 427, 1, 0, 0, 0, 429, 430, 1, 0, 0, 0, 430, 433, 1, 0, 0,
		0, 431, 429, 1, 0, 0, 0, 432, 424, 1, 0, 0, 0, 432, 433, 1, 0, 0, 0, 433,
		434, 1, 0, 0, 0, 434, 435, 5, 55, 0, 0, 435, 85, 1, 0, 0, 0, 436, 445,
		5, 54, 0, 0, 437, 442, 3, 88, 44, 0, 438, 439, 5, 58, 0, 0, 439, 441, 3,
		88, 44, 0, 440, 438, 1, 0, 0, 0, 441, 444, 1, 0, 0, 0, 442, 440, 1, 0,
		0, 0, 442, 443, 1, 0, 0, 0, 443, 446, 1, 0, 0, 0, 444, 442, 1, 0, 0, 0,
		445, 437, 1, 0, 0, 0, 445, 446, 1, 0, 0, 0, 446, 447, 1, 0, 0, 0, 447,
		448, 5, 55, 0, 0, 448, 87, 1, 0, 0, 0, 449, 450, 3, 60, 30, 0, 450, 451,
		5, 59, 0, 0, 451, 452, 3, 60, 30, 0, 452, 89, 1, 0, 0, 0, 453, 454, 5,
		68, 0, 0, 454, 463, 5, 54, 0, 0, 455, 460, 3, 92, 46, 0, 456, 457, 5, 58,
		0, 0, 457, 459, 3, 92, 46, 0, 458, 456, 1, 0, 0, 0, 459, 462, 1, 0, 0,
		0, 460, 458, 1, 0, 0, 0, 460, 461, 1, 0, 0, 0, 461, 464, 1, 0, 0, 0, 462,
		460, 1, 0, 0, 0, 463, 455, 1, 0, 0, 0, 463, 464, 1, 0, 0, 0, 464, 465,
		1, 0, 0, 0, 465, 466, 5, 55, 0, 0, 466, 91, 1, 0, 0, 0, 467, 468, 5, 68,
		0, 0, 468, 469, 5, 59, 0, 0, 469, 470, 3, 60, 30, 0, 470, 93, 1, 0, 0,
		0, 471, 476, 3, 60, 30, 0, 472, 473, 5, 58, 0, 0, 473, 475, 3, 60, 30,
		0, 474, 472, 1, 0, 0, 0, 475, 478, 1, 0, 0, 0, 476, 474, 1, 0, 0, 0, 476,
		477, 1, 0, 0, 0, 477, 95, 1, 0, 0, 0, 478, 476, 1, 0, 0, 0, 479, 480, 5,
		33, 0, 0, 480, 481, 5, 41, 0, 0, 481, 482, 3, 32, 16, 0, 482, 483, 5, 43,
		0, 0, 483, 484, 5, 52, 0, 0, 484, 485, 3, 60, 30, 0, 485, 486, 5, 53, 0,
		0, 486, 97, 1, 0, 0, 0, 487, 488, 5, 32, 0, 0, 488, 489, 5, 52, 0, 0, 489,
		492, 3, 32, 16, 0, 490, 491, 5, 58, 0, 0, 491, 493, 3, 60, 30, 0, 492,
		490, 1, 0, 0, 0, 492, 493, 1, 0, 0, 0, 493, 494, 1, 0, 0, 0, 494, 495,
		5, 53, 0, 0, 495, 99, 1, 0, 0, 0, 48, 103, 105, 117, 121, 133, 141, 152,
		156, 163, 168, 171, 177, 181, 190, 195, 198, 210, 223, 232, 243, 269, 282,
		295, 301, 313, 318, 323, 332, 340, 348, 356, 364, 372, 378, 384, 393, 398,
		401, 412, 421, 429, 432, 442, 445, 460, 463, 476, 492,
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
	ArcParserDEFER           = 10
	ArcParserEXTERN          = 11
	ArcParserSELF            = 12
	ArcParserINT8            = 13
	ArcParserINT16           = 14
	ArcParserINT32           = 15
	ArcParserINT64           = 16
	ArcParserUINT8           = 17
	ArcParserUINT16          = 18
	ArcParserUINT32          = 19
	ArcParserUINT64          = 20
	ArcParserUSIZE           = 21
	ArcParserISIZE           = 22
	ArcParserFLOAT32         = 23
	ArcParserFLOAT64         = 24
	ArcParserBYTE            = 25
	ArcParserBOOL            = 26
	ArcParserCHAR            = 27
	ArcParserSTRING          = 28
	ArcParserVOID            = 29
	ArcParserVECTOR          = 30
	ArcParserMAP             = 31
	ArcParserALLOCA          = 32
	ArcParserCAST            = 33
	ArcParserPLUS            = 34
	ArcParserMINUS           = 35
	ArcParserSTAR            = 36
	ArcParserSLASH           = 37
	ArcParserPERCENT         = 38
	ArcParserEQ              = 39
	ArcParserNE              = 40
	ArcParserLT              = 41
	ArcParserLE              = 42
	ArcParserGT              = 43
	ArcParserGE              = 44
	ArcParserAND             = 45
	ArcParserOR              = 46
	ArcParserNOT             = 47
	ArcParserAMP             = 48
	ArcParserAT              = 49
	ArcParserASSIGN          = 50
	ArcParserARROW           = 51
	ArcParserLPAREN          = 52
	ArcParserRPAREN          = 53
	ArcParserLBRACE          = 54
	ArcParserRBRACE          = 55
	ArcParserLBRACKET        = 56
	ArcParserRBRACKET        = 57
	ArcParserCOMMA           = 58
	ArcParserCOLON           = 59
	ArcParserSEMICOLON       = 60
	ArcParserDOT             = 61
	ArcParserELLIPSIS        = 62
	ArcParserBOOLEAN_LITERAL = 63
	ArcParserINTEGER_LITERAL = 64
	ArcParserFLOAT_LITERAL   = 65
	ArcParserSTRING_LITERAL  = 66
	ArcParserCHAR_LITERAL    = 67
	ArcParserIDENTIFIER      = 68
	ArcParserWS              = 69
	ArcParserLINE_COMMENT    = 70
	ArcParserBLOCK_COMMENT   = 71
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
	ArcParserRULE_deferStmt                = 29
	ArcParserRULE_expression               = 30
	ArcParserRULE_logicalOrExpression      = 31
	ArcParserRULE_logicalAndExpression     = 32
	ArcParserRULE_equalityExpression       = 33
	ArcParserRULE_relationalExpression     = 34
	ArcParserRULE_additiveExpression       = 35
	ArcParserRULE_multiplicativeExpression = 36
	ArcParserRULE_unaryExpression          = 37
	ArcParserRULE_postfixExpression        = 38
	ArcParserRULE_postfixOp                = 39
	ArcParserRULE_primaryExpression        = 40
	ArcParserRULE_literal                  = 41
	ArcParserRULE_vectorLiteral            = 42
	ArcParserRULE_mapLiteral               = 43
	ArcParserRULE_mapEntry                 = 44
	ArcParserRULE_structLiteral            = 45
	ArcParserRULE_fieldInit                = 46
	ArcParserRULE_argumentList             = 47
	ArcParserRULE_castExpression           = 48
	ArcParserRULE_allocaExpression         = 49
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
	p.SetState(105)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2174) != 0 {
		p.SetState(103)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case ArcParserIMPORT:
			{
				p.SetState(100)
				p.ImportDecl()
			}

		case ArcParserNAMESPACE:
			{
				p.SetState(101)
				p.NamespaceDecl()
			}

		case ArcParserLET, ArcParserCONST, ArcParserFUNC, ArcParserSTRUCT, ArcParserEXTERN:
			{
				p.SetState(102)
				p.TopLevelDecl()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(107)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(108)
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

	p.SetState(121)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(110)
			p.Match(ArcParserIMPORT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(111)
			p.Match(ArcParserSTRING_LITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(112)
			p.Match(ArcParserIMPORT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(113)
			p.Match(ArcParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(117)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == ArcParserSTRING_LITERAL {
			{
				p.SetState(114)
				p.ImportSpec()
			}

			p.SetState(119)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(120)
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
		p.SetState(123)
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
		p.SetState(125)
		p.Match(ArcParserNAMESPACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(126)
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
	p.SetState(133)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ArcParserFUNC:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(128)
			p.FunctionDecl()
		}

	case ArcParserSTRUCT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(129)
			p.StructDecl()
		}

	case ArcParserLET:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(130)
			p.VariableDecl()
		}

	case ArcParserCONST:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(131)
			p.ConstDecl()
		}

	case ArcParserEXTERN:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(132)
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
		p.SetState(135)
		p.Match(ArcParserEXTERN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(136)
		p.Match(ArcParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(137)
		p.Match(ArcParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(141)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ArcParserFUNC {
		{
			p.SetState(138)
			p.ExternMember()
		}

		p.SetState(143)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(144)
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
		p.SetState(146)
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
		p.SetState(148)
		p.Match(ArcParserFUNC)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(149)
		p.Match(ArcParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(150)
		p.Match(ArcParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(152)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64((_la-13)) & ^0x3f) == 0 && ((int64(1)<<(_la-13))&36591781341036543) != 0 {
		{
			p.SetState(151)
			p.ExternParameterList()
		}

	}
	{
		p.SetState(154)
		p.Match(ArcParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(156)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64((_la-13)) & ^0x3f) == 0 && ((int64(1)<<(_la-13))&36028831387615231) != 0 {
		{
			p.SetState(155)
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

	p.SetState(171)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ArcParserINT8, ArcParserINT16, ArcParserINT32, ArcParserINT64, ArcParserUINT8, ArcParserUINT16, ArcParserUINT32, ArcParserUINT64, ArcParserUSIZE, ArcParserISIZE, ArcParserFLOAT32, ArcParserFLOAT64, ArcParserBYTE, ArcParserBOOL, ArcParserCHAR, ArcParserSTRING, ArcParserVOID, ArcParserVECTOR, ArcParserMAP, ArcParserSTAR, ArcParserAMP, ArcParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(158)
			p.Type_()
		}
		p.SetState(163)
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
					p.SetState(159)
					p.Match(ArcParserCOMMA)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(160)
					p.Type_()
				}

			}
			p.SetState(165)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}
		p.SetState(168)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == ArcParserCOMMA {
			{
				p.SetState(166)
				p.Match(ArcParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(167)
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
			p.SetState(170)
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
		p.SetState(173)
		p.Match(ArcParserFUNC)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(174)
		p.Match(ArcParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(175)
		p.Match(ArcParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(177)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ArcParserELLIPSIS || _la == ArcParserIDENTIFIER {
		{
			p.SetState(176)
			p.ParameterList()
		}

	}
	{
		p.SetState(179)
		p.Match(ArcParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(181)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64((_la-13)) & ^0x3f) == 0 && ((int64(1)<<(_la-13))&36028831387615231) != 0 {
		{
			p.SetState(180)
			p.Type_()
		}

	}
	{
		p.SetState(183)
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

	p.SetState(198)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ArcParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(185)
			p.Parameter()
		}
		p.SetState(190)
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
					p.SetState(186)
					p.Match(ArcParserCOMMA)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(187)
					p.Parameter()
				}

			}
			p.SetState(192)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 13, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}
		p.SetState(195)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == ArcParserCOMMA {
			{
				p.SetState(193)
				p.Match(ArcParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(194)
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
			p.SetState(197)
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
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(200)
		p.Match(ArcParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(201)
		p.Match(ArcParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(202)
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
		p.SetState(204)
		p.Match(ArcParserSTRUCT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(205)
		p.Match(ArcParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(206)
		p.Match(ArcParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(210)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ArcParserIDENTIFIER {
		{
			p.SetState(207)
			p.StructField()
		}

		p.SetState(212)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(213)
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
		p.SetState(215)
		p.Match(ArcParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(216)
		p.Match(ArcParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(217)
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
		p.SetState(219)
		p.Match(ArcParserLET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(220)
		p.Match(ArcParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(223)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ArcParserCOLON {
		{
			p.SetState(221)
			p.Match(ArcParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(222)
			p.Type_()
		}

	}
	{
		p.SetState(225)
		p.Match(ArcParserASSIGN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(226)
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
		p.SetState(228)
		p.Match(ArcParserCONST)
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
	p.SetState(243)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ArcParserINT8, ArcParserINT16, ArcParserINT32, ArcParserINT64, ArcParserUINT8, ArcParserUINT16, ArcParserUINT32, ArcParserUINT64, ArcParserUSIZE, ArcParserISIZE, ArcParserFLOAT32, ArcParserFLOAT64, ArcParserBYTE, ArcParserBOOL, ArcParserCHAR, ArcParserSTRING, ArcParserVOID:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(237)
			p.PrimitiveType()
		}

	case ArcParserSTAR:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(238)
			p.PointerType()
		}

	case ArcParserAMP:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(239)
			p.ReferenceType()
		}

	case ArcParserVECTOR:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(240)
			p.VectorType()
		}

	case ArcParserMAP:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(241)
			p.MapType()
		}

	case ArcParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(242)
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
		p.SetState(245)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1073733632) != 0) {
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
		p.SetState(247)
		p.Match(ArcParserSTAR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(248)
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
		p.SetState(250)
		p.Match(ArcParserAMP)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(251)
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
		p.SetState(253)
		p.Match(ArcParserVECTOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(254)
		p.Match(ArcParserLT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(255)
		p.Type_()
	}
	{
		p.SetState(256)
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
		p.SetState(258)
		p.Match(ArcParserMAP)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(259)
		p.Match(ArcParserLT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(260)
		p.Type_()
	}
	{
		p.SetState(261)
		p.Match(ArcParserCOMMA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(262)
		p.Type_()
	}
	{
		p.SetState(263)
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
		p.SetState(265)
		p.Match(ArcParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(269)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&-9200431710288738920) != 0) || ((int64((_la-64)) & ^0x3f) == 0 && ((int64(1)<<(_la-64))&31) != 0) {
		{
			p.SetState(266)
			p.Statement()
		}

		p.SetState(271)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(272)
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
	p.SetState(282)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 21, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(274)
			p.VariableDecl()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(275)
			p.ConstDecl()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(276)
			p.AssignmentStmt()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(277)
			p.ExpressionStmt()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(278)
			p.ReturnStmt()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(279)
			p.IfStmt()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(280)
			p.DeferStmt()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(281)
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
		p.SetState(284)
		p.LeftHandSide()
	}
	{
		p.SetState(285)
		p.Match(ArcParserASSIGN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(286)
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
	p.SetState(295)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 22, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(288)
			p.Match(ArcParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(289)
			p.Match(ArcParserSTAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(290)
			p.Expression()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(291)
			p.Expression()
		}
		{
			p.SetState(292)
			p.Match(ArcParserDOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(293)
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
		p.SetState(297)
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
		p.SetState(299)
		p.Match(ArcParserRETURN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(301)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 23, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(300)
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
		p.SetState(303)
		p.Match(ArcParserIF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(304)
		p.Expression()
	}
	{
		p.SetState(305)
		p.Block()
	}
	p.SetState(313)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 24, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(306)
				p.Match(ArcParserELSE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(307)
				p.Match(ArcParserIF)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(308)
				p.Expression()
			}
			{
				p.SetState(309)
				p.Block()
			}

		}
		p.SetState(315)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 24, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(318)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ArcParserELSE {
		{
			p.SetState(316)
			p.Match(ArcParserELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(317)
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
	p.EnterRule(localctx, 58, ArcParserRULE_deferStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(320)
		p.Match(ArcParserDEFER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(323)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 26, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(321)
			p.AssignmentStmt()
		}

	case 2:
		{
			p.SetState(322)
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
	p.EnterRule(localctx, 60, ArcParserRULE_expression)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(325)
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
	p.EnterRule(localctx, 62, ArcParserRULE_logicalOrExpression)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(327)
		p.LogicalAndExpression()
	}
	p.SetState(332)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ArcParserOR {
		{
			p.SetState(328)
			p.Match(ArcParserOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(329)
			p.LogicalAndExpression()
		}

		p.SetState(334)
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
	p.EnterRule(localctx, 64, ArcParserRULE_logicalAndExpression)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(335)
		p.EqualityExpression()
	}
	p.SetState(340)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ArcParserAND {
		{
			p.SetState(336)
			p.Match(ArcParserAND)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(337)
			p.EqualityExpression()
		}

		p.SetState(342)
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
	p.EnterRule(localctx, 66, ArcParserRULE_equalityExpression)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(343)
		p.RelationalExpression()
	}
	p.SetState(348)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ArcParserEQ || _la == ArcParserNE {
		{
			p.SetState(344)
			_la = p.GetTokenStream().LA(1)

			if !(_la == ArcParserEQ || _la == ArcParserNE) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(345)
			p.RelationalExpression()
		}

		p.SetState(350)
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
	p.EnterRule(localctx, 68, ArcParserRULE_relationalExpression)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(351)
		p.AdditiveExpression()
	}
	p.SetState(356)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&32985348833280) != 0 {
		{
			p.SetState(352)
			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&32985348833280) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(353)
			p.AdditiveExpression()
		}

		p.SetState(358)
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
	p.EnterRule(localctx, 70, ArcParserRULE_additiveExpression)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(359)
		p.MultiplicativeExpression()
	}
	p.SetState(364)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 31, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(360)
				_la = p.GetTokenStream().LA(1)

				if !(_la == ArcParserPLUS || _la == ArcParserMINUS) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}
			{
				p.SetState(361)
				p.MultiplicativeExpression()
			}

		}
		p.SetState(366)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 31, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 72, ArcParserRULE_multiplicativeExpression)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(367)
		p.UnaryExpression()
	}
	p.SetState(372)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 32, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(368)
				_la = p.GetTokenStream().LA(1)

				if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&481036337152) != 0) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}
			{
				p.SetState(369)
				p.UnaryExpression()
			}

		}
		p.SetState(374)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 32, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 74, ArcParserRULE_unaryExpression)
	var _la int

	p.SetState(378)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ArcParserMINUS, ArcParserSTAR, ArcParserNOT, ArcParserAMP:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(375)
			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&422315544281088) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(376)
			p.UnaryExpression()
		}

	case ArcParserALLOCA, ArcParserCAST, ArcParserLPAREN, ArcParserLBRACE, ArcParserBOOLEAN_LITERAL, ArcParserINTEGER_LITERAL, ArcParserFLOAT_LITERAL, ArcParserSTRING_LITERAL, ArcParserCHAR_LITERAL, ArcParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(377)
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
	p.EnterRule(localctx, 76, ArcParserRULE_postfixExpression)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(380)
		p.PrimaryExpression()
	}
	p.SetState(384)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 34, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(381)
				p.PostfixOp()
			}

		}
		p.SetState(386)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 34, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 78, ArcParserRULE_postfixOp)
	var _la int

	p.SetState(401)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 37, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(387)
			p.Match(ArcParserDOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(388)
			p.Match(ArcParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(389)
			p.Match(ArcParserDOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(390)
			p.Match(ArcParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(391)
			p.Match(ArcParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(393)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64((_la-32)) & ^0x3f) == 0 && ((int64(1)<<(_la-32))&135296811035) != 0 {
			{
				p.SetState(392)
				p.ArgumentList()
			}

		}
		{
			p.SetState(395)
			p.Match(ArcParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(396)
			p.Match(ArcParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(398)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64((_la-32)) & ^0x3f) == 0 && ((int64(1)<<(_la-32))&135296811035) != 0 {
			{
				p.SetState(397)
				p.ArgumentList()
			}

		}
		{
			p.SetState(400)
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
	p.EnterRule(localctx, 80, ArcParserRULE_primaryExpression)
	p.SetState(412)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 38, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(403)
			p.Literal()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(404)
			p.Match(ArcParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(405)
			p.Match(ArcParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(406)
			p.Expression()
		}
		{
			p.SetState(407)
			p.Match(ArcParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(409)
			p.CastExpression()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(410)
			p.AllocaExpression()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(411)
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
	p.EnterRule(localctx, 82, ArcParserRULE_literal)
	p.SetState(421)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 39, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(414)
			p.Match(ArcParserINTEGER_LITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(415)
			p.Match(ArcParserFLOAT_LITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(416)
			p.Match(ArcParserSTRING_LITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(417)
			p.Match(ArcParserCHAR_LITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(418)
			p.Match(ArcParserBOOLEAN_LITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(419)
			p.VectorLiteral()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(420)
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
	p.EnterRule(localctx, 84, ArcParserRULE_vectorLiteral)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(423)
		p.Match(ArcParserLBRACE)
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

	if (int64((_la-32)) & ^0x3f) == 0 && ((int64(1)<<(_la-32))&135296811035) != 0 {
		{
			p.SetState(424)
			p.Expression()
		}
		p.SetState(429)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == ArcParserCOMMA {
			{
				p.SetState(425)
				p.Match(ArcParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(426)
				p.Expression()
			}

			p.SetState(431)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(434)
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
	p.EnterRule(localctx, 86, ArcParserRULE_mapLiteral)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(436)
		p.Match(ArcParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(445)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64((_la-32)) & ^0x3f) == 0 && ((int64(1)<<(_la-32))&135296811035) != 0 {
		{
			p.SetState(437)
			p.MapEntry()
		}
		p.SetState(442)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == ArcParserCOMMA {
			{
				p.SetState(438)
				p.Match(ArcParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(439)
				p.MapEntry()
			}

			p.SetState(444)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(447)
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
	p.EnterRule(localctx, 88, ArcParserRULE_mapEntry)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(449)
		p.Expression()
	}
	{
		p.SetState(450)
		p.Match(ArcParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(451)
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
	p.EnterRule(localctx, 90, ArcParserRULE_structLiteral)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(453)
		p.Match(ArcParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(454)
		p.Match(ArcParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(463)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ArcParserIDENTIFIER {
		{
			p.SetState(455)
			p.FieldInit()
		}
		p.SetState(460)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == ArcParserCOMMA {
			{
				p.SetState(456)
				p.Match(ArcParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(457)
				p.FieldInit()
			}

			p.SetState(462)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(465)
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
	p.EnterRule(localctx, 92, ArcParserRULE_fieldInit)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(467)
		p.Match(ArcParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(468)
		p.Match(ArcParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(469)
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
	p.EnterRule(localctx, 94, ArcParserRULE_argumentList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(471)
		p.Expression()
	}
	p.SetState(476)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ArcParserCOMMA {
		{
			p.SetState(472)
			p.Match(ArcParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(473)
			p.Expression()
		}

		p.SetState(478)
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
	p.EnterRule(localctx, 96, ArcParserRULE_castExpression)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(479)
		p.Match(ArcParserCAST)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(480)
		p.Match(ArcParserLT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(481)
		p.Type_()
	}
	{
		p.SetState(482)
		p.Match(ArcParserGT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(483)
		p.Match(ArcParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(484)
		p.Expression()
	}
	{
		p.SetState(485)
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
	p.EnterRule(localctx, 98, ArcParserRULE_allocaExpression)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(487)
		p.Match(ArcParserALLOCA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(488)
		p.Match(ArcParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(489)
		p.Type_()
	}
	p.SetState(492)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ArcParserCOMMA {
		{
			p.SetState(490)
			p.Match(ArcParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(491)
			p.Expression()
		}

	}
	{
		p.SetState(494)
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
