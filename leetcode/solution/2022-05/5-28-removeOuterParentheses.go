package _022_05

// https://leetcode.cn/problems/remove-outermost-parentheses/
// 有效括号字符串为空 ""、"(" + A + ")" 或 A + B ，其中 A 和 B 都是有效的括号字符串，+ 代表字符串的连接。
//
//例如，""，"()"，"(())()" 和 "(()(()))" 都是有效的括号字符串。
//如果有效字符串 s 非空，且不存在将其拆分为 s = A + B 的方法，我们称其为原语（primitive），其中 A 和 B 都是非空有效括号字符串。
//
//给出一个非空有效字符串 s，考虑将其进行原语化分解，使得：s = P_1 + P_2 + ... + P_k，其中 P_i 是有效括号字符串原语。
//
//对 s 进行原语化分解，删除分解中每个原语字符串的最外层括号，返回 s 。
func RemoveOuterParentheses(s string) string {
	if s == "" {
		return ""
	}
	stack := make([]byte, 0)
	cnt := 0
	res := ""
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			cnt++
			stack = append(stack, s[i])
		} else if s[i] == ')' {
			cnt--
			// 如果等于0，说明要加了
			if cnt == 0 {
				for j := 1; j < len(stack); j++ {
					res += string(stack[j])
				}
				// pop (
				stack = stack[:0]
				continue
			}
			stack = append(stack, s[i])
		}
	}
	return res
}
