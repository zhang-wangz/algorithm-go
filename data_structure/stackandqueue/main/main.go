package main

import "algorithm-pattern/data_structure/stackandqueue/queue/solution"

func main() {
	//s := "3[a]2[bc]"
	//res := solution.DecodeString(s)
	//println(res)
	s := [][]int{{0, 0, 0}, {0, 1, 0}, {1, 1, 1}}
	mat := solution.UpdateMatrix(s)
	for _, v := range mat {
		for _, v1 := range v {
			print(v1, " ")
		}
		println()
	}
}
