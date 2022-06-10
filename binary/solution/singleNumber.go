package solution

// 给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
func singleNumber(nums []int) int {
	var res int
	for i := 0; i < len(nums); i++ {
		res = res ^ nums[i]
	}
	return res
}
