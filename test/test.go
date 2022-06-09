package main

import (
	"math"
)

type Project struct {
	Name string `json:"name"`
	Url  string `json:"url"`
	F    bool   `json:"f,omitempty"`
	Docs string `json:"docs,omitempty"`
}

func main() {
	a := int(math.Floor(3.7))
	print(a)
}
