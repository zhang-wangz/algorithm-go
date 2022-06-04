package solution

import "strings"

// https://leetcode.cn/problems/unique-email-addresses/

func NumUniqueEmails(emails []string) int {
	m := map[string]int{}
	for _, email := range emails {
		str := strings.Split(email, "@")
		if ok1, s := isValidLocal(str[0]); ok1 {
			if ok2, y := isValidY(str[1]); ok2 {
				m[s+"@"+y] = 1
			}
		}
	}
	return len(m)
}
func isValidLocal(s string) (bool, string) {
	b := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if s[i] == '+' {
			break
		}
		if s[i] == '.' {
			continue
		}
		b = append(b, s[i])
	}
	return true, string(b)
}

func isValidY(s string) (bool, string) {
	//if len(strings.Split(s,"+"))!=1{
	//	return false,""
	//}
	return true, s
}
