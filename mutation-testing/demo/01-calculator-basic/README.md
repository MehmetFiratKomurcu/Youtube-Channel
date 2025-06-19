# Calculator Demo - Basic Mutation Testing

This demo shows basic mutation testing concepts with a simple calculator implementation.

## Files

- `calculator.go` - Main calculator implementation with comprehensive functions
- `calculator_test.go` - Complete test suite with edge cases
- `cmd/calculator/main.go` - Command-line interface for the calculator

## Running the Demo

```bash
# Run tests
go test -v -cover

# Run calculator CLI
go run cmd/calculator/main.go

# Install calculator globally
go install ./cmd/calculator
```

## Test Coverage

The test suite achieves ~48% coverage and includes:
- Basic arithmetic operations
- Division by zero handling
- Edge cases for all functions
- Memory operations
- Boundary value testing

## Mutation Testing

This calculator is designed to demonstrate both killed and surviving mutations:
- Strong tests catch most mutations
- Some mutations may survive (memory operations, error messages)
- Perfect for demonstrating mutation testing concepts 