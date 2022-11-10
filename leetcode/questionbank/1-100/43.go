package main

import "strconv"

func multiply(num1 string, num2 string) (ans string) {
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	m, n := len(num1), len(num2)
	for i := n - 1; i >= 0; i-- {
		cur := ""
		for j := n - 1; j > i; j-- {
			cur += "0"
		}
		add := 0
		y := int(num2[i] - '0')
		for j := m - 1; j >= 0; j-- {
			// 单个相乘
			x := int(num1[j] - '0')
			product := x*y + add
			cur = strconv.Itoa(product%10) + cur
			add = product / 10
		}
		for ; add != 0; add /= 10 {
			cur = strconv.Itoa(add%10) + cur
		}
		ans = addStrings(ans, cur)
	}
	return
}

func addStrings(num1, num2 string) string {
	i, j := len(num1)-1, len(num2)-1
	add, ans := 0, ""
	for ; i >= 0 || j >= 0 || add != 0; i, j = i-1, j-1 {
		x, y := 0, 0
		if i >= 0 {
			x = int(num1[i] - '0')
		}
		if j >= 0 {
			y = int(num2[j] - '0')
		}
		pro := x + y + add
		ans = strconv.Itoa(pro%10) + ans
		add = pro / 10
	}
	return ans
}
