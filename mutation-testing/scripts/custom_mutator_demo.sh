#!/bin/bash

# Custom Mutator Demo Script for YouTube Video
# Shows how custom mutators work in go-mutesting

# Colors for output
GREEN='\033[0;32m'
BLUE='\033[0;34m'
YELLOW='\033[1;33m'
RED='\033[0;31m'
NC='\033[0m' # No Color

echo -e "${BLUE}🧬 GO-MUTESTING: CUSTOM MUTATOR DEMO${NC}"
echo "=============================================="
echo ""

echo -e "${GREEN}This demo shows how custom mutators work${NC}"
echo "With custom mutators, you can create domain-specific mutations"
echo ""

echo -e "${YELLOW}1. Running Custom Mutator Demo...${NC}"
echo ""

# Run the custom demo
go run demo/03-custom-mutator/custom-demo/main.go

echo ""
echo -e "${YELLOW}2. Custom Mutator Source Code:${NC}"
echo ""

echo -e "${GREEN}📁 Our custom mutator example files:${NC}"
echo "   • demo/02-mutation-examples/custom_mutator_example.go - Real custom mutator implementations"
echo "   • demo/03-custom-mutator/custom-demo/main.go - Runnable demo"
echo ""

echo -e "${YELLOW}3. Custom Mutator Types:${NC}"
echo ""

echo -e "${GREEN}📝 String Literal Mutator:${NC}"
echo "   • Replaces string values with different values"
echo "   • Empty string, test string, error messages"
echo "   • Ideal for string validation tests"
echo ""

echo -e "${GREEN}🔢 Number Mutator:${NC}"
echo "   • Replaces numeric values with boundary values"  
echo "   • 0, 1, -1, large numbers"
echo "   • Perfect for edge case testing"
echo ""

echo -e "${GREEN}🚨 Error Message Mutator:${NC}"
echo "   • Specifically tests error messages"
echo "   • Checks test consistency with different error messages"
echo "   • Critical for domain-specific error handling"
echo ""

echo -e "${YELLOW}4. Custom Mutator Implementation:${NC}"
echo ""

echo -e "${GREEN}💻 Real usage steps:${NC}"
echo "   1. Clone go-mutesting source code"
echo "   2. Add custom mutator to mutator/registry.go file"
echo "   3. Rebuild the binary"
echo "   4. Use with --enable=custom/mutator"
echo ""

echo -e "${YELLOW}5. Advantages:${NC}"
echo ""

echo -e "${GREEN}✅ Benefits:${NC}"
echo "   • Business logic-specific tests"
echo "   • Protocol-specific mutations"
echo "   • Domain validation"
echo "   • Advanced error handling tests"
echo "   • String/number boundary tests"
echo ""

echo -e "${RED}⚠️  Note:${NC}"
echo "Custom mutators must be added to go-mutesting source code"
echo "They cannot be added via config file, code-level integration is required"
echo ""

echo -e "${BLUE}🎯 Demo completed!${NC}"
echo "With custom mutators, you can customize mutation testing for your domain-specific needs"
echo ""

echo -e "${GREEN}📖 More information:${NC}"
echo "   • github.com/avito-tech/go-mutesting"
echo "   • Custom mutator implementation guide"
echo "   • This demo: demo/03-custom-mutator/custom-demo/main.go" 