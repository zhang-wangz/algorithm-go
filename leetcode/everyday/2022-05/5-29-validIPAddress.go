package _022_05

import (
	"strings"
)

// https://leetcode.cn/problems/validate-ip-address/
// 验证ip4或者ip6

func ValidIPAddress(queryIP string) string {
	str := strings.Split(queryIP, ".")
	f := false
	if len(str) == 1 {
		str = strings.Split(queryIP, ":")
		f = true
	}
	// ipv4
	if !f {
		if len(str) != 4 {
			return "Neither"
		}
		for _, item := range str {
			num := 0
			if item == "" {
				return "Neither"
			}
			for i := 0; i < len(item); i++ {
				// 防止首位为0
				if i == 0 && item[i] == '0' && len(item) != 1 {
					return "Neither"
				}
				if item[i] >= '0' && item[i] <= '9' {
					num = num*10 + int(item[i]-'0')
				} else {
					return "Neither"
				}
			}
			if num < 0 || num > 255 {
				return "Neither"
			}
		}
		return "IPv4"
	} else {
		if len(str) != 8 {
			return "Neither"
		}
		for _, item := range str {
			if len(item) > 4 || len(item) < 1 {
				return "Neither"
			}
			for i := 0; i < len(item); i++ {
				c := item[i]
				if (c >= '0' && c <= '9') || (c >= 'a' && c <= 'f') || (c >= 'A' && c <= 'F') {
					continue
				} else {
					return "Neither"
				}
			}
		}
		return "IPv6"
	}
}
