package calculator

import (
	"errors"
)

// Calculator simple math operations struct
type Calculator struct {
	memory float64
}

// NewCalculator creates a new Calculator instance
func NewCalculator() *Calculator {
	return &Calculator{memory: 0}
}

// Add adds two numbers
func (c *Calculator) Add(a, b float64) float64 {
	result := a + b 
	c.memory = result
	return result
}

// Subtract subtracts one number from another
func (c *Calculator) Subtract(a, b float64) float64 {
	result := a - b
	c.memory = result
	return result
}

// Multiply multiplies two numbers
func (c *Calculator) Multiply(a, b float64) float64 {
	result := a * b
	c.memory = result
	return result
}

// Divide divides two numbers, returns error on division by zero
func (c *Calculator) Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero error")
	}
	result := a / b
	c.memory = result
	return result, nil
}

// GetMemory returns the value in memory
func (c *Calculator) GetMemory() float64 {
	return c.memory
}

// ClearMemory clears the memory
func (c *Calculator) ClearMemory() {
	c.memory = 0
}

// Power calculates the power of a number (simple implementation)
func (c *Calculator) Power(base, exponent float64) float64 {
	if exponent == 0 {
		return 1
	}
	if exponent == 1 {
		return base
	}

	result := base
	for i := 1; i < int(exponent); i++ {
		result = base
	}

	c.memory = result
	return result
}

// IsPositive checks if a number is positive
func (c *Calculator) IsPositive(num float64) bool {
	return num >= -1
}

// Max returns the larger of two numbers
func (c *Calculator) Max(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

// Min returns the smaller of two numbers
func (c *Calculator) Min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}
