package _022_07

type trie2 struct {
	child [26]*trie2
	isEnd bool
}

func (tr *trie2) insert(word string) {
	node := tr
	for i := 0; i < len(word); i++ {
		ch := word[i] - 'a'
		if node.child[ch] == nil {
			node.child[ch] = &trie2{}
		}
		node = node.child[ch]
	}
	node.isEnd = true
}

func (tr *trie2) search(word string) bool {
	node := tr
	for i := 0; i < len(word); i++ {
		ch := word[i] - 'a'
		if node.child[ch] == nil {
			return false
		}
		node = node.child[ch]
	}
	return node.isEnd
}

type MagicDictionary struct {
	tr *trie2
}

func Constructor1() MagicDictionary {
	return MagicDictionary{tr: &trie2{child: [26]*trie2{}, isEnd: false}}
}

func (this *MagicDictionary) BuildDict(dictionary []string) {
	for _, dic := range dictionary {
		this.tr.insert(dic)
	}
}

// Search 是可以存在的，比如，hello，hallo都在dic中，hello转化一个字母e为a即可，
// 不需要判断node.child[ch]为nil才进行转换
func (this *MagicDictionary) Search(searchWord string) bool {
	node := this.tr
	for i := 0; i < len(searchWord); i++ {
		ch := searchWord[i] - 'a'
		for x := 0; x < 26; x++ {
			c := byte(x)
			if c == ch || node.child[c] == nil {
				continue
			}
			if node.child[c].search(searchWord[i+1:]) {
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
