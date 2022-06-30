package _1_60

import "strconv"

func EvalRPN(tokens []string) int {
	j := 0
	stack := make([]int, 0)
	for j < len(tokens) {
		if tokens[j] != "+" && tokens[j] != "-" && tokens[j] != "*" && tokens[j] != "/" {
			num, _ := strconv.Atoi(tokens[j])
			stack = append(stack, num)
		} else {
			b := stack[len(stack)-1]
			a := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			ans := 0
			switch tokens[j] {
			case "+":
				ans = a + b
			case "-":
				ans = a - b
			case "*":
				ans = a * b
			case "/":
				ans = a / b
			}
			stack = append(stack, ans)
		}
		j++
	}
	return stack[len(stack)-1]
}
