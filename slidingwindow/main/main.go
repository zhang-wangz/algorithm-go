package main

import "fmt"

func main() {
	s := "sdads"
	for _, e := range s {
		fmt.Printf("%T", e)
	}
}
