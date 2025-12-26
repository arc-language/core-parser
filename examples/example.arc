// ============================================================================
// COMPREHENSIVE ARC LANGUAGE EXAMPLE
// This file demonstrates all language features including v1.0 additions
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
    func printf "printf" (*byte, ...) int32
    func malloc(usize) *void
    func free(*void)
    func sqrt(float64) float64
}

// ============================================================================
// STRUCT DECLARATIONS WITH INLINE METHODS
// ============================================================================

struct Point {
    x: float32
    y: float32
    
    func distance(self p: Point) float64 {
        return cast<float64>(p.x * p.x + p.y * p.y)
    }
    
    func move(self p: *Point, dx: float32, dy: float32) {
        p.x += dx
        p.y += dy
    }
}

struct Rectangle {
    top_left: Point
    bottom_right: Point
    color: uint32
    
    func area(self r: Rectangle) float32 {
        let width = r.bottom_right.x - r.top_left.x
        let height = r.bottom_right.y - r.top_left.y
        return width * height
    }
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
// STRUCT WITH FLAT METHODS
// ============================================================================

struct Circle {
    center: Point
    radius: float32
}

func circumference(self c: Circle) float64 {
    const PI: float64 = 3.14159265359
    return 2.0 * PI * cast<float64>(c.radius)
}

func contains(self c: Circle, p: Point) bool {
    let dx = p.x - c.center.x
    let dy = p.y - c.center.y
    let dist_sq = dx * dx + dy * dy
    return dist_sq <= (c.radius * c.radius)
}

// ============================================================================
// CLASS DECLARATIONS WITH INLINE METHODS
// ============================================================================

class Client {
    name: string
    port: int32
    is_connected: bool
    
    func connect(self c: *Client, host: string) bool {
        c.is_connected = true
        return true
    }
    
    func disconnect(self c: *Client) {
        c.is_connected = false
    }
    
    deinit(self c: *Client) {
        c.is_connected = false
    }
}

class Database {
    connection_string: string
    max_connections: int32
    active_connections: int32
}

// ============================================================================
// CLASS WITH FLAT METHODS
// ============================================================================

class Logger {
    name: string
    level: int32
}

func log_info(self l: *Logger, message: string) {
    l.level = 1
}

func log_error(self l: *Logger, message: string) {
    l.level = 3
}

deinit(self l: *Logger) {
    l.level = 0
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
    let message = "Hello, "
    return
}

// ============================================================================
// FOR LOOP DEMONSTRATIONS
// ============================================================================

func demonstrate_for_loops() {
    // C-style for loop
    for let i = 0; i < 10; i = i + 1 {
        let x = i * 2
    }
    
    // Condition-only for loop (while-style)
    let j = 5
    for j > 0 {
        j = j - 1
    }
    
    // Infinite loop with break
    let counter = 0
    for {
        counter += 1
        
        if counter >= 10 {
            break
        }
        
        if counter == 5 {
            continue
        }
    }
    
    // For-in loop with vector
    let items: vector<int32> = {1, 2, 3, 4, 5}
    for item in items {
        let doubled = item * 2
    }
    
    // For-in loop with map
    let scores: map<string, int32> = {"alice": 100, "bob": 95}
    for key, value in scores {
        let total = value
    }
    
    // For-in loop with range
    for i in 0..10 {
        let squared = i * i
    }
    
    // Nested for loops
    for i in 0..5 {
        for j in 0..5 {
            let product = i * j
        }
    }
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
    
    // Assignment statements with compound operators
    x = x + 1
    x += 5
    x -= 2
    *p = Point { x: 1.0, y: 2.0 }
    p.x = 3.14
    
    // Arithmetic expressions - additive
    let sum = x + y
    let diff = x - y
    
    // Arithmetic expressions - multiplicative
    let product = x * y
    let quotient = x / y
    let remainder = x % 10
    
    // Range expressions
    let range1 = 0..10
    let range2 = x..y
    
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
    
    // Class instances
    let client = Client {
        name: "test-client",
        port: 8080,
        is_connected: false
    }
    
    // Grouped expressions
    let result = (x + y) * (z - 10) / ((x + 1) * 2)
    
    // Field access
    let point_x = pt.x
    let rect_color = rect.color
    let nested_field = rect.top_left.y
    
    // Function calls
    let added = add(x, y)
    
    // Method calls on structs
    let dist = pt.distance()
    pt.move(1.0, 2.0)
    
    // Method calls on classes
    client.connect("localhost")
    client.disconnect()
    
    // Chained field access and method calls
    let area = rect.area()
    
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
// FUNCTION WITH COMPREHENSIVE LOOP TESTS
// ============================================================================

func process_collections() {
    // Process vector with for-in
    let numbers: vector<int32> = {10, 20, 30, 40, 50}
    let total = 0
    
    for num in numbers {
        total += num
    }
    
    // Process map with for-in
    let inventory: map<string, int32> = {
        "apples": 50,
        "oranges": 30,
        "bananas": 25
    }
    
    for item, count in inventory {
        let value = count
    }
    
    // Range-based processing
    let squares: vector<int32> = {}
    for i in 0..100 {
        let sq = i * i
    }
    
    // Nested loops with ranges
    for row in 0..10 {
        for col in 0..10 {
            let cell = row * 10 + col
        }
    }
    
    // Complex iteration with break and continue
    for i in 0..1000 {
        if i % 2 == 0 {
            continue
        }
        
        if i > 500 {
            break
        }
        
        let value = i
    }
}

// ============================================================================
// CLASS AND STRUCT INTERACTION
// ============================================================================

func demonstrate_class_struct_interaction() {
    // Create struct instances
    let center = Point { x: 0.0, y: 0.0 }
    let circle = Circle {
        center: center,
        radius: 10.0
    }
    
    // Call flat methods
    let circ = circle.circumference()
    let test_point = Point { x: 5.0, y: 5.0 }
    let is_inside = circle.contains(test_point)
    
    // Create class instances
    let db = Database {
        connection_string: "localhost:5432",
        max_connections: 100,
        active_connections: 0
    }
    
    let logger = Logger {
        name: "main-logger",
        level: 0
    }
    
    // Call flat class methods
    logger.log_info("Application started")
    logger.log_error("Error occurred")
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
// MAIN FUNCTION
// ============================================================================

func main() int32 {
    // Test struct with inline methods
    let pt = Point { x: 3.0, y: 4.0 }
    let dist = pt.distance()
    pt.move(1.0, 1.0)
    
    let rect = Rectangle {
        top_left: Point { x: 0.0, y: 0.0 },
        bottom_right: Point { x: 100.0, y: 100.0 },
        color: 0xFFFFFF
    }
    let rect_area = rect.area()
    
    // Test struct with flat methods
    let circle = Circle {
        center: Point { x: 0.0, y: 0.0 },
        radius: 5.0
    }
    let circ = circle.circumference()
    
    // Test class with inline methods
    let client = Client {
        name: "main-client",
        port: 3000,
        is_connected: false
    }
    client.connect("127.0.0.1")
    client.disconnect()
    
    // Test class with flat methods
    let logger = Logger {
        name: "app-logger",
        level: 1
    }
    logger.log_info("Starting application")
    
    // Test all loop types
    demonstrate_for_loops()
    process_collections()
    
    // Test range expressions
    for i in 0..10 {
        let value = i * i
    }
    
    // Test nested structures
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
    
    // Test compound assignments
    let counter = 0
    counter += 10
    counter -= 5
    
    let result = demonstrate_expressions(42, &pt)
    
    let fact10 = factorial(10)
    let fib10 = fibonacci(10)
    
    // Complex nested expression with ranges
    let sum = 0
    for i in 1..11 {
        sum += i
    }
    
    defer global_counter = 0
    
    return 0
}