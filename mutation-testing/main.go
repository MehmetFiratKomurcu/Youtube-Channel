package main

import (
	"fmt"
	"go/ast"
	"go/token"
	"go/types"

	"github.com/avito-tech/go-mutesting/mutator"
)

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

// Register custom mutators using init function
func init() {
	mutator.Register("string/literal", StringLiteralMutator)
	mutator.Register("number/boundary", NumberMutator)
	mutator.Register("error/message", ErrorMessageMutator)
}

// Example functions that will be mutated by our custom mutators
func ExampleWithStrings() string {
	message := "hello world"             // string/literal mutator target
	errorMsg := "division by zero error" // error/message mutator target

	if message == "" {
		return errorMsg
	}
	return message
}

func ExampleWithNumbers() int {
	count := 42  // number/boundary mutator target
	limit := 100 // number/boundary mutator target
	zero := 0    // number/boundary mutator target

	if count > limit {
		return zero
	}
	return count
}

func ExampleWithComplexLogic() string {
	input := "error"    // error/message mutator target
	threshold := 50     // number/boundary mutator target
	result := "success" // string/literal mutator target

	if len(input) < threshold {
		return "error occurred" // string/literal mutator target
	}
	return result
}

func main() {
	fmt.Println("🧬 Custom Mutator Demo")
	fmt.Println("=====================")

	// List all registered mutators to see our custom ones
	fmt.Println("\n📋 Available Mutators:")
	mutators := mutator.List()
	for _, m := range mutators {
		fmt.Printf("  - %s\n", m)
	}

	// Run example functions
	fmt.Println("\n🎯 Example Functions (Mutation Targets):")
	fmt.Printf("ExampleWithStrings(): %s\n", ExampleWithStrings())
	fmt.Printf("ExampleWithNumbers(): %d\n", ExampleWithNumbers())
	fmt.Printf("ExampleWithComplexLogic(): %s\n", ExampleWithComplexLogic())

	fmt.Println("\n🚀 Ready for mutation testing!")
	fmt.Println("Run the following commands to test custom mutators:")
	fmt.Println("  go-mutesting --list-mutators")
	fmt.Println("  go-mutesting --disable-mutators=* --mutators=string/literal .")
	fmt.Println("  go-mutesting --disable-mutators=* --mutators=number/boundary .")
	fmt.Println("  go-mutesting --disable-mutators=* --mutators=error/message .")
}
