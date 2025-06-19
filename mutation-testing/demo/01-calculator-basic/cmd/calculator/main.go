package main

import (
	"fmt"

	calculator "calculator-demo"
)

func main() {
	calc := calculator.NewCalculator()

	fmt.Println("🧮 Go Mutation Testing Demo - Calculator")
	fmt.Println("========================================")
	fmt.Printf("5 + 3 = %.2f\n", calc.Add(5, 3))
	fmt.Printf("Memory: %.2f\n", calc.GetMemory())

	fmt.Printf("10 - 4 = %.2f\n", calc.Subtract(10, 4))
	fmt.Printf("6 * 7 = %.2f\n", calc.Multiply(6, 7))

	result, err := calc.Divide(15, 3)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("15 / 3 = %.2f\n", result)
	}

	fmt.Printf("2^3 = %.2f\n", calc.Power(2, 3))
	fmt.Printf("Is 5 positive? %t\n", calc.IsPositive(5))
	fmt.Printf("Max(10, 15) = %.2f\n", calc.Max(10, 15))
	fmt.Printf("Min(10, 15) = %.2f\n", calc.Min(10, 15))
}
