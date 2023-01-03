package _022_07

import "strings"

type trie1 struct {
	child [26]*trie1
	isEnd bool
}

func (tr *trie1) insert(word string) {
	node := tr
	for i := 0; i < len(word); i++ {
		ch := word[i] - 'a'
		if node.child[ch] == nil {
			node.child[ch] = &trie1{}
		}
		node = node.child[ch]
	}
	node.isEnd = true
}

func (tr *trie1) search(word string) (f bool, ans string) {
	node := tr

	for i := 0; i < len(word); i++ {
		ch := word[i] - 'a'
		if node.child[ch] == nil {
			return false, ""
		}
		node = node.child[ch]
		ans += string(ch + 'a')
		if node.isEnd {
			return true, ans
		}
	}
	return false, ""
}

func replaceWords(dictionary []string, sentence string) string {
	t := &trie1{}
	for _, w := range dictionary {
		t.insert(w)
	}
	sen := strings.Split(sentence, " ")
	for i, se := range sen {
		if ok, ans := t.search(se); ok {
			sen[i] = ans
		}
	}
	return strings.Join(sen, " ")
}
