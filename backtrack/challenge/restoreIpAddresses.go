package challenge

import "strings"

// https://leetcode.cn/problems/restore-ip-addresses/
// 有效 IP 地址 正好由四个整数（每个整数位于 0 到 255 之间组成，且不能含有前导 0），整数之间用 '.' 分隔。
//
// 例如："0.1.2.201" 和 "192.168.1.1" 是 有效 IP 地址，
// 但是 "0.011.255.245"、"192.168.1.312" 和 "192.168@1.1" 是 无效 IP 地址。
// 给定一个只包含数字的字符串 s ，用以表示一个 IP 地址，返回所有可能的有效 IP 地址，
// 这些地址可以通过在 s 中插入 '.' 来形成。你 不能 重新排序或删除 s 中的任何数字。你可以按 任何 顺序返回答案。

func RestoreIpAddresses(s string) []string {
	res := make([]string, 0)
	list := make([]string, 0)
	if len(s) > 12 {
		return res
	}
	dfs6(s, &res, list)
	return res
}

func dfs6(s string, res *[]string, list []string) {
	if len(list) > 4 {
		return
	}
	if len(s) > (4-len(list))*3 {
		return
	}
	if s == "" && len(list) == 4 {
		ans := make([]string, len(list))
		copy(ans, list)
		*res = append(*res, strings.Join(ans, "."))
	}
	for i := 0; i < len(s); i++ {
		str := s[:i+1]
		if isValidIP(str) {
			list = append(list, str)
			dfs6(s[i+1:], res, list)
			list = list[:len(list)-1]
		}
	}
}

func isValidIP(s string) bool {
	num := 0
	for i := 0; i < len(s); i++ {
		if s[0] == '0' && len(s) != 1 {
			return false
		}
		if s[i] >= '0' && s[i] <= '9' {
			num = int(s[i]-'0') + num*10
		}
	}
	if num >= 0 && num <= 255 {
		return true
	} else {
		return false
	}
}
