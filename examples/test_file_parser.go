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

// ArcPrintListener is a comprehensive listener that prints all parse tree nodes
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

// ============================================================================
// COMPILATION UNIT
// ============================================================================

func (l *ArcPrintListener) EnterCompilationUnit(ctx *parser.CompilationUnitContext) {
	l.printIndent("╔═══ Compilation Unit ═══")
	l.indent++
}

func (l *ArcPrintListener) ExitCompilationUnit(ctx *parser.CompilationUnitContext) {
	l.indent--
	l.printIndent("╚═══ End Compilation Unit ═══")
}

// ============================================================================
// IMPORTS
// ============================================================================

func (l *ArcPrintListener) EnterImportDecl(ctx *parser.ImportDeclContext) {
	if ctx.STRING_LITERAL() != nil {
		l.printIndent(fmt.Sprintf("📦 Import: %s", ctx.STRING_LITERAL().GetText()))
	} else {
		l.printIndent("📦 Import (multiple):")
		l.indent++
	}
}

func (l *ArcPrintListener) ExitImportDecl(ctx *parser.ImportDeclContext) {
	if ctx.LPAREN() != nil {
		l.indent--
	}
}

func (l *ArcPrintListener) EnterImportSpec(ctx *parser.ImportSpecContext) {
	l.printIndent(fmt.Sprintf("  - %s", ctx.STRING_LITERAL().GetText()))
}

// ============================================================================
// NAMESPACE
// ============================================================================

func (l *ArcPrintListener) EnterNamespaceDecl(ctx *parser.NamespaceDeclContext) {
	l.printIndent(fmt.Sprintf("🏛️  Namespace: %s", ctx.IDENTIFIER().GetText()))
}

// ============================================================================
// EXTERN DECLARATIONS
// ============================================================================

func (l *ArcPrintListener) EnterExternDecl(ctx *parser.ExternDeclContext) {
	l.printIndent(fmt.Sprintf("🔗 Extern Block: %s", ctx.IDENTIFIER().GetText()))
	l.indent++
}

func (l *ArcPrintListener) ExitExternDecl(ctx *parser.ExternDeclContext) {
	l.indent--
}

func (l *ArcPrintListener) EnterExternFunctionDecl(ctx *parser.ExternFunctionDeclContext) {
	funcName := ctx.IDENTIFIER().GetText()
	returnType := "void"
	if ctx.Type_() != nil {
		returnType = ctx.Type_().GetText()
	}
	params := ""
	if ctx.ExternParameterList() != nil {
		params = ctx.ExternParameterList().GetText()
	}
	l.printIndent(fmt.Sprintf("  ⚡ extern func %s(%s) -> %s", funcName, params, returnType))
}

func (l *ArcPrintListener) EnterExternParameterList(ctx *parser.ExternParameterListContext) {
	if ctx.ELLIPSIS() != nil {
		l.printIndent("    ... (variadic)")
	}
}

// ============================================================================
// FUNCTION DECLARATIONS
// ============================================================================

func (l *ArcPrintListener) EnterFunctionDecl(ctx *parser.FunctionDeclContext) {
	funcName := ctx.IDENTIFIER().GetText()
	returnType := "void"
	if ctx.Type_() != nil {
		returnType = ctx.Type_().GetText()
	}
	l.printIndent(fmt.Sprintf("⚡ Function: %s -> %s", funcName, returnType))
	l.indent++
}

func (l *ArcPrintListener) ExitFunctionDecl(ctx *parser.FunctionDeclContext) {
	l.indent--
}

func (l *ArcPrintListener) EnterParameterList(ctx *parser.ParameterListContext) {
	if ctx.ELLIPSIS() != nil {
		l.printIndent("  📝 Parameters: (variadic)")
	} else if len(ctx.AllParameter()) > 0 {
		l.printIndent("  📝 Parameters:")
	}
}

func (l *ArcPrintListener) EnterParameter(ctx *parser.ParameterContext) {
	isSelf := ctx.SELF() != nil
	paramName := ctx.IDENTIFIER().GetText()
	paramType := ctx.Type_().GetText()
	selfPrefix := ""
	if isSelf {
		selfPrefix = "self "
	}
	l.printIndent(fmt.Sprintf("    - %s%s: %s", selfPrefix, paramName, paramType))
}

// ============================================================================
// STRUCT DECLARATIONS
// ============================================================================

func (l *ArcPrintListener) EnterStructDecl(ctx *parser.StructDeclContext) {
	structName := ctx.IDENTIFIER().GetText()
	l.printIndent(fmt.Sprintf("🏗️  Struct: %s", structName))
	l.indent++
}

func (l *ArcPrintListener) ExitStructDecl(ctx *parser.StructDeclContext) {
	l.indent--
}

func (l *ArcPrintListener) EnterStructField(ctx *parser.StructFieldContext) {
	fieldName := ctx.IDENTIFIER().GetText()
	fieldType := ctx.Type_().GetText()
	l.printIndent(fmt.Sprintf("  • %s: %s", fieldName, fieldType))
}

// ============================================================================
// VARIABLE AND CONSTANT DECLARATIONS
// ============================================================================

func (l *ArcPrintListener) EnterVariableDecl(ctx *parser.VariableDeclContext) {
	varName := ctx.IDENTIFIER().GetText()
	typeInfo := "<inferred>"
	if ctx.Type_() != nil {
		typeInfo = ctx.Type_().GetText()
	}
	l.printIndent(fmt.Sprintf("📌 let %s: %s", varName, typeInfo))
	l.indent++
}

func (l *ArcPrintListener) ExitVariableDecl(ctx *parser.VariableDeclContext) {
	l.indent--
}

func (l *ArcPrintListener) EnterConstDecl(ctx *parser.ConstDeclContext) {
	constName := ctx.IDENTIFIER().GetText()
	typeInfo := "<inferred>"
	if ctx.Type_() != nil {
		typeInfo = ctx.Type_().GetText()
	}
	l.printIndent(fmt.Sprintf("🔒 const %s: %s", constName, typeInfo))
	l.indent++
}

func (l *ArcPrintListener) ExitConstDecl(ctx *parser.ConstDeclContext) {
	l.indent--
}

// ============================================================================
// TYPES
// ============================================================================

func (l *ArcPrintListener) EnterPrimitiveType(ctx *parser.PrimitiveTypeContext) {
	l.printIndent(fmt.Sprintf("  🔤 Primitive Type: %s", ctx.GetText()))
}

func (l *ArcPrintListener) EnterPointerType(ctx *parser.PointerTypeContext) {
	l.printIndent(fmt.Sprintf("  👉 Pointer Type: *%s", ctx.Type_().GetText()))
}

func (l *ArcPrintListener) EnterReferenceType(ctx *parser.ReferenceTypeContext) {
	l.printIndent(fmt.Sprintf("  📎 Reference Type: &%s", ctx.Type_().GetText()))
}

func (l *ArcPrintListener) EnterVectorType(ctx *parser.VectorTypeContext) {
	l.printIndent(fmt.Sprintf("  📊 Vector Type: vector<%s>", ctx.Type_().GetText()))
}

func (l *ArcPrintListener) EnterMapType(ctx *parser.MapTypeContext) {
	types := ctx.AllType_()
	l.printIndent(fmt.Sprintf("  🗺️  Map Type: map<%s, %s>", types[0].GetText(), types[1].GetText()))
}

// ============================================================================
// BLOCKS AND STATEMENTS
// ============================================================================

func (l *ArcPrintListener) EnterBlock(ctx *parser.BlockContext) {
	l.printIndent("{ Block")
	l.indent++
}

func (l *ArcPrintListener) ExitBlock(ctx *parser.BlockContext) {
	l.indent--
	l.printIndent("}")
}

func (l *ArcPrintListener) EnterAssignmentStmt(ctx *parser.AssignmentStmtContext) {
	lhs := ctx.LeftHandSide().GetText()
	l.printIndent(fmt.Sprintf("✏️  Assignment: %s =", lhs))
	l.indent++
}

func (l *ArcPrintListener) ExitAssignmentStmt(ctx *parser.AssignmentStmtContext) {
	l.indent--
}

func (l *ArcPrintListener) EnterReturnStmt(ctx *parser.ReturnStmtContext) {
	if ctx.Expression() != nil {
		l.printIndent("↩️  Return:")
		l.indent++
	} else {
		l.printIndent("↩️  Return (void)")
	}
}

func (l *ArcPrintListener) ExitReturnStmt(ctx *parser.ReturnStmtContext) {
	if ctx.Expression() != nil {
		l.indent--
	}
}

func (l *ArcPrintListener) EnterIfStmt(ctx *parser.IfStmtContext) {
	elseIfCount := len(ctx.AllIF()) - 1
	hasElse := len(ctx.AllELSE()) > elseIfCount
	info := ""
	if elseIfCount > 0 {
		info += fmt.Sprintf(" [%d else-if]", elseIfCount)
	}
	if hasElse {
		info += " [else]"
	}
	l.printIndent(fmt.Sprintf("🔀 If Statement%s", info))
	l.indent++
}

func (l *ArcPrintListener) ExitIfStmt(ctx *parser.IfStmtContext) {
	l.indent--
}

func (l *ArcPrintListener) EnterDeferStmt(ctx *parser.DeferStmtContext) {
	l.printIndent("⏳ Defer Statement:")
	l.indent++
}

func (l *ArcPrintListener) ExitDeferStmt(ctx *parser.DeferStmtContext) {
	l.indent--
}

func (l *ArcPrintListener) EnterExpressionStmt(ctx *parser.ExpressionStmtContext) {
	l.printIndent("📊 Expression Statement:")
	l.indent++
}

func (l *ArcPrintListener) ExitExpressionStmt(ctx *parser.ExpressionStmtContext) {
	l.indent--
}

// ============================================================================
// EXPRESSIONS
// ============================================================================

func (l *ArcPrintListener) EnterLogicalOrExpression(ctx *parser.LogicalOrExpressionContext) {
	if len(ctx.AllLogicalAndExpression()) > 1 {
		l.printIndent(fmt.Sprintf("🔗 Logical OR (%d operands)", len(ctx.AllLogicalAndExpression())))
		l.indent++
	}
}

func (l *ArcPrintListener) ExitLogicalOrExpression(ctx *parser.LogicalOrExpressionContext) {
	if len(ctx.AllLogicalAndExpression()) > 1 {
		l.indent--
	}
}

func (l *ArcPrintListener) EnterLogicalAndExpression(ctx *parser.LogicalAndExpressionContext) {
	if len(ctx.AllEqualityExpression()) > 1 {
		l.printIndent(fmt.Sprintf("🔗 Logical AND (%d operands)", len(ctx.AllEqualityExpression())))
		l.indent++
	}
}

func (l *ArcPrintListener) ExitLogicalAndExpression(ctx *parser.LogicalAndExpressionContext) {
	if len(ctx.AllEqualityExpression()) > 1 {
		l.indent--
	}
}

func (l *ArcPrintListener) EnterEqualityExpression(ctx *parser.EqualityExpressionContext) {
	if len(ctx.AllRelationalExpression()) > 1 {
		ops := make([]string, 0)
		for i := 0; i < len(ctx.AllEQ()); i++ {
			ops = append(ops, "==")
		}
		for i := 0; i < len(ctx.AllNE()); i++ {
			ops = append(ops, "!=")
		}
		l.printIndent(fmt.Sprintf("⚖️  Equality [%s]", strings.Join(ops, ", ")))
		l.indent++
	}
}

func (l *ArcPrintListener) ExitEqualityExpression(ctx *parser.EqualityExpressionContext) {
	if len(ctx.AllRelationalExpression()) > 1 {
		l.indent--
	}
}

func (l *ArcPrintListener) EnterRelationalExpression(ctx *parser.RelationalExpressionContext) {
	if len(ctx.AllAdditiveExpression()) > 1 {
		ops := make([]string, 0)
		for i := 0; i < len(ctx.AllLT()); i++ {
			ops = append(ops, "<")
		}
		for i := 0; i < len(ctx.AllLE()); i++ {
			ops = append(ops, "<=")
		}
		for i := 0; i < len(ctx.AllGT()); i++ {
			ops = append(ops, ">")
		}
		for i := 0; i < len(ctx.AllGE()); i++ {
			ops = append(ops, ">=")
		}
		l.printIndent(fmt.Sprintf("📏 Relational [%s]", strings.Join(ops, ", ")))
		l.indent++
	}
}

func (l *ArcPrintListener) ExitRelationalExpression(ctx *parser.RelationalExpressionContext) {
	if len(ctx.AllAdditiveExpression()) > 1 {
		l.indent--
	}
}

func (l *ArcPrintListener) EnterAdditiveExpression(ctx *parser.AdditiveExpressionContext) {
	if len(ctx.AllMultiplicativeExpression()) > 1 {
		ops := make([]string, 0)
		for i := 0; i < len(ctx.AllPLUS()); i++ {
			ops = append(ops, "+")
		}
		for i := 0; i < len(ctx.AllMINUS()); i++ {
			ops = append(ops, "-")
		}
		l.printIndent(fmt.Sprintf("➕ Additive [%s]", strings.Join(ops, ", ")))
		l.indent++
	}
}

func (l *ArcPrintListener) ExitAdditiveExpression(ctx *parser.AdditiveExpressionContext) {
	if len(ctx.AllMultiplicativeExpression()) > 1 {
		l.indent--
	}
}

func (l *ArcPrintListener) EnterMultiplicativeExpression(ctx *parser.MultiplicativeExpressionContext) {
	if len(ctx.AllUnaryExpression()) > 1 {
		ops := make([]string, 0)
		for i := 0; i < len(ctx.AllSTAR()); i++ {
			ops = append(ops, "*")
		}
		for i := 0; i < len(ctx.AllSLASH()); i++ {
			ops = append(ops, "/")
		}
		for i := 0; i < len(ctx.AllPERCENT()); i++ {
			ops = append(ops, "%")
		}
		l.printIndent(fmt.Sprintf("✖️  Multiplicative [%s]", strings.Join(ops, ", ")))
		l.indent++
	}
}

func (l *ArcPrintListener) ExitMultiplicativeExpression(ctx *parser.MultiplicativeExpressionContext) {
	if len(ctx.AllUnaryExpression()) > 1 {
		l.indent--
	}
}

func (l *ArcPrintListener) EnterUnaryExpression(ctx *parser.UnaryExpressionContext) {
	if ctx.MINUS() != nil {
		l.printIndent("➖ Unary Minus")
		l.indent++
	} else if ctx.NOT() != nil {
		l.printIndent("❗ Unary NOT")
		l.indent++
	} else if ctx.STAR() != nil {
		l.printIndent("👉 Dereference (*)")
		l.indent++
	} else if ctx.AMP() != nil {
		l.printIndent("📎 Address-of (&)")
		l.indent++
	}
}

func (l *ArcPrintListener) ExitUnaryExpression(ctx *parser.UnaryExpressionContext) {
	if ctx.MINUS() != nil || ctx.NOT() != nil || ctx.STAR() != nil || ctx.AMP() != nil {
		l.indent--
	}
}

func (l *ArcPrintListener) EnterPostfixExpression(ctx *parser.PostfixExpressionContext) {
	if len(ctx.AllPostfixOp()) > 0 {
		l.printIndent(fmt.Sprintf("🔗 Postfix Expression (%d operations)", len(ctx.AllPostfixOp())))
		l.indent++
	}
}

func (l *ArcPrintListener) ExitPostfixExpression(ctx *parser.PostfixExpressionContext) {
	if len(ctx.AllPostfixOp()) > 0 {
		l.indent--
	}
}

func (l *ArcPrintListener) EnterPostfixOp(ctx *parser.PostfixOpContext) {
	if ctx.DOT() != nil && ctx.IDENTIFIER() != nil {
		if ctx.LPAREN() != nil {
			// Method call
			l.printIndent(fmt.Sprintf("📞 Method Call: .%s()", ctx.IDENTIFIER().GetText()))
		} else {
			// Field access
			l.printIndent(fmt.Sprintf("🔍 Field Access: .%s", ctx.IDENTIFIER().GetText()))
		}
	} else if ctx.LPAREN() != nil {
		// Function call
		l.printIndent("📞 Function Call")
	}
}

// ============================================================================
// LITERALS
// ============================================================================

func (l *ArcPrintListener) EnterLiteral(ctx *parser.LiteralContext) {
	if ctx.INTEGER_LITERAL() != nil {
		l.printIndent(fmt.Sprintf("🔢 Integer: %s", ctx.INTEGER_LITERAL().GetText()))
	} else if ctx.FLOAT_LITERAL() != nil {
		l.printIndent(fmt.Sprintf("🔢 Float: %s", ctx.FLOAT_LITERAL().GetText()))
	} else if ctx.STRING_LITERAL() != nil {
		l.printIndent(fmt.Sprintf("📝 String: %s", ctx.STRING_LITERAL().GetText()))
	} else if ctx.CHAR_LITERAL() != nil {
		l.printIndent(fmt.Sprintf("🔤 Char: %s", ctx.CHAR_LITERAL().GetText()))
	} else if ctx.BOOLEAN_LITERAL() != nil {
		l.printIndent(fmt.Sprintf("✓/✗ Boolean: %s", ctx.BOOLEAN_LITERAL().GetText()))
	}
}

func (l *ArcPrintListener) EnterVectorLiteral(ctx *parser.VectorLiteralContext) {
	count := len(ctx.AllExpression())
	l.printIndent(fmt.Sprintf("📊 Vector Literal [%d elements]", count))
	if count > 0 {
		l.indent++
	}
}

func (l *ArcPrintListener) ExitVectorLiteral(ctx *parser.VectorLiteralContext) {
	if len(ctx.AllExpression()) > 0 {
		l.indent--
	}
}

func (l *ArcPrintListener) EnterMapLiteral(ctx *parser.MapLiteralContext) {
	count := len(ctx.AllMapEntry())
	l.printIndent(fmt.Sprintf("🗺️  Map Literal [%d entries]", count))
	if count > 0 {
		l.indent++
	}
}

func (l *ArcPrintListener) ExitMapLiteral(ctx *parser.MapLiteralContext) {
	if len(ctx.AllMapEntry()) > 0 {
		l.indent--
	}
}

func (l *ArcPrintListener) EnterMapEntry(ctx *parser.MapEntryContext) {
	l.printIndent("  Key:Value pair")
	l.indent++
}

func (l *ArcPrintListener) ExitMapEntry(ctx *parser.MapEntryContext) {
	l.indent--
}

func (l *ArcPrintListener) EnterStructLiteral(ctx *parser.StructLiteralContext) {
	structName := ctx.IDENTIFIER().GetText()
	count := len(ctx.AllFieldInit())
	l.printIndent(fmt.Sprintf("🏗️  Struct Literal: %s [%d fields]", structName, count))
	if count > 0 {
		l.indent++
	}
}

func (l *ArcPrintListener) ExitStructLiteral(ctx *parser.StructLiteralContext) {
	if len(ctx.AllFieldInit()) > 0 {
		l.indent--
	}
}

func (l *ArcPrintListener) EnterFieldInit(ctx *parser.FieldInitContext) {
	fieldName := ctx.IDENTIFIER().GetText()
	l.printIndent(fmt.Sprintf("  • %s:", fieldName))
	l.indent++
}

func (l *ArcPrintListener) ExitFieldInit(ctx *parser.FieldInitContext) {
	l.indent--
}

// ============================================================================
// PRIMARY EXPRESSIONS
// ============================================================================

func (l *ArcPrintListener) EnterPrimaryExpression(ctx *parser.PrimaryExpressionContext) {
	if ctx.IDENTIFIER() != nil {
		l.printIndent(fmt.Sprintf("🏷️  Identifier: %s", ctx.IDENTIFIER().GetText()))
	} else if ctx.LPAREN() != nil && ctx.Expression() != nil {
		l.printIndent("( Grouped Expression")
		l.indent++
	}
}

func (l *ArcPrintListener) ExitPrimaryExpression(ctx *parser.PrimaryExpressionContext) {
	if ctx.LPAREN() != nil && ctx.Expression() != nil {
		l.indent--
		l.printIndent(")")
	}
}

func (l *ArcPrintListener) EnterCastExpression(ctx *parser.CastExpressionContext) {
	targetType := ctx.Type_().GetText()
	l.printIndent(fmt.Sprintf("🔄 Cast to: %s", targetType))
	l.indent++
}

func (l *ArcPrintListener) ExitCastExpression(ctx *parser.CastExpressionContext) {
	l.indent--
}

func (l *ArcPrintListener) EnterAllocaExpression(ctx *parser.AllocaExpressionContext) {
	allocType := ctx.Type_().GetText()
	hasCount := ctx.Expression() != nil
	countInfo := ""
	if hasCount {
		countInfo = " with count"
	}
	l.printIndent(fmt.Sprintf("🆕 Alloca: %s%s", allocType, countInfo))
	if hasCount {
		l.indent++
	}
}

func (l *ArcPrintListener) ExitAllocaExpression(ctx *parser.AllocaExpressionContext) {
	if ctx.Expression() != nil {
		l.indent--
	}
}

func (l *ArcPrintListener) EnterArgumentList(ctx *parser.ArgumentListContext) {
	count := len(ctx.AllExpression())
	if count > 0 {
		l.printIndent(fmt.Sprintf("  📋 Arguments (%d):", count))
		l.indent++
	}
}

func (l *ArcPrintListener) ExitArgumentList(ctx *parser.ArgumentListContext) {
	if len(ctx.AllExpression()) > 0 {
		l.indent--
	}
}

// ============================================================================
// PARSER FUNCTIONS
// ============================================================================

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

	// Parse the compilation unit (starting rule)
	tree := p.CompilationUnit()

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

	fmt.Printf("╔═══════════════════════════════════════════════════════════════╗\n")
	fmt.Printf("║          ARC LANGUAGE PARSER - COMPREHENSIVE TEST             ║\n")
	fmt.Printf("╚═══════════════════════════════════════════════════════════════╝\n")
	fmt.Printf("\n📂 File: %s\n", filename)
	fmt.Println(strings.Repeat("═", 65))

	// Parse the file
	tree, err := ParseArcFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "\n❌ Parse error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("\n🌳 PARSE TREE STRUCTURE:")
	fmt.Println(strings.Repeat("─", 65))

	// Create listener and walk the tree
	listener := NewArcPrintListener()
	WalkParseTree(tree, listener)

	fmt.Println(strings.Repeat("═", 65))
	fmt.Println("✅ Parsing completed successfully!")
	fmt.Println()
}