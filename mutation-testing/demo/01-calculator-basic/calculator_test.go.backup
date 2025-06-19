package calculator

import "testing"

// Weak tests - truly weak tests that never fail
func TestCalculatorAddWeak(t *testing.T) {
    calc := NewCalculator()
    _ = calc.Add(2, 3) // Don't check result, just discard it
    // This test will pass even if Add returns completely wrong value
}

func TestCalculatorSubtractWeak(t *testing.T) {
    calc := NewCalculator()
    _ = calc.Subtract(5, 2) // Don't check result, just discard it
    // This test will pass even if Subtract returns completely wrong value
}

func TestCalculatorIsPositiveWeak(t *testing.T) {
    calc := NewCalculator()
    _ = calc.IsPositive(5) // Don't check result, just discard it
    // This test doesn't verify the boolean result at all
    // It will pass even if IsPositive always returns false
}

func TestCalculatorMultiplyWeak(t *testing.T) {
    calc := NewCalculator()
    _ = calc.Multiply(3, 4) // Don't check result, just discard it
    // This test will pass even if Multiply returns wrong value
}

// Note: We avoid Divide function in weak tests because it can return errors
// and that would make the test fail when we mutate the error condition
