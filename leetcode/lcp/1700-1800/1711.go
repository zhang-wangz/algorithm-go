package main

func countPairs(deliciousness []int) (ans int) {
	cnt := map[int]int{}
	ms := deliciousness[0]
	for _, v := range deliciousness {
		if v > ms {
			ms = v
		}
	}
	ts := ms * 2
	for _, v := range deliciousness {
		for sum := 1; sum <= ts; sum <<= 1 {
			ans += cnt[sum-v]
		}
		cnt[v]++
	}
	return ans
}
