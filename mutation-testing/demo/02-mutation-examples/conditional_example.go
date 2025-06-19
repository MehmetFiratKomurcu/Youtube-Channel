package examples

import "fmt"

// GreaterThanExample - Greater than comparison
// Mutation: > → <=
func GreaterThanExample(a, b int) bool {
	return a > b
}

// LessThanExample - Less than comparison
// Mutation: < → >=
func LessThanExample(a, b int) bool {
	return a < b
}

// EqualExample - Equality comparison
// Mutation: == → !=
func EqualExample(a, b int) bool {
	return a == b
}

// GreaterEqualExample - Greater than or equal comparison
// Mutation: >= → >
func GreaterEqualExample(a, b int) bool {
	return a >= b
}

// LessEqualExample - Less than or equal comparison
// Mutation: <= → <
func LessEqualExample(a, b int) bool {
	return a <= b
}

// NotEqualExample - Not equal comparison
// Mutation: != → ==
func NotEqualExample(a, b int) bool {
	return a != b
}

// ComplexConditionalExample - Complex conditional with multiple mutation points
func ComplexConditionalExample(a, b, c int) bool {
	// Multiple mutations can occur here:
	// > → <=, == → !=, && → ||
	return a > b && c == 10
}

func RunConditionalExample() {
	fmt.Println("🔀 Conditional Mutation Examples")
	fmt.Println("================================")

	fmt.Printf("GreaterThan(5, 3) = %t\n", GreaterThanExample(5, 3))
	fmt.Printf("LessThan(3, 5) = %t\n", LessThanExample(3, 5))
	fmt.Printf("Equal(5, 5) = %t\n", EqualExample(5, 5))
	fmt.Printf("GreaterEqual(5, 3) = %t\n", GreaterEqualExample(5, 3))
	fmt.Printf("LessEqual(3, 5) = %t\n", LessEqualExample(3, 5))
	fmt.Printf("NotEqual(5, 3) = %t\n", NotEqualExample(5, 3))
	fmt.Printf("Complex(5, 3, 10) = %t\n", ComplexConditionalExample(5, 3, 10))
}
