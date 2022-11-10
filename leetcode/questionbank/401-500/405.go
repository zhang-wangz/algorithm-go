package main

func toHex(num int) (ans string) {
	if num == 0 {
		return "0"
	}
	if num < 0 {
		num += 1 << 32
	}
	hex := "0123456789abcdef"
	for num != 0 {
		ans += string(hex[num%16])
		num /= 16
	}
	b := []byte(ans)
	for l, r := 0, len(b); l < r; l++ {
		b[l], b[r] = b[r], b[l]
		r--
	}
	return string(b)
}
