package main

// 模拟题
// 给你一个括号字符串s，它只包含字符'(' 和')'。一个括号字符串被称为平衡的当它满足：
// 任何左括号'('必须对应两个连续的右括号'))'。
// 左括号'('必须在对应的连续两个右括号'))'之前。
// 你可以在任意位置插入字符 '(' 和 ')' 使字符串平衡。
// 求最少的插入次数
func minInsertions(s string) (ans int) {
	l := 0
	idx := 0
	n := len(s)
	for idx < n {
		c := s[idx]
		if c == '(' {
			l++
			idx++
		} else {
			if l > 0 {
				l--
			} else {
				ans++
			}
			if idx < n-1 && s[idx+1] == ')' {
				idx += 2
			} else {
				ans++
				idx++
			}
		}
	}
	ans += l * 2
	return
}
