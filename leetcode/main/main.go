package main

import (
	"algorithm-go/leetcode/solution"
)

func main() {
	//"fooleetbar"
	//"f00l"
	//[["o","0"]]
	words := []string{"abc", "deq", "mee", "aqq", "dkd", "ccc"}
	pattern := "abb"
	res := solution.FindAndReplacePattern(words, pattern)
	println(res)
}
