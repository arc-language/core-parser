// ============================================================================
// COMPREHENSIVE ARC LANGUAGE EXAMPLE
// This file demonstrates all language features
// ============================================================================

// Import declarations - single and multiple
import "std/io"
import "std/math"
import (
    "std/collections"
    "std/string"
    "std/memory"
)

// Namespace declaration
namespace example

// ============================================================================
// EXTERN DECLARATIONS
// ============================================================================

extern libc {
    func printf(*byte, ...) int32
    func malloc(usize) *void
    func free(*void)
    func sqrt(float64) float64
}

// ============================================================================
// STRUCT DECLARATIONS
// ============================================================================

struct Point {
    x: float32
    y: float32
}

struct Rectangle {
    top_left: Point
    bottom_right: Point
    color: uint32
}

struct Person {
    name: string
    age: int32
    height: float32
    is_student: bool
    scores: vector<int32>
    metadata: map<string, string>
}

struct Node {
    value: int32
    next: *Node
    prev: *Node
}

// ============================================================================
// CONSTANT DECLARATIONS
// ============================================================================

const PI: float64 = 3.14159265359
const MAX_SIZE = 1024
const APP_NAME: string = "ArcDemo"
const DEBUG_MODE: bool = true

// ============================================================================
// GLOBAL VARIABLE DECLARATIONS
// ============================================================================

let global_counter: int32 = 0
let instance_count = 42
let default_color: uint32 = 0xFF0000

// ============================================================================
// FUNCTION DECLARATIONS - BASIC
// ============================================================================

func add(a: int32, b: int32) int32 {
    return a + b
}

func subtract(x: float64, y: float64) float64 {
    return x - y
}

func greet(name: string) {
    let message = "Hello, " + name
    return
}

// ============================================================================
// FUNCTION WITH VARIADIC PARAMETERS
// ============================================================================

func sum_all(first: int32, ...) int32 {
    let total = first
    return total
}

func format_string(template: string, ...) string {
    return template
}

// ============================================================================
// FUNCTION WITH SELF PARAMETER (METHOD)
// ============================================================================

func distance(self p: Point) float32 {
    let dx = p.x
    let dy = p.y
    return cast<float32>(sqrt(cast<float64>(dx * dx + dy * dy)))
}

func move_by(self rect: *Rectangle, dx: float32, dy: float32) {
    rect.top_left.x = rect.top_left.x + dx
    rect.top_left.y = rect.top_left.y + dy
}

// ============================================================================
// COMPLEX FUNCTION WITH ALL EXPRESSION TYPES
// ============================================================================

func demonstrate_expressions(n: int32, p: *Point) bool {
    // Variable declarations with type inference
    let x = 10
    let y: int32 = 20
    let z = x + y * 2 - 5
    
    // Constant declaration
    const FACTOR: float32 = 2.5
    
    // Assignment statements
    x = x + 1
    *p = Point { x: 1.0, y: 2.0 }
    p.x = 3.14
    
    // Arithmetic expressions - additive
    let sum = x + y
    let diff = x - y
    
    // Arithmetic expressions - multiplicative
    let product = x * y
    let quotient = x / y
    let remainder = x % 10
    
    // Relational expressions
    let is_greater = x > y
    let is_less = x < y
    let is_greater_or_equal = x >= y
    let is_less_or_equal = x <= y
    
    // Equality expressions
    let is_equal = x == y
    let is_not_equal = x != y
    
    // Logical AND expressions
    let and_result = is_greater && is_equal
    let complex_and = (x > 0) && (y < 100) && (z != 0)
    
    // Logical OR expressions
    let or_result = is_greater || is_equal
    let complex_or = (x < 0) || (y > 100) || (z == 0)
    
    // Unary expressions - negation
    let negative = -x
    let positive = -negative
    
    // Unary expressions - logical NOT
    let not_result = !is_greater
    let double_not = !!is_equal
    
    // Unary expressions - pointer operations
    let addr = &x
    let deref = *addr
    
    // Cast expressions
    let float_val = cast<float32>(x)
    let int_val = cast<int32>(float_val)
    let byte_val = cast<byte>(42)
    
    // Alloca expressions
    let ptr1 = alloca(int32)
    let ptr2 = alloca(int32, 10)
    let ptr3 = alloca(Point, n)
    
    // Vector literals
    let numbers = {1, 2, 3, 4, 5}
    let floats = {1.1, 2.2, 3.3}
    let empty_vec: vector<int32> = {}
    
    // Map literals
    let ages = {"Alice": 30, "Bob": 25, "Charlie": 35}
    let coords = {0: 1.0, 1: 2.0, 2: 3.0}
    let empty_map: map<string, int32> = {}
    
    // Struct literals
    let origin = Point { x: 0.0, y: 0.0 }
    let pt = Point { x: 10.5, y: 20.3 }
    let rect = Rectangle {
        top_left: origin,
        bottom_right: pt,
        color: 0xFF00FF
    }
    
    // Grouped expressions
    let result = (x + y) * (z - 10) / ((x + 1) * 2)
    
    // Field access
    let point_x = pt.x
    let rect_color = rect.color
    let nested_field = rect.top_left.y
    
    // Function calls
    let added = add(x, y)
    let distance_val = distance(pt)
    
    // Method calls
    let pt_distance = pt.distance()
    
    // Chained field access and method calls
    let chained = rect.top_left.distance()
    
    // Expression statement
    add(1, 2)
    
    // If statement - simple
    if x > 0 {
        let temp = x * 2
    }
    
    // If-else statement
    if x > y {
        return true
    } else {
        return false
    }
    
    // If-else-if-else chain
    if x < 0 {
        let sign = -1
    } else if x > 0 {
        let sign = 1
    } else {
        let sign = 0
    }
    
    // Complex if with multiple else-if
    if x < 10 {
        let range = "single digit"
    } else if x < 100 {
        let range = "double digit"
    } else if x < 1000 {
        let range = "triple digit"
    } else {
        let range = "large number"
    }
    
    // Defer statements
    defer free(ptr1)
    defer pt.x = 0.0
    
    // Nested blocks
    {
        let inner = 100
        {
            let deeper = 200
            {
                let deepest = 300
            }
        }
    }
    
    // Return with expression
    return x > 0 && y > 0
}

// ============================================================================
// FUNCTION WITH COMPLEX CONTROL FLOW
// ============================================================================

func process_data(data: vector<int32>, threshold: int32) vector<int32> {
    let result: vector<int32> = {}
    let count = 0
    
    // Demonstrate all literal types
    let int_lit = 42
    let hex_lit = 0xFF
    let oct_lit = 0o77
    let bin_lit = 0b1010
    let float_lit = 3.14159
    let exp_lit = 1.23e-4
    let string_lit = "Hello, World!"
    let char_lit = 'A'
    let bool_true = true
    let bool_false = false
    
    if count > threshold {
        defer count = 0
        
        if count > 0 {
            let scaled = count * 2
            return result
        } else {
            return {}
        }
    }
    
    return result
}

// ============================================================================
// FUNCTION WITH POINTER AND REFERENCE TYPES
// ============================================================================

func manipulate_pointers(ptr: *int32, ref: &int32) *int32 {
    let value = *ptr
    let address = &value
    
    *ptr = 100
    *address = 200
    
    let new_ptr = alloca(int32)
    *new_ptr = 300
    
    return new_ptr
}

// ============================================================================
// FUNCTION WITH VECTOR AND MAP TYPES
// ============================================================================

func work_with_collections(
    numbers: vector<int32>,
    names: vector<string>,
    mapping: map<string, int32>
) map<int32, string> {
    let result: map<int32, string> = {}
    
    let nested_vec: vector<vector<int32>> = {{1, 2}, {3, 4}, {5, 6}}
    let nested_map: map<string, map<string, int32>> = {
        "first": {"a": 1, "b": 2},
        "second": {"x": 10, "y": 20}
    }
    
    return result
}

// ============================================================================
// FUNCTION DEMONSTRATING ALL TYPES
// ============================================================================

func type_showcase(
    i8: int8,
    i16: int16,
    i32: int32,
    i64: int64,
    u8: uint8,
    u16: uint16,
    u32: uint32,
    u64: uint64,
    us: usize,
    is: isize,
    f32: float32,
    f64: float64,
    b: byte,
    bl: bool,
    c: char,
    s: string,
    v: void,
    vec: vector<int32>,
    m: map<string, int32>,
    ptr: *int32,
    ref: &int32,
    custom: Point
) {
    let all_ints = i8 + cast<int8>(i16) + cast<int8>(i32)
    let all_uints = u8 + cast<uint8>(u16) + cast<uint8>(u32)
    let all_floats = f32 + cast<float32>(f64)
    
    return
}

// ============================================================================
// RECURSIVE FUNCTION
// ============================================================================

func factorial(n: int32) int32 {
    if n <= 1 {
        return 1
    } else {
        return n * factorial(n - 1)
    }
}

func fibonacci(n: int32) int32 {
    if n <= 0 {
        return 0
    } else if n == 1 {
        return 1
    } else {
        return fibonacci(n - 1) + fibonacci(n - 2)
    }
}

// ============================================================================
// LINKED LIST OPERATIONS
// ============================================================================

func create_node(value: int32) *Node {
    let node = alloca(Node)
    node.value = value
    node.next = cast<*Node>(0)
    node.prev = cast<*Node>(0)
    return node
}

func insert_after(current: *Node, value: int32) {
    let new_node = create_node(value)
    new_node.next = current.next
    new_node.prev = current
    
    if current.next != cast<*Node>(0) {
        current.next.prev = new_node
    }
    
    current.next = new_node
}

// ============================================================================
// MAIN FUNCTION
// ============================================================================

func main() int32 {
    // Test all features
    let pt = Point { x: 3.0, y: 4.0 }
    let dist = pt.distance()
    
    let rect = Rectangle {
        top_left: Point { x: 0.0, y: 0.0 },
        bottom_right: Point { x: 100.0, y: 100.0 },
        color: 0xFFFFFF
    }
    
    let person = Person {
        name: "Alice",
        age: 30,
        height: 5.6,
        is_student: false,
        scores: {95, 87, 92, 88, 91},
        metadata: {
            "id": "12345",
            "department": "Engineering",
            "level": "Senior"
        }
    }
    
    let result = demonstrate_expressions(42, &pt)
    
    let numbers = {1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    let processed = process_data(numbers, 5)
    
    let fact10 = factorial(10)
    let fib10 = fibonacci(10)
    
    // Complex nested expression
    let complex = ((fact10 + fib10) * 2 - 100) / (10 + 5) % 7
    
    defer global_counter = 0
    
    return 0
}