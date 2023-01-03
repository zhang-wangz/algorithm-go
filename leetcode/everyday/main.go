package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	s := "2121 sssd 2 343 sdr"
	r := strings.FieldsFunc(s, func(r rune) bool {
		return !unicode.IsDigit(r)
	})
	for _, w := range r {
		fmt.Println(w)
	}
}
