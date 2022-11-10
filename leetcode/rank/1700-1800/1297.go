package main

func maxFreq(s string, maxLetters int, minSize int, maxSize int) (ans int) {
	n := len(s)
	cnt := map[string]int{}
	valid := func(substr string) bool {
		mp := map[byte]bool{}
		for _, c := range substr {
			mp[byte(c)] = true
		}
		return len(mp) <= maxLetters
	}
	for i := 0; i+minSize <= n; i++ {
		if sub := s[i : i+minSize]; valid(sub) {
			cnt[sub]++
			if cnt[sub] > ans {
				ans = cnt[sub]
			}
		}
	}
	return
}

func maxFreq2(s string, maxLetters int, minSize int, maxSize int) (ans int) {
	cnt := map[byte]int{}
	m := map[string]int{}
	left, right := 0, 0
	n := len(s)
	kind := 0
	for right < n {
		c := s[right]
		cnt[c]++
		right++
		if cnt[c] == 1 {
			kind++
		}
		for kind > maxLetters || right-left > minSize {
			cnt[s[left]]--
			if cnt[s[left]] == 0 {
				kind--
			}
			left++
		}
		if kind <= maxLetters {
			if right-left == minSize {
				str := s[left:right]
				m[str]++
			}
		}
	}
	for _, v := range m {
		if v > ans {
			ans = v
		}
	}
	return
}
