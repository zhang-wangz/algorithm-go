package main

import (
	"algorithm-go/leetcode/solution"
)

func main() {
	//"fooleetbar"
	//"f00l"
	//[["o","0"]]
	//words := []string{"abc", "deq", "mee", "aqq", "dkd", "ccc"}
	//pattern := "abb"
	// [62,100,4]
	//2
	nums := []int{1, 3, 1}
	k := 1
	res := solution.SmallestDistancePair(nums, k)
	println(res)
}
