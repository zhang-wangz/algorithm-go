package main

import "algorithm-go/backtrack/challenge"

func main() {
	//challenge.LetterCombinations("23")
	//a := make([]challenge.Ans, 1)
	//a[0].Cnt++
	//print(a[0].Cnt)
	s := "25525511135"
	r := challenge.RestoreIpAddresses(s)
	println(r)
}
