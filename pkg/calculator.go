package calculator

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

func Calc(expression string) (float64, error) {
	var opsStack []string
	var rpnQueue []string
	var stack []float64

	precedence := map[string]int{"+": 1, "-": 1, "*": 2, "/": 2}
	tokenPattern := regexp.MustCompile(`\d+(\.\d+)?|[+\-*/()]`)

	tokens := tokenPattern.FindAllString(expression, -1)

	for _, token := range tokens {
		if isOperator(token) {
			for len(opsStack) > 0 && opsStack[len(opsStack)-1] != "(" &&
				precedence[opsStack[len(opsStack)-1]] >= precedence[token] {
				rpnQueue = append(rpnQueue, opsStack[len(opsStack)-1])
				opsStack = opsStack[:len(opsStack)-1]
			}
			opsStack = append(opsStack, token)
		} else if token == "(" {
			opsStack = append(opsStack, token)
		} else if token == ")" {
			for len(opsStack) > 0 && opsStack[len(opsStack)-1] != "(" {
				rpnQueue = append(rpnQueue, opsStack[len(opsStack)-1])
				opsStack = opsStack[:len(opsStack)-1]
			}
			if len(opsStack) == 0 {
				return 0, errors.New("mismatched parentheses")
			}
			opsStack = opsStack[:len(opsStack)-1]
		} else {
			rpnQueue = append(rpnQueue, token)
		}
	}

	for len(opsStack) > 0 {
		rpnQueue = append(rpnQueue, opsStack[len(opsStack)-1])
		opsStack = opsStack[:len(opsStack)-1]
	}

	for _, token := range rpnQueue {
		if isOperator(token) {
			if len(stack) < 2 {
				return 0, errors.New("invalid expression")
			}
			b, a := stack[len(stack)-1], stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			switch token {
			case "+":
				stack = append(stack, a+b)
			case "-":
				stack = append(stack, a-b)
			case "*":
				stack = append(stack, a*b)
			case "/":
				if b == 0 {
					return 0, errors.New("division by zero")
				}
				stack = append(stack, a/b)
			}
		} else {
			num, err := strconv.ParseFloat(token, 64)
			if err != nil {
				return 0, fmt.Errorf("invalid number: %s", token)
			}
			stack = append(stack, num)
		}
	}

	if len(stack) != 1 {
		return 0, errors.New("invalid expression")
	}
	return stack[0], nil
}

func isOperator(token string) bool {
	return token == "+" || token == "-" || token == "*" || token == "/"
}
