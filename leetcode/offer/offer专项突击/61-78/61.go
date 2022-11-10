package _1_78

import "container/heap"

// 实现go的优先队列接口
type pair struct {
	i, j int
}
type hp struct {
	nums1, nums2 []int
	data         []pair
}

func (h *hp) Len() int {
	return len(h.data)
}
func (h *hp) Less(i, j int) bool {
	a, b := h.data[i], h.data[j]
	return h.nums1[a.i]+h.nums2[a.j] < h.nums1[b.i]+h.nums2[b.j]
}
func (h *hp) Swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}
func (h *hp) Push(val any) {
	h.data = append(h.data, val.(pair))
}
func (h *hp) Pop() any {
	v := h.data[len(h.data)-1]
	h.data = h.data[:len(h.data)-1]
	return v
}

func kSmallestPairs(nums1 []int, nums2 []int, k int) (ans [][]int) {
	m, n := len(nums1), len(nums2)
	h := &hp{nums1, nums2, nil}
	// 相当于i<min(m,k)
	for i := 0; i < m && i < k; i++ {
		heap.Push(h, pair{i, 0})
	}
	for h.Len() > 0 && len(ans) < k {
		t := heap.Pop(h).(pair)
		i, j := t.i, t.j
		ans = append(ans, []int{nums1[i], nums2[j]})
		if j+1 < n {
			heap.Push(h, pair{i, j + 1})
		}
	}
	return ans
}

// 堆排 // // 全部放进去的堆排
type node struct {
	key [2]int
	val int
}

func kSmallestPairs2(nums1 []int, nums2 []int, k int) [][]int {
	nums := make([]*node, 0)
	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			t := &node{
				key: [2]int{nums1[i], nums2[j]},
				val: nums1[i] + nums2[j],
			}
			nums = append(nums, t)
		}
	}
	ans := make([][]int, 0)
	if len(nums) < k {
		k = len(nums)
	}
	// 根据mp的val做一个堆，最后的数是小值
	sink(nums, len(nums))
	n := len(nums) - 1
	for i := n; i >= n-k+1; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		sink(nums, i)
	}
	for i := n; i >= n-k+1; i-- {
		ans = append(ans, nums[i].key[:])
	}
	return ans
}

func sink(nums []*node, len int) {
	for i := len/2 - 1; i >= 0; i-- {
		left := 2*i + 1
		if left < len && nums[left].val < nums[i].val {
			nums[i], nums[left] = nums[left], nums[i]
		}
		right := 2*i + 2
		if right < len && nums[right].val < nums[i].val {
			nums[i], nums[right] = nums[right], nums[i]
		}
	}
}
