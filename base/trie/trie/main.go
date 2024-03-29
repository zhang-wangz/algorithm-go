package main

// 字典树
type Trie struct {
	child [26]*Trie
	isEnd bool
}

// Constructor10 /** Initialize your data structure here. */
func Constructor10() Trie {
	return Trie{child: [26]*Trie{}, isEnd: false}
}

// Insert /** Inserts a word into the trie. */
func (tr *Trie) Insert(word string) {
	node := tr
	for i := 0; i < len(word); i++ {
		c := word[i] - 'a'
		if node.child[c] == nil {
			node.child[c] = &Trie{}
		}
		node = node.child[c]
	}
	node.isEnd = true
}
func (tr *Trie) searchPrefix(prefix string) *Trie {
	node := tr
	for _, ch := range prefix {
		ch -= 'a'
		if node.child[ch] == nil {
			return nil
		}
		node = node.child[ch]
	}
	return node
}

// Search /** Returns if the word is in the trie. */
func (tr *Trie) Search(word string) bool {
	node := tr.searchPrefix(word)
	return node != nil && node.isEnd
}
