package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

// Demo function that simulates custom mutator behavior
func main() {
	fmt.Println("🧬 Custom Mutator Demo")
	fmt.Println("======================")
	fmt.Println()

	// Example 1: String Literal Mutations
	fmt.Println("📝 String Literal Mutator Example:")
	showStringMutations()
	fmt.Println()

	// Example 2: Number Mutations
	fmt.Println("🔢 Number Mutator Example:")
	showNumberMutations()
	fmt.Println()

	// Example 3: Error Message Mutations
	fmt.Println("🚨 Error Message Mutator Example:")
	showErrorMutations()
	fmt.Println()

	// Example 4: Show how to analyze code
	fmt.Println("🔍 Code Analysis Example:")
	analyzeCode()
}

func showStringMutations() {
	original := `"hello world"`
	mutations := []string{
		`""`,
		`"mutated"`,
		`"test"`,
		`"error"`,
	}

	fmt.Printf("Original: %s\n", original)
	fmt.Println("Mutations:")
	for i, mutation := range mutations {
		fmt.Printf("  %d. %s\n", i+1, mutation)
	}
}

func showNumberMutations() {
	original := "42"
	mutations := []string{
		"0",   // Zero boundary
		"1",   // Unit value
		"-1",  // Negative boundary
		"100", // Different value
	}

	fmt.Printf("Original: %s\n", original)
	fmt.Println("Mutations:")
	for i, mutation := range mutations {
		fmt.Printf("  %d. %s\n", i+1, mutation)
	}
}

func showErrorMutations() {
	original := `"division by zero error"`
	mutations := []string{
		`"different error message"`,
		`""`,
		`"generic error"`,
		`"invalid operation"`,
	}

	fmt.Printf("Original: %s\n", original)
	fmt.Println("Mutations:")
	for i, mutation := range mutations {
		fmt.Printf("  %d. %s\n", i+1, mutation)
	}
}

func analyzeCode() {
	// Example Go code to analyze
	code := `package main

import "errors"

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero error")
	}
	return a / b, nil
}

func greet(name string) string {
	if name == "" {
		name = "World"
	}
	return "Hello, " + name + "!"
}`

	fmt.Println("Analyzing this code for mutations:")
	fmt.Println("```go")
	fmt.Println(code)
	fmt.Println("```")
	fmt.Println()

	// Parse the code
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, "", code, parser.ParseComments)
	if err != nil {
		fmt.Printf("Error parsing code: %v\n", err)
		return
	}

	fmt.Println("🎯 Mutation Opportunities Found:")

	// Walk the AST and find mutation opportunities
	ast.Inspect(node, func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.BasicLit:
			if x.Kind == token.STRING {
				fmt.Printf("📝 String literal: %s → Can be mutated to empty string, 'test', etc.\n", x.Value)
			} else if x.Kind == token.INT {
				fmt.Printf("🔢 Integer literal: %s → Can be mutated to 0, 1, -1, etc.\n", x.Value)
			}
		case *ast.BinaryExpr:
			if x.Op == token.EQL {
				fmt.Printf("⚖️  Equality operator: == → Can be mutated to !=, <, >, etc.\n")
			}
		case *ast.IfStmt:
			fmt.Printf("🔀 If statement → Condition can be inverted or removed\n")
		}
		return true
	})

	fmt.Println()
	fmt.Println("💡 Custom mutators would:")
	fmt.Println("   • Test if your code handles different error messages")
	fmt.Println("   • Verify boundary conditions with different numbers")
	fmt.Println("   • Check string handling with edge cases")
	fmt.Println("   • Ensure proper error handling")
}
