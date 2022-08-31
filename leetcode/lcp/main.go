package main

import (
	"fmt"
	"strconv"
)

func main() {
	f, _ := strconv.ParseFloat(fmt.Sprintf("%0.2f", 2.00), 64)
	fmt.Printf("%0.2f", f)
}
