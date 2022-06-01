package main

import (
	"algorithm-pattern/slidingwindow/solution"
	"fmt"
)

func main() {
	s := "abcabcbb"
	res := solution.LengthOfLongestSubstring(s)
	fmt.Println(res)
}
