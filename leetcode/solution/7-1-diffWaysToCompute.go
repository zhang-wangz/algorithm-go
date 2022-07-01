package solution

import "strconv"

// https://leetcode.cn/problems/different-ways-to-add-parentheses/
// 分治

func diffWaysToCompute(expression string) (ans []int) {
	if expression == "" || len(expression) == 0 {
		return ans
	}
	for i := 0; i < len(expression); i++ {
		if expression[i] == '+' || expression[i] == '-' || expression[i] == '*' {
			opr := expression[i]
			left := diffWaysToCompute(expression[:i])
			right := diffWaysToCompute(expression[i+1:])
			for _, l := range left {
				for _, r := range right {
					if opr == '+' {
						ans = append(ans, l+r)
					} else if opr == '-' {
						ans = append(ans, l-r)
					} else {
						ans = append(ans, l*r)
					}
				}
			}
		}
	}
	input, err := strconv.Atoi(expression)
	if err != nil {
		ans = append(ans, input)
	}
	return ans
}
