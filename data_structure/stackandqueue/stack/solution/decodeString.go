package solution

import "strconv"

// 给定一个经过编码的字符串，返回它解码后的字符串。
// 编码规则为: k[encoded_string]，表示其中方括号内部的 encoded_string 正好重复 k 次。注意 k 保证为正整数。
// 你可以认为输入字符串总是有效的；输入字符串中没有额外的空格，且输入的方括号总是符合格式要求的。
// 此外，你可以认为原始数据不包含数字，所有的数字只表示重复的次数 k ，例如不会出现像 3a 或 2[4] 的输入。
// 输入：s = "3[a]2[bc]"
// 输出："aaabcbc"

func DecodeString(s string) string {
	if len(s) == 0 {
		return ""
	}
	stack := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case ']':
			tmp := make([]byte, 0)
			for stack[len(stack)-1] != '[' {
				tmp = append(tmp, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			// pop [
			stack = stack[:len(stack)-1]
			// 统计数字
			idx := 1
			for len(stack)-idx >= 0 && (stack[len(stack)-idx]-'0') >= 0 && (stack[len(stack)-idx]-'0') <= 9 {
				idx++
			}
			// 因为idx++，所以这里减去idx需要+1
			num := stack[len(stack)-idx+1:]
			stack = stack[:len(stack)-idx+1]

			cnt, _ := strconv.Atoi(string(num))
			for i := 0; i < cnt; i++ {
				for j := len(tmp) - 1; j >= 0; j-- {
					stack = append(stack, tmp[j])
				}
			}
		default:
			stack = append(stack, s[i])
		}
	}
	return string(stack)
}
