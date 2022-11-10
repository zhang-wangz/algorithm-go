package _1_78

import (
	"math"
	"strings"
)

// 将前缀插入字典数，然后根据每个词搜索字典数有没有符合的前缀，找到一个isEnd的就返回true
// 然后replace同理，在isEnd前面并且存在的就添加进stringBuilder
func replaceWords(dictionary []string, sentence string) string {
	tr := Trie{child: [26]*Trie{}, isEnd: false}
	strs := strings.Split(sentence, " ")
	for i := 0; i < len(strs); i++ {
		word := strs[i]
		tr.Insert(word)
		ml := math.MaxInt
		for _, prefix := range dictionary {
			if len(prefix) < ml && tr.StartsWith(prefix) {
				ml = len(prefix)
				strs[i] = prefix
			}
		}
		tr.remove(word)
	}
	return strings.Join(strs, " ")
}
