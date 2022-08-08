package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	isDebug := false

	var n, m int
	var strs []string
	var s string
	if isDebug {
		n = 3
		m = 5
		strs = []string{
			"#.#.#",
			"#####",
			"#.#.#",
		}
	} else {
		fmt.Fscan(in, &n, &m)
	}

	dp := make([]int, 1<<m)
	ndp := make([]int, 1<<m)
	hrz := make([]int, 1<<m)
	vtc := make([]int, 1<<m)
	for i := range ndp {
		ndp[i] = math.MaxInt32
	}
	ndp[0] = 0
	cMax := 1 << m
	for i := 1; i < cMax; i++ {
		lowBit := i & -i
		hrz[i] = hrz[i^lowBit]
		if (lowBit<<1)&i == 0 {
			hrz[i]++
		}
		vtc[i] = vtc[i^lowBit] + 1
	}

	for i := 0; i < n; i++ {
		var cs int
		if isDebug {
			s = strs[i]
		} else {
			fmt.Fscan(in, &s)
		}
		for _, c := range s {
			cs = cs << 1
			if c == '#' {
				cs |= 1
			}
		}
		for ci := 0; ci < cMax; ci++ {
			cans := math.MaxInt32
			if (ci & cs) == ci {
				for pre := ci; ; pre = (pre - 1) & ci {
					cans = min(cans, ndp[pre]+vtc[pre^ci])
					if pre == 0 {
						break
					}
				}
			}
			dp[ci] = cans + hrz[ci^cs]
		}
		for ci := 0; ci < cMax; ci++ {
			cans := math.MaxInt32
			for mask, pre := ci^(cMax-1), ci^(cMax-1); ; pre = (pre - 1) & mask {
				cans = min(cans, dp[ci|pre])
				if pre == 0 {
					break
				}
			}
			ndp[ci] = cans
		}
	}
	fmt.Fprintln(out, ndp[0])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
