package main

import "sort"

func main() {
	a := sort.IntSlice{}
	a = append(a, 2)
	a = append(a, 1)
	sort.Sort(a)
	print(a[0])
}
