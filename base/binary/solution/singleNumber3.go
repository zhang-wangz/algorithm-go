package solution

// 给定一个整数数组 nums，其中恰好有两个元素只出现一次，其余所有元素均出现两次。 找出只出现一次的那两个元素。
// [2,3,4,5,4,5]
func singleNumber3(nums []int) []int {
	var res int
	for i := 0; i < len(nums); i++ {
		res = res ^ nums[i]
	}
	// res = a ^ b
	// 找到最接近找的位置
	r := []int{res, res}
	// 根据最后一位1的位置进行进行区分
	// 获取最后一个1
	diff := (res & (res - 1)) ^ res
	for i := 0; i < len(nums); i++ {
		if diff&nums[i] == 0 {
			r[0] = r[0] ^ nums[i]
		} else {
			r[1] = r[1] ^ nums[i]
		}
	}
	return r
}
