package main

import (
	"fmt"
	"math/bits"
)

func readBinaryWatch(turnedOn int) (ans []string) {
	for i := 0; i < 12; i++ {
		for j := 0; j < 60; j++ {
			if bits.OnesCount(uint(i))+bits.OnesCount8(uint8(j)) == turnedOn {
				ans = append(ans, fmt.Sprintf("%d:%02d", i, j))
			}
		}
	}
	return
}
