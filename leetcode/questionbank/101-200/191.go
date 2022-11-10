package main

import "math/bits"

func hammingWeight(num uint32) int {
	return bits.OnesCount(uint(num))
}
