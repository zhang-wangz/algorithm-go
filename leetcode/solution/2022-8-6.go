package main

import (
	"sort"
	"strings"
)

func makeLargestSpecial(s string) string {
	if len(s) <= 2 {
		return s
	}
	cnt := 0
	str := sort.StringSlice{}
	for left, cur := 0, 0; cur < len(s); cur++ {
		if s[cur] == '1' {
			cnt++
		} else if cnt--; cnt == 0 {
			str = append(str, "1"+makeLargestSpecial(s[left+1:cur])+"0")
			left = cur + 1
		}
	}
	sort.Sort(sort.Reverse(str))
	return strings.Join(str, "")
}
