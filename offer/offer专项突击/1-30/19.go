package offer专项突击_1_30

func ValidPalindrome(s string) bool {
	return validstring(0, len(s)-1, s, false)
}
func validstring(start, end int, s string, flag bool) bool {
	for start <= end {
		if s[start] == s[end] {
			start++
			end--
		} else {
			if flag {
				return false
			} else {
				return validstring(start+1, end, s, !flag) || validstring(start, end-1, s, !flag)
			}
		}
	}
	return true
}
