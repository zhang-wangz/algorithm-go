package main

const BitLen = 14

type TrieXor struct {
	child [2]*TrieXor
	cnt   int
}

func (t *TrieXor) insert(num int) {
	node := t
	for i := BitLen; i >= 0; i-- {
		k := (num >> i) & 1
		if node.child[k] == nil {
			node.child[k] = &TrieXor{}
		}
		node = node.child[k]
		node.cnt++
	}
}

func (t *TrieXor) getCnt(num int, x int) int {
	node := t
	res := 0
	for i := BitLen; i >= 0; i-- {
		k := (num >> i) & 1
		if node == nil {
			return res
		}
		if (x>>i)&1 == 1 {
			if node.child[k] != nil {
				res += node.child[k].cnt
			}
			// 因为需要xor 与 x保持一致，此时x的位是1，所以需要与1 xor
			node = node.child[k^1]
		} else {
			node = node.child[k]
		}
	}
	return res
}

func countPairs(nums []int, low int, high int) (ans int) {
	tr := &TrieXor{}
	for _, num := range nums {
		ans += tr.getCnt(num, high + 1) - tr.getCnt(num, low)
		tr.insert(num)
	}
	return ans
}
