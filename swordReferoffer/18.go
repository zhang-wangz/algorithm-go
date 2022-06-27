package swordReferoffer

func IsPalindrome(s string) bool {
	if s == "" {
		return true
	}
	i, j := 0, len(s)-1
	for i < j {
		if valid(s[i]) && valid(s[j]) {
			if s[i] >= '0' && s[i] <= '9' && s[j] >= '0' && s[j] <= '9' {
				if s[i] == s[j] {
					i++
					j--
				} else {
					return false
				}
			} else {
				t1 := s[i]
				t2 := s[j]
				if s[i] >= 'A' && s[i] <= 'Z' {
					t1 = byte(s[i] + 32)
				}
				if s[j] >= 'A' && s[j] <= 'Z' {
					t2 = byte(s[j] + 32)
				}
				if t1 == t2 {
					i++
					j--
				} else {
					return false
				}
			}
		} else {
			if !valid(s[i]) {
				i++
			}
			if !valid(s[j]) {
				j--
			}
		}
	}
	return true
}

func valid(byte2 byte) bool {
	if byte2 >= '0' && byte2 <= '9' {
		return true
	}
	if byte2 >= 'a' && byte2 <= 'z' {
		return true
	}
	if byte2 >= 'A' && byte2 <= 'Z' {
		return true
	}
	return false
}
