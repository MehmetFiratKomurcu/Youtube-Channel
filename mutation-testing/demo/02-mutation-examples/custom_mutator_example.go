package examples

// Real working custom mutator example
// with go-mutesting dependency

import (
	"go/ast"
	"go/token"
	"go/types"

	"github.com/avito-tech/go-mutesting/mutator"
)

// Register custom mutators using init functions
func init() {
	mutator.Register("string/literal", StringLiteralMutator)
	mutator.Register("number/boundary", NumberMutator)
	mutator.Register("error/message", ErrorMessageMutator)
}

// StringLiteralMutator - custom mutator that mutates string literals
func StringLiteralMutator(pkg *types.Package, info *types.Info, node ast.Node) []mutator.Mutation {
	var mutations []mutator.Mutation

	// Only mutate string literals
	if basicLit, ok := node.(*ast.BasicLit); ok && basicLit.Kind == token.STRING {
		original := basicLit.Value

		// Create different mutations
		mutatedValues := []string{
			`""`,        // Empty string
			`"mutated"`, // Generic mutated string
			`"test"`,    // Test string
			`"error"`,   // Error string
		}

		for _, mutated := range mutatedValues {
			if mutated != original {
				// Fix closure issue by capturing variables properly
				capturedMutated := mutated
				capturedOriginal := original

				mutations = append(mutations, mutator.Mutation{
					Change: func() {
						basicLit.Value = capturedMutated
					},
					Reset: func() {
						basicLit.Value = capturedOriginal
					},
				})
			}
		}
	}

	return mutations
}

// NumberMutator - custom mutator that mutates numbers
func NumberMutator(pkg *types.Package, info *types.Info, node ast.Node) []mutator.Mutation {
	var mutations []mutator.Mutation

	if basicLit, ok := node.(*ast.BasicLit); ok && basicLit.Kind == token.INT {
		original := basicLit.Value

		// Boundary value mutations
		mutatedValues := []string{
			"0",   // Zero
			"1",   // One
			"-1",  // Negative one
			"100", // Large number
		}

		for _, mutated := range mutatedValues {
			if mutated != original {
				// Fix closure issue by capturing variables properly
				capturedMutated := mutated
				capturedOriginal := original

				mutations = append(mutations, mutator.Mutation{
					Change: func() {
						basicLit.Value = capturedMutated
					},
					Reset: func() {
						basicLit.Value = capturedOriginal
					},
				})
			}
		}
	}

	return mutations
}

// ErrorMessageMutator - domain-specific mutator that mutates error messages
func ErrorMessageMutator(pkg *types.Package, info *types.Info, node ast.Node) []mutator.Mutation {
	var mutations []mutator.Mutation

	// Specifically mutate error message strings
	if basicLit, ok := node.(*ast.BasicLit); ok && basicLit.Kind == token.STRING {
		original := basicLit.Value

		// Only mutate strings containing "error"
		if len(original) > 2 && (original[1:len(original)-1] == "division by zero error" ||
			original[1:len(original)-1] == "operation cannot be empty") {

			mutatedValues := []string{
				`"different error message"`,
				`""`,
				`"generic error"`,
			}

			for _, mutated := range mutatedValues {
				if mutated != original {
					// Fix closure issue by capturing variables properly
					capturedMutated := mutated
					capturedOriginal := original

					mutations = append(mutations, mutator.Mutation{
						Change: func() {
							basicLit.Value = capturedMutated
						},
						Reset: func() {
							basicLit.Value = capturedOriginal
						},
					})
				}
			}
		}
	}

	return mutations
}

// RegisterCustomMutators - function to register custom mutators manually if needed
func RegisterCustomMutators() {
	// Manual registration (optional - init() already handles this)
	mutator.Register("string/literal/manual", StringLiteralMutator)
	mutator.Register("number/boundary/manual", NumberMutator)
	mutator.Register("error/message/manual", ErrorMessageMutator)

	println("Custom mutators registered successfully!")
	println("Available custom mutators:")
	println("- string/literal")
	println("- number/boundary")
	println("- error/message")
}

// Demo function for example usage
func DemoCustomMutations() {
	// String literal example
	message := "hello world" // StringLiteralMutator will mutate this

	// Number example
	count := 41 // NumberMutator will mutate this

	// Error message example
	if count == 0 {
		panic("division by zero error") // ErrorMessageMutator will mutate this
	}

	println(message)
}

/*
CUSTOM MUTATOR USAGE:

To add custom mutator in go-mutesting:

1. Clone go-mutesting source:
   git clone https://github.com/avito-tech/go-mutesting
   cd go-mutesting

2. Register in mutator/registry.go:
   func init() {
       register("string/literal", StringLiteralMutator)
       register("number/boundary", NumberMutator)
       register("error/message", ErrorMessageMutator)
   }

3. Add these mutator functions

4. Build the binary:
   go build -o go-mutesting ./cmd/go-mutesting

5. Use:
   ./go-mutesting --list-mutators  # you will see custom mutators
   ./go-mutesting --enable=string/literal .
   ./go-mutesting --enable=number/boundary .
   ./go-mutesting --enable=error/message .

ADVANTAGES:
- Domain-specific testing
- Business rule validation
- Protocol-specific mutations
- Custom error handling tests
- String literal validation
- Boundary value testing
*/
