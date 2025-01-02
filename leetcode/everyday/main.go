package main

import "fmt"

func main() {
	a := make([]int, 0, 3)
	a = append(a, 2)
	a = append(a, 3)
	fmt.Println(a)
}
