package solution

import "strings"

type trie struct {
	child [26]*trie
	isEnd bool
}

func (tr *trie) insert(word string) {
	node := tr
	for i := 0; i < len(word); i++ {
		ch := word[i] - 'a'
		if node.child[ch] != nil {
			node.child[ch] = &trie{}
		}
		node = node.child[ch]
	}
	node.isEnd = true
}

func (tr *trie) search(word string) (f bool, ans string) {
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
	t := &trie{}
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
