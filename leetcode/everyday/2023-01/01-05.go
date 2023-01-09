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

func countPairs2(nums []int, low int, high int) (ans int) {
	tr := &TrieXor{}
	for _, num := range nums {
		ans += tr.getCnt(num, high+1) - tr.getCnt(num, low)
		tr.insert(num)
	}
	return ans
}

func countPairs(nums []int, low int, high int) (ans int) {
	m := map[int]int{}
	for _, num := range nums {
		m[num]++
	}
	for high++; high > 0; high >>= 1 {
		next := map[int]int{}
		for k, v := range m {
			if high&1 == 1 {
				// x ^ y = t => x ^ t = y
				ans += v * m[k^(high-1)]
			}
			if low&1 == 1 {
				ans -= v * m[k^(low-1)]
			}
			next[k>>1] += v
		}
		m = next
		low >>= 1
	}
	return ans / 2
}
