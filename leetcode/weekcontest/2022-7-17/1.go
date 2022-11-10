package _022_7_17

func numberOfPairs(nums []int) []int {
	mp := map[int]int{}
	for i := range nums {
		mp[nums[i]]++
	}
	cnt := 0
	for k, _ := range mp {
		cnt += mp[k] / 2
	}
	return []int{cnt, len(nums) - cnt*2}
}
