package main

type MapSum struct {
	*tr
	mp map[string]int
}
type tr struct {
	child [26]*tr
	v     int
}

// Constructor /** Initialize your data structure here. */
func Constructor() MapSum {
	return MapSum{
		&tr{
			child: [26]*tr{},
		},
		map[string]int{},
	}
}

func (msum *MapSum) Insert(key string, val int) {
	node := msum.tr
	for _, ch := range key {
		ch -= 'a'
		if node.child[ch] == nil {
			node.child[ch] = &tr{v: val}
		} else {
			node.child[ch].v = node.child[ch].v - msum.mp[key] + val
		}
		node = node.child[ch]
	}
	msum.mp[key] = val
}

func (msum *MapSum) Sum(prefix string) int {
	node := msum.tr
	for _, ch := range prefix {
		ch -= 'a'
		if node.child[ch] == nil {
			return 0
		}
		node = node.child[ch]
	}
	return node.v
}
