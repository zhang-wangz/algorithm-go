package challenge

// https://leetcode.cn/problems/letter-combinations-of-a-phone-number/
// 给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。答案可以按 任意顺序 返回。
// 给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。

func LetterCombinations(digits string) []string {
	if digits == "" {
		return []string{}
	}
	m := map[int]string{
		2: "abc", 3: "def", 4: "ghi", 5: "jkl", 6: "mno", 7: "pqrs", 8: "tuv", 9: "wxyz",
	}
	nums := make([]int, 0)
	res := make([]string, 0)
	str := ""
	for _, num := range digits {
		n := int(num - '0')
		nums = append(nums, n)
	}
	dfsLetter(nums, str, &res, m, 0)
	return res
}

func dfsLetter(nums []int, str string, res *[]string, m map[int]string, p int) {
	if p == len(nums) {
		tmp := make([]byte, len(str))
		copy(tmp, str)
		*res = append(*res, string(tmp))
		return
	}
	letter := m[nums[p]]
	for i := 0; i < len(letter); i++ {
		dfsLetter(nums, str+string(letter[i]), res, m, p+1)
	}
}
