package swordReferoffer

// https://leetcode.cn/problems/JFETK5/
// 给定两个 01 字符串 a 和 b ，请计算它们的和，并以二进制字符串的形式输出。
// 输入为 非空 字符串且只包含数字 1 和 0。
func AddBinary(a string, b string) string {
	la := len(a)
	lb := len(b)
	var l int
	if la > lb {
		l = lb
	} else {
		l = la
	}
	stack := make([]byte, 0)
	pre := 0
	for i := 0; i < l; i++ {
		sum := int(a[la-1-i]-'0') + int(b[lb-1-i]-'0') + pre
		if sum > 1 {
			stack = append(stack, byte(sum-2+'0'))
			pre = 1
		} else if sum == 1 {
			stack = append(stack, '1')
			pre = 0
		} else {
			stack = append(stack, '0')
			pre = 0
		}
	}
	if l < la {
		for i := 0; i < la-l; i++ {
			sum := int(a[la-l-1-i]-'0') + pre
			if sum > 1 {
				stack = append(stack, byte(sum-2+'0'))
				pre = 1
			} else if sum == 1 {
				stack = append(stack, '1')
				pre = 0
			} else {
				stack = append(stack, '0')
				pre = 0
			}
		}
	}
	if l < lb {
		for i := 0; i < lb-l; i++ {
			sum := int(b[lb-l-1-i]-'0') + pre
			if sum > 1 {
				stack = append(stack, byte(sum-2+'0'))
				pre = 1
			} else if sum == 1 {
				stack = append(stack, '1')
				pre = 0
			} else {
				stack = append(stack, '0')
				pre = 0
			}
		}
	}
	if pre >= 1 {
		stack = append(stack, '1')
	}
	res := ""
	for i := len(stack) - 1; i >= 0; i-- {
		res += string(stack[i])
	}
	return res
}
