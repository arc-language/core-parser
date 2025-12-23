package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	"github.com/arc-language/core-parser"
)

// CustomErrorListener implements ANTLR error listener for better error reporting
type CustomErrorListener struct {
	*antlr.DefaultErrorListener
	Errors []string
}

func NewCustomErrorListener() *CustomErrorListener {
	return &CustomErrorListener{
		DefaultErrorListener: antlr.NewDefaultErrorListener(),
		Errors:               make([]string, 0),
	}
}

func (c *CustomErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	errorMsg := fmt.Sprintf("line %d:%d %s", line, column, msg)
	c.Errors = append(c.Errors, errorMsg)
	fmt.Fprintf(os.Stderr, "ERROR: %s\n", errorMsg)
}

// ArcPrintListener is a custom listener that prints the parse tree structure
type ArcPrintListener struct {
	*parser.BaseArcParserListener
	indent int
}

func NewArcPrintListener() *ArcPrintListener {
	return &ArcPrintListener{
		BaseArcParserListener: &parser.BaseArcParserListener{},
		indent:                0,
	}
}

func (l *ArcPrintListener) printIndent(s string) {
	for i := 0; i < l.indent; i++ {
		fmt.Print("  ")
	}
	fmt.Println(s)
}

func (l *ArcPrintListener) EnterProgram(ctx *parser.ProgramContext) {
	l.printIndent("Program")
	l.indent++
}

func (l *ArcPrintListener) ExitProgram(ctx *parser.ProgramContext) {
	l.indent--
}

func (l *ArcPrintListener) EnterFunctionDecl(ctx *parser.FunctionDeclContext) {
	funcName := ctx.IDENTIFIER().GetText()
	returnType := ""
	if ctx.ReturnType() != nil {
		returnType = fmt.Sprintf(" -> %s", ctx.ReturnType().GetText())
	}
	l.printIndent(fmt.Sprintf("Function: %s%s", funcName, returnType))
	l.indent++
}

func (l *ArcPrintListener) ExitFunctionDecl(ctx *parser.FunctionDeclContext) {
	l.indent--
}

func (l *ArcPrintListener) EnterVariableDecl(ctx *parser.VariableDeclContext) {
	varName := ctx.IDENTIFIER().GetText()
	keyword := "let"
	if ctx.CONST() != nil {
		keyword = "const"
	}
	typeInfo := ""
	if ctx.TypeSpec() != nil {
		typeInfo = fmt.Sprintf(": %s", ctx.TypeSpec().GetText())
	}
	l.printIndent(fmt.Sprintf("Variable Declaration (%s): %s%s", keyword, varName, typeInfo))
	l.indent++
}

func (l *ArcPrintListener) ExitVariableDecl(ctx *parser.VariableDeclContext) {
	l.indent--
}

func (l *ArcPrintListener) EnterStructDecl(ctx *parser.StructDeclContext) {
	structName := ctx.IDENTIFIER().GetText()
	l.printIndent(fmt.Sprintf("Struct: %s", structName))
	l.indent++
}

func (l *ArcPrintListener) ExitStructDecl(ctx *parser.StructDeclContext) {
	l.indent--
}

func (l *ArcPrintListener) EnterParameter(ctx *parser.ParameterContext) {
	paramName := ctx.IDENTIFIER().GetText()
	typeInfo := ""
	if ctx.TypeSpec() != nil {
		typeInfo = fmt.Sprintf(": %s", ctx.TypeSpec().GetText())
	}
	l.printIndent(fmt.Sprintf("Parameter: %s%s", paramName, typeInfo))
}

func (l *ArcPrintListener) EnterReturnStmt(ctx *parser.ReturnStmtContext) {
	l.printIndent("Return Statement")
	l.indent++
}

func (l *ArcPrintListener) ExitReturnStmt(ctx *parser.ReturnStmtContext) {
	l.indent--
}

func (l *ArcPrintListener) EnterIfStmt(ctx *parser.IfStmtContext) {
	l.printIndent("If Statement")
	l.indent++
}

func (l *ArcPrintListener) ExitIfStmt(ctx *parser.IfStmtContext) {
	l.indent--
}

func (l *ArcPrintListener) EnterAssignStmt(ctx *parser.AssignStmtContext) {
	varName := ctx.IDENTIFIER().GetText()
	l.printIndent(fmt.Sprintf("Assignment: %s", varName))
	l.indent++
}

func (l *ArcPrintListener) ExitAssignStmt(ctx *parser.AssignStmtContext) {
	l.indent--
}

func (l *ArcPrintListener) EnterCallExpr(ctx *parser.CallExprContext) {
	l.printIndent("Function Call")
	l.indent++
}

func (l *ArcPrintListener) ExitCallExpr(ctx *parser.CallExprContext) {
	l.indent--
}

func (l *ArcPrintListener) EnterIntLiteral(ctx *parser.IntLiteralContext) {
	value := ctx.INTEGER_LITERAL().GetText()
	l.printIndent(fmt.Sprintf("Int Literal: %s", value))
}

func (l *ArcPrintListener) EnterFloatLiteral(ctx *parser.FloatLiteralContext) {
	value := ctx.FLOAT_LITERAL().GetText()
	l.printIndent(fmt.Sprintf("Float Literal: %s", value))
}

func (l *ArcPrintListener) EnterStringLiteral(ctx *parser.StringLiteralContext) {
	value := ctx.STRING_LITERAL().GetText()
	l.printIndent(fmt.Sprintf("String Literal: %s", value))
}

func (l *ArcPrintListener) EnterBoolLiteral(ctx *parser.BoolLiteralContext) {
	value := ctx.GetText()
	l.printIndent(fmt.Sprintf("Bool Literal: %s", value))
}

func (l *ArcPrintListener) EnterIdentifierExpr(ctx *parser.IdentifierExprContext) {
	name := ctx.IDENTIFIER().GetText()
	l.printIndent(fmt.Sprintf("Identifier: %s", name))
}

// ParseArcFile parses an Arc language source file
func ParseArcFile(filename string) (antlr.Tree, error) {
	// Read the input file
	input, err := antlr.NewFileStream(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %v", err)
	}

	// Create the lexer
	lexer := parser.NewArcLexer(input)
	
	// Remove default error listeners and add custom one
	lexer.RemoveErrorListeners()
	errorListener := NewCustomErrorListener()
	lexer.AddErrorListener(errorListener)

	// Create token stream
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the parser
	p := parser.NewArcParser(stream)
	
	// Remove default error listeners and add custom one
	p.RemoveErrorListeners()
	p.AddErrorListener(errorListener)

	// Parse the program (starting rule)
	tree := p.Program()

	// Check for parsing errors
	if len(errorListener.Errors) > 0 {
		return tree, fmt.Errorf("parsing failed with %d error(s)", len(errorListener.Errors))
	}

	return tree, nil
}

// WalkParseTree walks the parse tree using a listener
func WalkParseTree(tree antlr.Tree, listener antlr.ParseTreeListener) {
	antlr.ParseTreeWalkerDefault.Walk(listener, tree)
}

func main() {
	// Check command line arguments
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <arc-source-file>\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Example: %s example.arc\n", os.Args[0])
		os.Exit(1)
	}

	filename := os.Args[1]

	// Check if file exists
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		fmt.Fprintf(os.Stderr, "Error: File '%s' does not exist\n", filename)
		os.Exit(1)
	}

	fmt.Printf("Parsing Arc file: %s\n", filename)
	fmt.Println(strings.Repeat("=", 60))

	// Parse the file
	tree, err := ParseArcFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Parse error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("\nParse tree structure:")
	fmt.Println(strings.Repeat("-", 60))

	// Create listener and walk the tree
	listener := NewArcPrintListener()
	WalkParseTree(tree, listener)

	fmt.Println(strings.Repeat("=", 60))
	fmt.Println("✓ Parsing completed successfully!")
}