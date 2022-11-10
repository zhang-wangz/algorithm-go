package main

import (
	"algorithm-go/base/slidingwindow/solution"
	"fmt"
)

func main() {
	s := "abcabcbb"
	res := solution.LengthOfLongestSubstring(s)
	fmt.Println(res)
}
