package main

import (
	"fmt"
	"strings"
)

func main() {

	str := "[]}"
	valid := validParentheses(str)
	fmt.Printf("String is Valid: %v\n", valid)
}

func validParentheses(s string) bool {

	stack := ""

	for _, char := range s {
		if strings.Contains("({[", string(char)) {
			stack = stack + string(char)
		} else if len(stack) == 0 {
			return false
		} else {
			value := stack[len(stack)-1:]

			if char == ')' && value != "(" {
				return false
			}

			if char == '}' && value != "{" {
				return false
			}

			if char == ']' && value != "[" {
				return false
			}

			stack = stack[:len(stack)-1]

		}
	}

	if len(stack) == 0 {
		return true
	}

	return false
}
