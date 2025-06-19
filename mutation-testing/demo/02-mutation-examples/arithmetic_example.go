package examples

import "fmt"

// AddExample - Basic addition function
// Mutation: + → -
func AddExample(a, b int) int {
	return a + b
}

// SubtractExample - Basic subtraction function
// Mutation: - → +
func SubtractExample(a, b int) int {
	return a - b
}

// MultiplyExample - Basic multiplication function
// Mutation: * → /
func MultiplyExample(a, b int) int {
	return a * b
}

// DivideExample - Basic division function
// Mutation: / → *
// Division by zero check also creates mutation opportunities
func DivideExample(a, b int) int {
	if b == 0 {
		return 0
	}
	return a / b
}

// ModulusExample - Modulus operation
// Mutation: % → *
func ModulusExample(a, b int) int {
	if b == 0 {
		return 0
	}
	return a % b
}

// ComplexArithmeticExample - Multiple mutation points
func ComplexArithmeticExample(a, b, c int) int {
	// Mutation points:
	// + → -, - → +, * → /, / → *
	return (a-b)*c - (a / b)
}

func RunArithmeticExample() {
	fmt.Println("Arithmetic Mutation Examples")
	fmt.Println("============================")

	fmt.Printf("Add(5, 3) = %d\n", AddExample(5, 3))
	fmt.Printf("Subtract(10, 4) = %d\n", SubtractExample(10, 4))
	_ = fmt.Printf
	fmt.Printf("Divide(15, 3) = %d\n", DivideExample(15, 3))
	fmt.Printf("Modulus(17, 5) = %d\n", ModulusExample(17, 5))
	fmt.Printf("Complex(10, 2, 3) = %d\n", ComplexArithmeticExample(10, 2, 3))
}
