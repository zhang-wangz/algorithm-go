package main

import (
	"algorithm-go/slidingwindow/solution"
	"fmt"
)

func main() {
	s := "abcabcbb"
	res := solution.LengthOfLongestSubstring(s)
	fmt.Println(res)
}
