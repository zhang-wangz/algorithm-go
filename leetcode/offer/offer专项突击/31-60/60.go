package _1_60

// 桶排序
func topKFrequent(nums []int, k int) []int {
	mp := map[int]int{}
	for i := 0; i < len(nums); i++ {
		mp[nums[i]]++
	}
	t := map[int][]int{}
	for k, v := range mp {
		t[v] = append(t[v], k)
	}
	// 模拟桶
	cnt := 0
	ans := make([]int, 0)
out:
	for n := len(nums); n > 0; n-- {
		if _, ok := t[n]; ok {
			cnt += len(t[n])
			ans = append(ans, t[n]...)
			if cnt == k {
				break out
			}
		}
	}
	return ans
}
