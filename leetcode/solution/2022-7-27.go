package main

import (
	"strconv"
	"unicode"
)

func fractionAddition(expression string) string {
	fz, fm := 0, 1
	i := 0
	for i < len(expression) {
		sign := 1
		if expression[i] == '-' {
			sign = -1
		} else if expression[i] == '+' {
			i++
		}
		if sign == -1 {
			i++
		}
		tfz := 0
		for i < len(expression) && unicode.IsDigit(rune(expression[i])) {
			tfz = tfz*10 + int(expression[i]-'0')
			i++
		}
		tfm := 0
		i++
		for i < len(expression) && unicode.IsDigit(rune(expression[i])) {
			tfm = tfm*10 + int(expression[i]-'0')
			i++
		}
		fz = sign*tfz*fm + fz*tfm
		fm = tfm * fm
	}
	if fz == 0 {
		fm = 1
	} else {
		g := 0
		if fz < 0 {
			g = gcd(fm, -fz)
		} else {
			g = gcd(fm, fz)
		}
		fm /= g
		fz /= g
	}
	return strconv.Itoa(fz) + "/" + strconv.Itoa(fm)
}

func gcd(a, b int) int {
	if a < b {
		a, b = b, a
	}
	if b == 0 {
		return a
	}
	return gcd(a%b, b)
}
