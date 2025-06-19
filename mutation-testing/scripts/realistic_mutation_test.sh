#!/bin/bash

echo "=== Realistic Mutation Testing Demo ==="
echo "Let's show mutations that can survive even in tested functions"
echo

# Demo klasörüne geç
cd demo/01-calculator-basic

echo "1. Current Test Status:"
echo "---------------------"
go test -v -cover
echo

echo "2. Realistic Mutation Testing:"
echo "-----------------------------"

# Backup original file
cp calculator.go calculator.go.backup

mutation_count=0
alive_mutations=0

echo
echo "📋 Critical mutations in tested functions:"

# Mutation 1: Remove memory assignment in Add function
echo "🧬 Mutation 1: Remove memory assignment in Add function"
sed 's/c.memory = result//g' calculator.go.backup > calculator.go
mutation_count=$((mutation_count + 1))
if go test -q > /dev/null 2>&1; then
    echo "   ❌ MUTATION SURVIVED - Memory assignment not tested!"
    alive_mutations=$((alive_mutations + 1))
else
    echo "   ✅ Killed"
fi

# Mutation 2: Remove memory assignment in Multiply function  
echo "🧬 Mutation 2: Remove memory assignment in Multiply function"
sed '/result := a \* b/a\
# c.memory = result' calculator.go.backup | sed 's/c.memory = result//g' > calculator.go
mutation_count=$((mutation_count + 1))
if go test -q > /dev/null 2>&1; then
    echo "   ❌ MUTATION SURVIVED - Multiply memory assignment not tested!"
    alive_mutations=$((alive_mutations + 1))
else
    echo "   ✅ Killed"
fi

# Mutation 3: Change error message in Divide function
echo "🧬 Mutation 3: Change Divide error message"
sed 's/division by zero error/another error/g' calculator.go.backup > calculator.go
mutation_count=$((mutation_count + 1))
if go test -q > /dev/null 2>&1; then
    echo "   ❌ MUTATION SURVIVED - Error message not tested!"
    alive_mutations=$((alive_mutations + 1))
else
    echo "   ✅ Killed"
fi

# Mutation 4: Change Add function calculation (subtle)
echo "🧬 Mutation 4: Replace a+b with a+b+0 in Add function (equivalent change)"
sed 's/result := a + b/result := a + b + 0/g' calculator.go.backup > calculator.go
mutation_count=$((mutation_count + 1))
if go test -q > /dev/null 2>&1; then
    echo "   ❌ MUTATION SURVIVED - Mathematically equivalent change!"
    alive_mutations=$((alive_mutations + 1))
else
    echo "   ✅ Killed"
fi

# Mutation 5: Remove memory assignment in Divide function
echo "🧬 Mutation 5: Remove memory assignment in Divide function"
# This is more complex because Divide function has memory assignment in two places
sed '/result := a \/ b/,/c.memory = result/s/c.memory = result//g' calculator.go.backup > calculator.go
mutation_count=$((mutation_count + 1))
if go test -q > /dev/null 2>&1; then
    echo "   ❌ MUTATION SURVIVED - Divide memory assignment not tested!"
    alive_mutations=$((alive_mutations + 1))
else
    echo "   ✅ Killed"
fi

# Restore original file
mv calculator.go.backup calculator.go

echo
echo "=== MUTATION TESTING RESULTS ==="
echo "📊 Total Mutations: $mutation_count"
echo "💀 Killed: $((mutation_count - alive_mutations))"
echo "🔴 Survived: $alive_mutations"

if command -v bc > /dev/null; then
    mutation_score=$(echo "scale=2; ($mutation_count - $alive_mutations) * 100 / $mutation_count" | bc -l)
    echo "🎯 Mutation Score: $mutation_score%"
else
    echo "🎯 Mutation Score: $(( ($mutation_count - $alive_mutations) * 100 / $mutation_count ))%"
fi

echo
if [ $alive_mutations -gt 0 ]; then
    echo "⚠️  WARNING: $alive_mutations mutations survived!"
    echo
    echo "🔍 This indicates:"
    echo "  ✅ Basic function operations are tested"
    echo "  ❌ But side effects are not tested"
    echo "  ❌ Error messages are not validated"
    echo "  ❌ Memory state changes are not controlled"
    echo
    echo "💡 Improvement suggestions:"
    echo "  - Add tests that verify memory state"
    echo "  - Write tests that check error messages"
    echo "  - Test function side effects"
    echo
else
    echo "🎉 Perfect! All mutations were killed."
    echo "Your test suite is very comprehensive for these functions."
fi

echo
echo "📚 Mutation Testing Benefits:"
echo "  1. Measures test quality (not just coverage)"
echo "  2. Finds missing test cases"
echo "  3. Detects dead code"
echo "  4. Shows real effectiveness of test suite"
echo
echo "🎯 YouTube Video Notes:"
echo "  - Even 100% coverage can have surviving mutations"
echo "  - Mutation testing is the real measure of test quality"
echo "  - go-mutesting automatically creates hundreds of mutations"
echo "  - Each surviving mutation represents a potential bug" 