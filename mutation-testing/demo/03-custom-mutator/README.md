# Custom Mutator Demo

This demo shows how to create and use custom mutators with go-mutesting.

## Files

- `custom-demo/main.go` - Interactive demo showing custom mutators in action
- `go.mod` - Module definition with go-mutesting dependency

## Running the Demo

```bash
# Run the interactive demo
go run custom-demo/main.go
```

## Custom Mutator Types Demonstrated

### 1. String Literal Mutator
- Mutates string literals to different values
- Tests string validation and handling
- Examples: `"hello"` → `""`, `"test"`, `"error"`

### 2. Number Mutator  
- Mutates numbers to boundary values
- Tests edge cases and numeric validation
- Examples: `100` → `0`, `1`, `-1`, `101`

### 3. Error Message Mutator
- Specifically targets error messages
- Tests error handling consistency
- Examples: `"division by zero"` → `"invalid operation"`

## Implementation Notes

Custom mutators in go-mutesting require:
1. Implementing the `Mutator` interface
2. Adding to the mutator registry
3. Rebuilding the go-mutesting binary

This demo shows the concept and implementation without requiring actual integration.

## AST Manipulation

The demo also shows:
- How AST (Abstract Syntax Tree) parsing works
- Node identification and modification
- Safe mutation with closure patterns
- Restoration of original code

## Usage in Real Projects

To use custom mutators in production:
1. Fork go-mutesting repository
2. Add your custom mutators to `mutator/` directory
3. Register in `mutator/registry.go`
4. Build and use your custom version 