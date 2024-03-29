package _1_78

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
	for _, ch := range word {
		c := ch - 'a'
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

func (tr *Trie) remove(word string) {
	node := tr
	node.child[word[0]-'a'] = nil
}

// StartsWith /** Returns if there is any word in the trie that starts with the given prefix. */
func (tr *Trie) StartsWith(prefix string) bool {
	node := tr.searchPrefix(prefix)
	return node != nil
}
