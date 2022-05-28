package main

import (
	"algorithm-pattern/leetcode/solution"
	"fmt"
)

func main() {
	//data := [][]int{{4, 2, 3}, {0, 0, 1}, {7, 6, 5}}
	//solution.CutOffTree(data)
	s := "(()())(())"
	res := solution.RemoveOuterParentheses(s)
	fmt.Println(res)
}
