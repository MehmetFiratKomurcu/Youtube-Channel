# Go Mutation Testing Demo - YouTube Video Content

Source Code for [Mutation Testing in Golang with go-mutesting](https://youtu.be/BlEHEupNVmQ )

[![Mutation Testing in Golang with go-mutesting Youtube Video Link](https://img.youtube.com/vi/BlEHEupNVmQ/0.jpg)](https://www.youtube.com/watch?v=BlEHEupNVmQ )

This project was created to explain **Mutation Testing** concepts and the `go-mutesting` library in the Go programming language.

## 📹 YouTube Video Content

### 1. What is Mutation Testing?

**Mutation Testing** is a technique used to measure test quality. Basic principle:

- Makes small changes (mutations) in your code
- Checks whether these changes are caught by your tests
- Uncaught mutations reveal weak points in your test suite

### 2. Why is it Important?

```
Code Coverage ≠ Test Quality
```

- Even with **100% code coverage**, your tests might be insufficient
- Mutation testing measures the **real quality** of your tests
- Finds missing test cases
- Detects dead code

### 3. go-mutesting Library

GitHub: https://github.com/avito-tech/go-mutesting

#### Features:
- **Automatic mutation generation**
- **Multiple mutator support**
- **Detailed reporting**
- **Configurable** (YAML config)
- **CI/CD integration**

#### Supported Mutator Types:

##### 🔢 Arithmetic Mutators
- `+` → `-`, `-` → `+`
- `*` → `/`, `/` → `*`
- `%` → `*`

##### 🔀 Bitwise Mutators  
- `&` → `|`, `|` → `&`
- `^` → `&`, `&^` → `&`
- `>>` → `<<`, `<<` → `>>`

##### 🔄 Assignment Mutators
- `+=` → `=`, `-=` → `=`
- `*=` → `=`, `/=` → `=`

##### 🔁 Loop Mutators
- `break` → `continue`
- `continue` → `break`
- Loop condition mutations

##### 🔢 Number Mutators
- `100` → `101` (incrementer)
- `100` → `99` (decrementer)

##### ⚖️ Conditional Mutators
- `>` → `<=`, `<` → `>=`
- `==` → `!=`, `!=` → `==`

##### 🌿 Branch Mutators
- Empties `if` branches
- Empties `else` branches
- Empties `case` bodies

##### 📝 Expression Mutators
- `&&` → `true`/`false`
- `||` → `true`/`false`

##### 🗑️ Statement Mutators
- Removes assignment statements
- Removes expression statements

## 🚀 Demo Project: Calculator

### Quick Start

#### Option 1: Install and Run (Recommended)
```bash
# Install the demo calculator
go install github.com/firatkomurcu/go-mutesting-demo/cmd/calculator@latest

# Run the demo
calculator
```

#### Option 2: Clone and Explore
```bash
# Clone the repository
git clone https://github.com/firatkomurcu/go-mutesting-demo.git
cd go-mutesting-demo

# Run the calculator demo
go run demo/01-calculator-basic/cmd/calculator/main.go

# Run tests
cd demo/01-calculator-basic && go test -v

# Try mutation testing demos
./scripts/youtube_demo.sh
./scripts/mixed_mutation_demo.sh
./scripts/custom_mutator_demo.sh
```

### Project Structure
```
go-mutesting-demo/
├── demo/                           # Demo projects
│   ├── 01-calculator-basic/        # Basic calculator demo
│   │   ├── calculator.go           # Calculator library
│   │   ├── calculator_test.go      # Comprehensive tests
│   │   ├── cmd/calculator/main.go  # Executable binary
│   │   └── README.md               # Calculator demo guide
│   ├── 02-mutation-examples/       # Mutation examples
│   │   ├── arithmetic_example.go   # Arithmetic mutations
│   │   ├── conditional_example.go  # Conditional mutations
│   │   ├── loop_example.go         # Loop mutations
│   │   ├── custom_mutator_example.go # Custom mutator code
│   │   └── README.md               # Examples guide
│   └── 03-custom-mutator/          # Custom mutator demo
│       ├── custom-demo/main.go     # Interactive custom demo
│       ├── go.mod                  # Module with dependencies
│       └── README.md               # Custom mutator guide
├── scripts/                        # Demo scripts
│   ├── youtube_demo.sh             # Interactive demo
│   ├── mixed_mutation_demo.sh      # Mixed results demo
│   ├── realistic_mutation_test.sh  # Realistic testing
│   └── custom_mutator_demo.sh      # Custom mutator demo
├── config/                         # Configuration files
│   └── config.yml                  # Basic configuration
├── docs/                           # Documentation
│   └── YOUTUBE_VIDEO_KAPSAMLI_REHBER.md # Complete video guide
└── README.md                       # This file
```

### Calculator Features
- ✅ Basic math operations (Add, Subtract, Multiply, Divide)
- ✅ Memory management (GetMemory, ClearMemory)
- ✅ Power calculation
- ✅ Utility functions (IsPositive, Max, Min)
- ✅ Error handling (division by zero)

## 📊 Test Quality Comparison

### Weak Test Suite (14.5% coverage)
```bash
# Only 3 basic tests
go test -cover
# coverage: 14.5% of statements
```

### Comprehensive Test Suite (68.9% coverage)
```bash
# All functions + edge cases
go test -cover  
# coverage: 68.9% of statements
```

## 🧬 Mutation Testing Results

### Manual Mutation Testing
```bash
./scripts/realistic_mutation_test.sh
```

This script tests:
1. **Memory assignment** mutations
2. **Error message** mutations  
3. **Mathematical equivalence** mutations
4. **Side effect** mutations

### Real go-mutesting Usage

#### Installation:
```bash
go install github.com/avito-tech/go-mutesting/cmd/go-mutesting@latest
```

#### Basic Usage:
```bash
# Test entire package
cd demo/01-calculator-basic
go-mutesting .

# Test specific files
go-mutesting calculator.go

# Verbose output
go-mutesting --verbose .

# Use specific mutators
go-mutesting --mutator arithmetic/base .
```

#### Config File (config/config.yml):
```yaml
skip_without_test: true
skip_with_build_tags: true
json_output: true
silent_mode: false
exclude_dirs: ["vendor", "testdata"]

# Blacklist specific functions/files
blacklist:
  - function: "main"
    reason: "Entry points don't need testing"
  - function: "String"
    reason: "Formatters are usually simple"
  - pattern: 'log\.Print.*'
    reason: "Logging doesn't affect business logic"

# Custom mutation patterns
custom_mutations:
  - name: "empty_string"
    pattern: '"[^"]{1,}"'
    replacement: '""'
    description: "Replace strings with empty"
```

## 🎯 Advanced Features

### Blacklist Feature
Exclude specific code from mutation testing:

```bash
# Run with blacklist demo
./blacklist_custom_demo.sh
```

**Common Blacklist Targets:**
- `main()` functions - Entry points
- `init()` functions - Initialization code  
- `String()` methods - Display formatters
- Constructor functions (`New*`)
- Logging statements
- Mock files (`*_mock.go`)

### Custom Mutations
Create domain-specific mutation patterns:

```yaml
custom_mutations:
  - name: "error_messages"
    pattern: '".*error.*"'
    replacement: '"mutation error"'
    description: "Change error message content"
    
  - name: "boolean_flip"
    pattern: '\btrue\b' 
    replacement: 'false'
    description: "Flip boolean values"
```

### Configuration Files
- `config.yml` - Basic configuration with simple blacklist
- `advanced_config.yml` - Advanced blacklist and custom mutations

## 📈 Mutation Score Metrics

**Mutation Score = (Killed Mutations / Total Mutations) × 100**

- **0-30%**: Very weak test suite
- **30-60%**: Medium quality test suite  
- **60-80%**: Good test suite
- **80-100%**: Excellent test suite

## 🎯 YouTube Video Scenario

### Part 1: Introduction (2-3 min)
- What is mutation testing?
- Why is it important?
- Code coverage vs test quality

### Part 2: go-mutesting Introduction (3-4 min)
- Library features
- Mutator types
- Installation

### Part 3: Demo Project (5-6 min)
- Calculator project introduction
- Weak tests vs comprehensive tests
- Coverage comparison

### Part 4: Mutation Testing Practice (8-10 min)
- Manual mutation testing
- go-mutesting usage
- Results analysis
- Examining surviving mutations

### Part 5: Best Practices (3-4 min)
- Mutation testing strategies
- CI/CD integration
- Performance considerations
- Managing false positives

### Part 6: Conclusion (1-2 min)
- Key points
- Resources
- Future videos

## 🔗 Useful Resources

- [go-mutesting GitHub](https://github.com/avito-tech/go-mutesting)
- [Mutation Testing Wikipedia](https://en.wikipedia.org/wiki/Mutation_testing)
- [Go Testing Documentation](https://golang.org/pkg/testing/)

## 💡 Video Notes

### Key Points to Emphasize:
1. **Coverage ≠ Quality**: Even 100% coverage isn't enough
2. **Mutation Score**: Real test quality metric
3. **Automatic Detection**: go-mutesting creates hundreds of mutations
4. **Practical Value**: Each surviving mutation represents a potential bug

### Demo Sequence:
1. Start with weak tests (14.5% coverage)
2. Show manual mutations
3. Switch to comprehensive tests (68.9% coverage)
4. Use go-mutesting for automatic testing
5. Compare results

### Code Examples:
- Simple mutations (+ → -, > → >=)
- Complex mutations (memory assignment)
- Edge case mutations (division by zero)

## 🎬 Video Recording Notes

- **Screen recording**: Terminal and code editor
- **Live coding**: Manually show mutations
- **Result analysis**: Explain mutation scores
- **Practical examples**: Real-world scenarios

---

**This project is designed to demonstrate the power and importance of Go mutation testing. In the YouTube video, we'll process this content step by step to provide viewers with comprehensive mutation testing education.** 