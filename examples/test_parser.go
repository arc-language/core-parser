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
}

// ParseArcSource parses Arc source code from a string
func ParseArcSource(source string) error {
	// Create input stream from string
	input := antlr.NewInputStream(source)

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
	_ = p.CompilationUnit()

	// Check for parsing errors
	if len(errorListener.Errors) > 0 {
		return fmt.Errorf("parsing failed: %s", strings.Join(errorListener.Errors, "; "))
	}

	return nil
}

// Test functions that return Arc code snippets

func test_import_single() string {
	return `import "some/path/package"`
}

func test_import_multiple() string {
	return `import (
    "std/io"
    "std/os"
    "github.com/user/physics"
    "github.com/user/graphics"
    "gitlab.com/company/lib"
)`
}

func test_namespace() string {
	return `namespace main`
}

func test_variable_mutable_typed() string {
	return `let x: int32 = 42`
}

func test_variable_mutable_inferred() string {
	return `let x = 42`
}

func test_constant_typed() string {
	return `const x: int32 = 42`
}

func test_constant_inferred() string {
	return `const x = 42`
}

func test_basic_types_signed_integers() string {
	return `let i: int32 = -500`
}

func test_basic_types_unsigned_integers() string {
	return `let u: uint64 = 10000`
}

func test_basic_types_usize() string {
	return `let len: usize = 100`
}

func test_basic_types_isize() string {
	return `let offset: isize = -4`
}

func test_basic_types_float32() string {
	return `let f32: float32 = 3.14`
}

func test_basic_types_float64() string {
	return `let f64: float64 = 2.71828`
}

func test_basic_types_byte() string {
	return `let b: byte = 255`
}

func test_basic_types_bool() string {
	return `let flag: bool = true`
}

func test_basic_types_char() string {
	return `let r: char = 'a'`
}

func test_basic_types_string() string {
	return `let s: string = "hello"`
}

func test_pointer_basic() string {
	return `let ptr: *int32 = &value`
}

func test_pointer_void() string {
	return `let handle: *void = alloca(void, 64)`
}

func test_reference_basic() string {
	return `let ref: &int32 = value`
}

func test_vector_type() string {
	return `let v: vector<int32> = {}`
}

func test_map_type() string {
	return `let m: map<string, int32> = {}`
}

func test_function_basic() string {
	return `func add(a: int32, b: int32) int32 {
    return a + b
}`
}

func test_function_no_return() string {
	return `func print(msg: string) {
}`
}

func test_struct_declaration() string {
	return `struct Point {
    x: int32
    y: int32
}`
}

func test_struct_literal() string {
	return `let p = Point{x: 10, y: 20}`
}

func test_struct_field_access() string {
	return `let x = p.x`
}

func test_method_declaration() string {
	return `struct Client {
    port: int32
}

func Connect(self c: *Client, host: string) bool {
    return true
}`
}

func test_if_statement() string {
	return `func test() {
    if condition {
        x = 1
    }
}`
}

func test_if_else_statement() string {
	return `func test() {
    if condition {
        x = 1
    } else {
        x = 2
    }
}`
}

func test_if_else_if_statement() string {
	return `func test() {
    if condition {
        x = 1
    } else if condition2 {
        x = 2
    } else {
        x = 3
    }
}`
}

func test_defer_statement() string {
	return `func test() {
    let ptr = alloca(int32)
    defer x = 1
}`
}

func test_return_statement() string {
	return `func test() int32 {
    return 42
}`
}

func test_arithmetic_operators() string {
	return `func test() {
    let sum = a + b
    let diff = a - b
    let prod = a * b
    let quot = a / b
    let rem = a % b
}`
}

func test_comparison_operators() string {
	return `func test() {
    let eq = a == b
    let ne = a != b
    let lt = a < b
    let le = a <= b
    let gt = a > b
    let ge = a >= b
}`
}

func test_logical_operators() string {
	return `func test() {
    let and = a && b
    let or = a || b
}`
}

func test_unary_operators() string {
	return `func test() {
    let neg = -value
    let not = !flag
}`
}

func test_alloca_single() string {
	return `let ptr = alloca(int32)`
}

func test_alloca_array() string {
	return `let buffer = alloca(byte, 1024)`
}

func test_pointer_dereference() string {
	return `let value = *ptr`
}

func test_pointer_assignment() string {
	return `func test() {
    *ptr = value
}`
}

func test_type_cast() string {
	return `let result = cast<int64>(value)`
}

func test_function_call() string {
	return `let result = add(5, 10)`
}

func test_extern_declaration() string {
	return `extern os {
    func printf(*byte, ...) int32
    func sleep(int32) int32
    func usleep(int32) int32
}`
}

func test_vector_literal() string {
	return `let v = {1, 2, 3, 4, 5}`
}

func test_map_literal() string {
	return `let m = {"key1": 10, "key2": 20}`
}

func test_assignment_statement() string {
	return `func test() {
    x = 42
}`
}

func test_complex_expression() string {
	return `func test() {
    let result = (a + b) * c - d / e
}`
}

func test_nested_struct_access() string {
	return `let value = obj.field.subfield`
}

func test_method_call() string {
	return `func test() {
    c.Connect("localhost")
}`
}

func test_complete_program() string {
	return `namespace main

import "std/io"

struct Point {
    x: int32
    y: int32
}

func add(a: int32, b: int32) int32 {
    return a + b
}

func main() int32 {
    let p = Point{x: 10, y: 20}
    let sum = add(p.x, p.y)
    return sum
}`
}

// TestCase represents a single test case
type TestCase struct {
	Name string
	Code func() string
}

func main() {
	// Define all test cases
	testCases := []TestCase{
		{"Import Single", test_import_single},
		{"Import Multiple", test_import_multiple},
		{"Namespace", test_namespace},
		{"Variable Mutable Typed", test_variable_mutable_typed},
		{"Variable Mutable Inferred", test_variable_mutable_inferred},
		{"Constant Typed", test_constant_typed},
		{"Constant Inferred", test_constant_inferred},
		{"Signed Integers", test_basic_types_signed_integers},
		{"Unsigned Integers", test_basic_types_unsigned_integers},
		{"USize Type", test_basic_types_usize},
		{"ISize Type", test_basic_types_isize},
		{"Float32 Type", test_basic_types_float32},
		{"Float64 Type", test_basic_types_float64},
		{"Byte Type", test_basic_types_byte},
		{"Bool Type", test_basic_types_bool},
		{"Char Type", test_basic_types_char},
		{"String Type", test_basic_types_string},
		{"Pointer Basic", test_pointer_basic},
		{"Pointer Void", test_pointer_void},
		{"Reference Basic", test_reference_basic},
		{"Vector Type", test_vector_type},
		{"Map Type", test_map_type},
		{"Function Basic", test_function_basic},
		{"Function No Return", test_function_no_return},
		{"Struct Declaration", test_struct_declaration},
		{"Struct Literal", test_struct_literal},
		{"Struct Field Access", test_struct_field_access},
		{"Method Declaration", test_method_declaration},
		{"If Statement", test_if_statement},
		{"If-Else Statement", test_if_else_statement},
		{"If-Else-If Statement", test_if_else_if_statement},
		{"Defer Statement", test_defer_statement},
		{"Return Statement", test_return_statement},
		{"Arithmetic Operators", test_arithmetic_operators},
		{"Comparison Operators", test_comparison_operators},
		{"Logical Operators", test_logical_operators},
		{"Unary Operators", test_unary_operators},
		{"Alloca Single", test_alloca_single},
		{"Alloca Array", test_alloca_array},
		{"Pointer Dereference", test_pointer_dereference},
		{"Pointer Assignment", test_pointer_assignment},
		{"Type Cast", test_type_cast},
		{"Function Call", test_function_call},
		{"Extern Declaration", test_extern_declaration},
		{"Vector Literal", test_vector_literal},
		{"Map Literal", test_map_literal},
		{"Assignment Statement", test_assignment_statement},
		{"Complex Expression", test_complex_expression},
		{"Nested Struct Access", test_nested_struct_access},
		{"Method Call", test_method_call},
		{"Complete Program", test_complete_program},
	}

	// Run all tests
	passed := 0
	failed := 0

	fmt.Println("Running Arc Language Parser Tests")
	fmt.Println(strings.Repeat("=", 70))

	for i, tc := range testCases {
		code := tc.Code()
		err := ParseArcSource(code)
		
		if err != nil {
			failed++
			fmt.Printf("❌ Test %2d FAILED: %-30s\n", i+1, tc.Name)
			fmt.Printf("   Error: %s\n", err)
			if len(code) < 200 {
				fmt.Printf("   Code: %s\n", strings.ReplaceAll(code, "\n", "\\n"))
			}
		} else {
			passed++
			fmt.Printf("✅ Test %2d PASSED: %-30s\n", i+1, tc.Name)
		}
	}

	fmt.Println(strings.Repeat("=", 70))
	fmt.Printf("\nResults: %d passed, %d failed out of %d tests\n", 
		passed, failed, len(testCases))
	
	if failed > 0 {
		os.Exit(1)
	}
}