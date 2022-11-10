package _1_78

type trie struct {
	child [26]*trie
	isEnd bool
}
type MagicDictionary struct {
	*trie
}

// Constructor /** Initialize your data structure here. */
func Constructor11() MagicDictionary {
	return MagicDictionary{&trie{child: [26]*trie{}, isEnd: false}}
}

func (dic *MagicDictionary) BuildDict(dictionary []string) {
	for _, word := range dictionary {
		dic.insert(word)
	}
}

func (tr *trie) insert(word string) {
	node := tr
	for _, ch := range word {
		ch -= 'a'
		if node.child[ch] == nil {
			node.child[ch] = &trie{}
		}
		node = node.child[ch]
	}
	node.isEnd = true
}
func (tr *trie) contains(word string) bool {
	node := tr
	for _, ch := range word {
		ch -= 'a'
		if node.child[ch] == nil {
			return false
		}
		node = node.child[ch]
	}
	return node.isEnd
}

func (dic *MagicDictionary) Search(searchWord string) bool {
	node := dic.trie
	for i := 0; i < len(searchWord); i++ {
		ch := int(searchWord[i] - 'a')
		// 替换
		for x := 0; x < 26; x++ {
			if x == ch || node.child[x] == nil {
				continue
			}
			if node.child[x].contains(searchWord[i+1:]) {
				return true
			}
		}
		if node.child[ch] == nil {
			return false
		}
		node = node.child[ch]
	}
	return false
}
