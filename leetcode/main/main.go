package main

import (
	"fmt"
	"sort"
)

func main() {
	s := sort.IntSlice{}
	s = append(s, 2)
	s = append(s, 98)
	sort.Sort(sort.Reverse(s))
	fmt.Println(s.Len())
}
