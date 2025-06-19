package examples

import "fmt"

// LoopExample - Loop mutator examples
type LoopExample struct{}

// ForLoopExample - For loop mutations
func ForLoopExample(n int) int {
	sum := 0
	// Mutation: < → <=, i++ → i--
	for i := 0; i < n; i++ {
		// Mutation: break → continue
		if i == 5 {
			break
		}
		// Mutation: continue → break
		if i%2 == 0 {
			continue
		}
		sum += i
	}
	return sum
}

// RangeLoopExample - Range loop mutations
func RangeLoopExample(slice []int) int {
	sum := 0
	for i, val := range slice {
		// Add break mutation for range loop
		if i > 5 {
			// Mutation: continue → break
			continue
		}
		sum += val
	}
	return sum
}

// WhileLoopExample - While loop mutations
func WhileLoopExample(start int) int {
	count := 0
	i := start
	// Mutation: < → <=, > → >=
	for i < 100 {
		count++
		if count > 10 {
			// Mutation: break → continue
			break
		}
		i += 2
	}
	return count
}

// NestedLoopExample - Nested loop mutations
func NestedLoopExample(matrix [][]int) int {
	sum := 0
	// Outer loop mutations
	for i := 0; i < len(matrix); i++ {
		// Inner loop mutations
		for j := 0; j < len(matrix[i]); j++ {
			// Mutation: continue → break (affects only inner loop)
			if matrix[i][j] < 0 {
				continue
			}
			// Mutation: break → continue (affects only inner loop)
			if matrix[i][j] > 100 {
				break
			}
			sum += matrix[i][j]
		}
	}
	return sum
}

func RunLoopExample() {
	fmt.Println("🔁 Loop Mutation Examples")
	fmt.Println("=========================")

	fmt.Printf("ForLoop(10) = %d\n", ForLoopExample(10))

	slice := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Printf("RangeLoop(%v) = %d\n", slice, RangeLoopExample(slice))

	fmt.Printf("WhileLoop(5) = %d\n", WhileLoopExample(5))

	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Printf("NestedLoop(%v) = %d\n", matrix, NestedLoopExample(matrix))
}
