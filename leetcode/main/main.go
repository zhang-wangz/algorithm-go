package main

import (
	"algorithm-pattern/leetcode/solution"
	"fmt"
)

func main() {
	//data := [][]int{{4, 2, 3}, {0, 0, 1}, {7, 6, 5}}
	//solution.CutOffTree(data)
	s := "2001:0db8:85a3:00000:0:8A2E:0370:7334"
	res := solution.ValidIPAddress(s)
	fmt.Println(res)
}
