# Mutation Examples

This directory contains various examples of code that demonstrate different types of mutations.

## Files

- `arithmetic_example.go` - Arithmetic operations and mutations
- `conditional_example.go` - Conditional logic mutations  
- `loop_example.go` - Loop and iteration mutations
- `custom_mutator_example.go` - Custom mutator implementations

## Running Examples

```bash
# Run individual examples
go run arithmetic_example.go
go run conditional_example.go
go run loop_example.go

# View custom mutator code
cat custom_mutator_example.go
```

## Mutation Types Demonstrated

### Arithmetic Mutations
- `+` → `-`, `*`, `/`
- `-` → `+`, `*`, `/`
- `*` → `+`, `-`, `/`
- `/` → `+`, `-`, `*`

### Conditional Mutations
- `>` → `>=`, `<`, `<=`, `==`, `!=`
- `<` → `<=`, `>`, `>=`, `==`, `!=`
- `==` → `!=`, `>`, `<`
- `&&` → `||`
- `||` → `&&`

### Loop Mutations
- `break` → `continue`
- `continue` → `break`
- Loop condition modifications
- Iterator increment/decrement changes

### Custom Mutations
- String literal mutations
- Number boundary mutations
- Error message mutations
- Domain-specific mutations 