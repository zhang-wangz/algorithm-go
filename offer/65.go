package main

import (
	"sort"
	"strings"
)

type trie1 struct {
	child [26]*trie1
}

func (trie *trie1) insert(word string) {
	node := trie
	for i := len(word) - 1; i >= 0; i-- {
		ch := word[i]
		ch -= 'a'
		if node.child[ch] == nil {
			node.child[ch] = &trie1{}
		}
		node = node.child[ch]
	}
}
func (trie *trie1) search(word string) bool {
	node := trie
	for i := len(word) - 1; i >= 0; i-- {
		ch := word[i]
		ch -= 'a'
		if node.child[ch] == nil {
			return false
		}
		node = node.child[ch]
	}
	return true
}

func minimumLengthEncoding(words []string) int {
	tr := &trie1{}
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) > len(words[j])
	})
	sb := strings.Builder{}
	for _, w := range words {
		if !tr.search(w) {
			tr.insert(w)
			sb.WriteString(w)
			sb.WriteByte('#')
		}
	}
	return sb.Len()
}
