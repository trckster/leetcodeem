package main

import "strconv"

func evalRPN(tokens []string) int {
	stack := make([]int, 0)

	for _, token := range tokens {
		switch token {
		case "+":
			number1 := stack[len(stack)-2]
			number2 := stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			stack = append(stack, number1+number2)
		case "-":
			number1 := stack[len(stack)-2]
			number2 := stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			stack = append(stack, number1-number2)
		case "*":
			number1 := stack[len(stack)-2]
			number2 := stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			stack = append(stack, number1*number2)
		case "/":
			number1 := stack[len(stack)-2]
			number2 := stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			stack = append(stack, number1/number2)
		default:
			number, _ := strconv.Atoi(token)
			stack = append(stack, number)
		}
	}

	return stack[0]
}
