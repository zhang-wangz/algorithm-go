package contest

func strongPasswordCheckerII(password string) bool {
	sp := "!@#$%^&*()-+"
	spf := false
	low := false
	high := false
	d := false
	if len(password) < 8 {
		return false
	}
	m := map[byte]bool{}
	for i := 0; i < len(sp); i++ {
		m[sp[i]] = true
	}
	for i := 0; i < len(password); i++ {
		if m[password[i]] {
			spf = true
		}
		if password[i] >= 'a' && password[i] <= 'z' {
			low = true
		} else if password[i] >= 'A' && password[i] <= 'Z' {
			high = true
		} else if password[i] >= '0' && password[i] <= '9' {
			d = true
		}
		if i != 0 && password[i] == password[i-1] {
			return false
		}

	}
	if high && low && spf && d {
		return true
	}
	return false
}
