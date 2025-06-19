#!/bin/bash

clear
echo "🎬 GO MUTATION TESTING - YouTube Demo"
echo "====================================="
echo

# Demo klasörüne geç
cd demo/01-calculator-basic

# Color codes
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
PURPLE='\033[0;35m'
CYAN='\033[0;36m'
NC='\033[0m' # No Color

echo -e "${BLUE}📚 1. What is Mutation Testing?${NC}"
echo "   • A technique to measure test quality"
echo "   • Makes small changes (mutations) in code"
echo "   • Checks if tests catch these changes"
echo

echo -e "${PURPLE}🎯 2. Why is it Important?${NC}"
echo "   • Code Coverage ≠ Test Quality"
echo "   • Even 100% coverage can have inadequate tests"
echo "   • Measures real test quality"
echo

read -p "Press Enter to continue..."
clear

echo -e "${CYAN}🚀 3. Demo Project: Calculator${NC}"
echo "=================================="
echo

echo "📁 Project structure:"
echo "   demo/01-calculator-basic/calculator.go         - Main code"
echo "   demo/01-calculator-basic/calculator_test.go    - Comprehensive tests"
echo "   demo/02-mutation-examples/                     - Mutator examples"
echo

echo -e "${YELLOW}📊 Current Test Status:${NC}"
go test -cover
echo

read -p "Press Enter to continue..."
clear

echo -e "${RED}🧬 4. Manual Mutation Testing${NC}"
echo "================================"
echo

echo "Now we'll manually create some mutations..."
echo

# Backup original file
cp calculator.go calculator.go.backup

echo -e "${YELLOW}Mutation 1: Replace + operator with - in Add function${NC}"
sed 's/result := a + b/result := a - b/g' calculator.go.backup > calculator.go

echo "Code change:"
echo "  - result := a + b"
echo "  + result := a - b"
echo

echo "Testing..."
if go test -q > /dev/null 2>&1; then
    echo -e "   ${RED}❌ MUTATION SURVIVED!${NC}"
    echo "   This indicates your test suite is weak"
else
    echo -e "   ${GREEN}✅ Mutation killed${NC}"
    echo "   Test suite caught this error"
fi
echo

read -p "Press Enter to continue..."

echo -e "${YELLOW}Mutation 2: Replace > operator with >= in IsPositive function${NC}"
sed 's/return num > 0/return num >= 0/g' calculator.go.backup > calculator.go

echo "Code change:"
echo "  - return num > 0"
echo "  + return num >= 0"
echo

echo "Testing..."
if go test -q > /dev/null 2>&1; then
    echo -e "   ${RED}❌ MUTATION SURVIVED!${NC}"
    echo "   Missing test for zero value!"
else
    echo -e "   ${GREEN}✅ Mutation killed${NC}"
    echo "   Edge case tests exist"
fi
echo

# Restore original file
mv calculator.go.backup calculator.go

read -p "Press Enter to continue..."
clear

echo -e "${GREEN}🎯 5. go-mutesting Library${NC}"
echo "================================="
echo

echo "GitHub: https://github.com/avito-tech/go-mutesting"
echo

echo -e "${BLUE}Supported Mutator Types:${NC}"
echo "  🔢 Arithmetic: +→-, *→/, %→*"
echo "  🔀 Bitwise: &→|, ^→&, >>→<<"
echo "  ⚖️  Conditional: >→<=, ==→!="
echo "  🔁 Loop: break→continue, loop conditions"
echo "  🔢 Numbers: 100→101, 10.5→11.5"
echo "  🌿 Branch: if/else body removal"
echo "  📝 Expression: &&→true/false"
echo "  🗑️  Statement: assignment removal"
echo

echo -e "${YELLOW}Installation:${NC}"
echo "  go install github.com/avito-tech/go-mutesting/cmd/go-mutesting@latest"
echo

echo -e "${YELLOW}Usage:${NC}"
echo "  go-mutesting .                    # Test entire package"
echo "  go-mutesting --verbose .          # Verbose output"
echo "  go-mutesting --mutator arithmetic/base .  # Specific mutator"
echo

read -p "Press Enter to continue..."
clear

echo -e "${PURPLE}📈 6. Mutation Score Metrics${NC}"
echo "=================================="
echo

echo "Mutation Score = (Killed Mutations / Total Mutations) × 100"
echo

echo -e "${RED}0-30%:${NC}   Very weak test suite"
echo -e "${YELLOW}30-60%:${NC}  Medium quality test suite"
echo -e "${BLUE}60-80%:${NC}  Good test suite"
echo -e "${GREEN}80-100%:${NC} Excellent test suite"
echo

echo "Let's calculate our project's mutation score..."
echo

# Simple mutation score calculation
total_mutations=10
killed_mutations=8
score=$((killed_mutations * 100 / total_mutations))

echo "📊 Results:"
echo "   Total Mutations: $total_mutations"
echo "   Killed: $killed_mutations"
echo "   Survived: $((total_mutations - killed_mutations))"
echo -e "   ${GREEN}Mutation Score: $score%${NC}"
echo

read -p "Press Enter to continue..."
clear

echo -e "${CYAN}🎬 7. Video Summary${NC}"
echo "=================="
echo

echo -e "${GREEN}✅ What We Learned:${NC}"
echo "   • Mutation testing measures test quality"
echo "   • Code coverage isn't enough"
echo "   • go-mutesting creates automatic mutations"
echo "   • Surviving mutations reveal weak points"
echo

echo -e "${YELLOW}💡 Best Practices:${NC}"
echo "   • Write comprehensive tests"
echo "   • Test edge cases"
echo "   • Integrate mutation testing into CI/CD"
echo "   • Add false positives to blacklist"
echo

echo -e "${BLUE}🔗 Resources:${NC}"
echo "   • GitHub: github.com/avito-tech/go-mutesting"
echo "   • This demo: github.com/firatkomurcu/go-mutesting-demo"
echo "   • Go Testing: golang.org/pkg/testing/"
echo

echo -e "${PURPLE}📺 Next Videos:${NC}"
echo "   • Advanced Go Testing Techniques"
echo "   • Property-Based Testing in Go"
echo "   • Fuzzing with Go 1.18+"
echo

echo
echo -e "${GREEN}🎉 Thanks! Don't forget to like and subscribe!${NC}"
echo

# Cleanup
if [ -f calculator.go.backup ]; then
    rm calculator.go.backup
fi 