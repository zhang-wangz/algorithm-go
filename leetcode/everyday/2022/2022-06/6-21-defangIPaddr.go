package _022_06

import "strings"

// https://leetcode.cn/problems/defanging-an-ip-address/
// 给你一个有效的 IPv4 地址 address，返回这个 IP 地址的无效化版本。
//
// 所谓无效化 IP 地址，其实就是用 "[.]" 代替了每个 "."。
func defangIPaddr(address string) string {
	res := strings.Replace(address, ".", "[.]", len(address))
	return res
}
