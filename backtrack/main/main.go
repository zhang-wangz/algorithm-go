package main

import "algorithm-pattern/backtrack/challenge"

func main() {
	//challenge.LetterCombinations("23")
	//a := make([]challenge.Ans, 1)
	//a[0].Cnt++
	//print(a[0].Cnt)
	s := "aab"
	r := challenge.Partition(s)
	println(r)
}
