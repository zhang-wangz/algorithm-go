package main

import "fmt"
type test struct {
	tmap map[string]map[string]string
}
func main() {
	// nums := []int{9, 2, 5, 4, 8, 2}
	// nums = solution.HeapSort(nums)
	// for _, val := range nums {
	// 	println(val)
	// }
	t := test{}
	key := "234"
	t.tmap = map[string]map[string]string{}
	t.tmap[key] = map[string]string{"2":"3"}
	fmt.Print(t)
}
