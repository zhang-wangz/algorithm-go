package main

import "strconv"

func maximumSwap(num int) int {
	str := []byte(strconv.Itoa(num))
	id1, id2, mx := -1, -1, len(str)-1
	for i := len(str) - 1; i >= 0; i-- {
		if str[i] > str[mx] {
			mx = i
		} else if str[i] < str[mx] {
			id1, id2 = i, mx
		}
	}
	if id1 == -1 {
		return num
	}
	str[id1], str[id2] = str[id2], str[id1]
	v, _ := strconv.Atoi(string(str))
	return v
}
