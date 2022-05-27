package solution

// https://leetcode.cn/problems/find-closest-lcci/
// 有个内含单词的超大文本文件，给定任意两个不同的单词，找出在这个文件中这两个单词的最短距离(相隔单词数)。
// 如果寻找过程在这个文件中会重复多次，而每次寻找的单词不同，你能对此优化吗?

// 如果重复多次，维护words单词中的下标列表，然后双指针进行比较
func findClosest(words []string, word1 string, word2 string) int {
	i, j := -1, -1
	minD := 1 << 31
	for idx, v := range words {
		if v == word1 {
			i = idx
		} else if v == word2 {
			j = idx
		}
		if i != -1 && j != -1 {
			minD = min(minD, abs(i, j))
		}
	}
	return minD
}
func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
