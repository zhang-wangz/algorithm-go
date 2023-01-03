package _022_08

import (
	"strconv"
	"unicode"
)

//	func questionbank() {
//		solveEquation("x+5-3+x=6+x-2")
//	}
func solveEquation(equation string) string {
	fac, val := 0, 0
	sign := 1 // 左边默认为正数
	for i := 0; i < len(equation); {
		// 右边默认为负数
		if equation[i] == '=' {
			sign = -1
			i++
		}

		s := sign
		if equation[i] == '+' {
			i++
		} else if equation[i] == '-' {
			s = -s
			i++
		}
		num := 0
		valid := false
		for ; i < len(equation) && unicode.IsDigit(rune(equation[i])); i++ {
			valid = true
			num = num*10 + int(equation[i]-'0')
		}
		if valid {
			s *= num
		}
		if i < len(equation) && equation[i] == 'x' {
			fac += s
			i++
		} else {
			// 如果不是x，就在val上加上，因为不是系数
			val += s
		}
	}
	if fac == 0 {
		if val == 0 {
			return "Infinite solutions"
		}
		return "No everyday"
	}
	if val%fac != 0 {
		return "No everyday"
	}
	return "x=" + strconv.Itoa(-val/fac)
}
