package _1_78

import "math"

// 非常规前缀数问题
type trie2 struct {
	child [2]*trie2
}

func (tr *trie2) insert(nums []int) {
	for _, num := range nums {
		node := tr
		for i := 30; i >= 0; i-- {
			ch := num >> i & 1
			if node.child[ch] == nil {
				node.child[ch] = &trie2{}
			}
			node = node.child[ch]
		}
	}
}
func (tr *trie2) searchMaxXOR(num int) int {
	node := tr
	x := 0
	for i := 30; i >= 0; i-- {
		// 如果是1的话走0，如果是0的话走1
		// ch为0，tch为1，ch为1，tch为0
		ch := num >> i & 1
		tch := ch ^ 1
		if node.child[tch] == nil {
			node = node.child[ch]
			x = x*2 + ch
		} else {
			node = node.child[tch]
			x = x*2 + tch
		}
	}
	return x ^ num
}
func findMaximumXOR(nums []int) int {
	tr := &trie2{child: [2]*trie2{}}
	tr.insert(nums)
	ans := math.MinInt
	for i := 0; i < len(nums); i++ {
		ans = int(math.Max(float64(ans), float64(tr.searchMaxXOR(nums[i]))))
	}
	return ans
}

// hashmap,其实是set
// 异或互换性质，aj = ai ^ x, x = ai ^ aj
func findMaximumXOR2(nums []int) (x int) {
	for i := 30; i >= 0; i-- {
		mp := map[int]bool{}
		for _, num := range nums {
			mp[num>>i] = true
		}

		xNext := x*2 + 1
		found := false
		for _, num := range nums {
			if mp[num>>i^xNext] {
				found = true
			}
		}
		if !found {
			x = xNext - 1
		} else {
			x = xNext
		}
	}
	return x
}
