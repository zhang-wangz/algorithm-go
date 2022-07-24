package main

func repeatedCharacter(s string) byte {
	m := map[byte]int{}
	for i := 0; i < len(s); i++ {
		m[s[i]]++
		if m[s[i]] == 2 {
			return s[i]
		}
	}
	return 's'
}
