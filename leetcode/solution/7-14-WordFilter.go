package solution

type trie struct {
	child     [26]*trie
	isEnd     bool
	isSuffEnd bool
}

func (tr *trie) insert(word string) {
	node := tr
	for i := 0; i < len(word); i++ {
		ch := word[i] - 'a'
		if node.child[ch] == nil {
			node.child[ch] = &trie{}
		}
		node = node.child[ch]
	}
	node.isEnd = true
}

func (tr *trie) insertSuff(word string) {
	node := tr
	for i := len(word) - 1; i >= 0; i-- {
		ch := word[i] - 'a'
		if node.child[ch] == nil {
			node.child[ch] = &trie{}
		}
		node = node.child[ch]
	}
	node.isSuffEnd = true
}
func (tr *trie) removeSuff(word string) {
	node := tr
	node.child[word[len(word)-1]-'a'] = nil
}
func (tr *trie) remove(word string) {
	node := tr
	node.child[word[0]-'a'] = nil
}

func (tr *trie) searchPref(word string) bool {
	node := tr
	for i := 0; i < len(word); i++ {
		ch := word[i] - 'a'
		if node.child[ch] == nil {
			return false
		}
		node = node.child[ch]
		if node.isEnd {
			return true
		}
	}
	return false
}
func (tr *trie) searchSuff(word string) bool {
	node := tr
	for i := len(word) - 1; i >= 0; i-- {
		ch := word[i] - 'a'
		if node.child[ch] == nil {
			return false
		}
		node = node.child[ch]
		if node.isSuffEnd {
			return true
		}
	}
	return false
}

type WordFilter struct {
	*trie
	word []string
}

func Constructor3(words []string) WordFilter {
	return WordFilter{&trie{}, words}
}

func (wf *WordFilter) F(pref string, suff string) int {
	tr := wf.trie
	tr.insert(pref)
	tr.insertSuff(suff)
	for i := len(wf.word) - 1; i >= 0; i-- {
		w := wf.word[i]
		if tr.searchPref(w) && tr.searchSuff(w) {
			tr.remove(pref)
			tr.removeSuff(suff)
			return i
		}
	}
	tr.remove(pref)
	tr.removeSuff(suff)
	return -1
}
