package main

import "algorithm-go/base/kmp"

func main() {
	a := "ABABABCAA"
	b := "ABABC"
	next := kmp.GetNext(b)
	kmp.Kmp(next, b, a)
}
