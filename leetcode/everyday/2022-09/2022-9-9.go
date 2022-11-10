package main

// 1598  文件夹操作日志搜集器
func minOperations(logs []string) int {
	ans := 0
	for _, l := range logs {
		if l == "./" {
			continue
		}
		if l != "../" {
			ans++
		} else if ans > 0 {
			ans--
		}
	}
	return ans
}
