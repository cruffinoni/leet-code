package main

import (
	"log"
	"strings"
)

func parseInnerExpr(expression string, f func(a, b bool) bool) bool {
	log.Printf("Parsing inner expression '%v'", expression)
	expression = strings.ReplaceAll(expression, ",", "")
	l := len(expression)
	if l == 1 {
		return expression[0] == 't'
	}
	cpy := make([]byte, l)
	copy(cpy, expression)
	lastRes := false
	for i := 1; i < l; i++ {
		res := f(cpy[i-1] == 't', cpy[i] == 't')
		cpy[i] = []byte(writeBool(res))[0]
		lastRes = res
	}
	log.Printf("Inner expression result: %v", lastRes)
	return lastRes
}

func parseExpr(expression string, length int) bool {
	log.Printf("Parsing expression '%v' (length=%d/%d)", expression, length, len(expression))
	switch expression[0] {
	case 't':
		return true
	case 'f':
		return false
	case '!':
		return expression[2] != 't'
	case '&':
		return parseInnerExpr(expression[2:length-1], func(a, b bool) bool {
			return a && b
		})
	case '|':
		return parseInnerExpr(expression[2:length-1], func(a, b bool) bool {
			return a || b
		})
	default:
		log.Fatalf("Unknown operator: %c", expression[0])
	}
	return false
}

func writeBool(b bool) string {
	if b {
		return "t"
	}
	return "f"
}

func parseBoolExpr(expression string) bool {
	log.Printf("Original expression: %v", expression)
	lastOpenPar := -1
	for i := 0; i < len(expression); i++ {
		log.Printf("Char: %c at %d", expression[i], i)
		if expression[i] == '(' {
			lastOpenPar = i
		}
		if expression[i] == ')' && lastOpenPar != -1 && i+1 < len(expression) {
			log.Printf("Last open parenthesis at %d", lastOpenPar)
			res := parseExpr(expression[lastOpenPar-1:i+1], i-lastOpenPar+2)
			expression = expression[:lastOpenPar-1] + writeBool(res) + expression[i+1:]
			log.Printf("New expression: %v", expression)
			i = 0
		}
	}
	return parseExpr(expression, len(expression))
}

type test struct {
	expression string
	expected   bool
}

func main() {
	tests := []test{
		//{expression: "!(f)", expected: true},
		//{expression: "|(f,t)", expected: true},
		//{expression: "&(t,f)", expected: false},
		//{expression: "&(t,t,t)", expected: true},
		//{expression: "&(t,f,t,t,t)", expected: false},
		//{expression: "&(t,f,t,t,t)", expected: false},
		//{expression: "&(t,t,t,t,t,!(t))", expected: false},
		//{expression: "|(&(t,f,t),!(t))", expected: false},
		{expression: "!(&(!(&(f)),&(t),|(f,f,t)))", expected: false},
	}
	for _, t := range tests {
		res := parseBoolExpr(t.expression)
		if res != t.expected {
			log.Printf("Test failed for %v, expected %v, got %v", t.expression, t.expected, res)
		} else {
			log.Printf("Test passed for %v", t.expression)
		}
	}
}
