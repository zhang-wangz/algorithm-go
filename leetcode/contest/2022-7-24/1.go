package main

func bestHand(ranks []int, suits []byte) string {
	mp1 := map[byte]int{}
	for i := range suits {
		mp1[suits[i]]++
	}
	if mp1[suits[0]] == 5 {
		return "Flush"
	}
	mp2 := map[int]int{}
	for i := range ranks {
		mp2[ranks[i]]++
	}
	min := 0
	for _, v := range mp2 {
		if v >= min {
			min = v
		}
	}
	if min >= 3 {
		return "Three of a Kind"
	} else if min >= 2 {
		return "Pair"
	} else {
		return "High Card"
	}
}
