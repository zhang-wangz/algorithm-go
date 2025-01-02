package solution

// 给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现了三次。找出那个只出现了一次的元素。
// 计算每一位上的1，然后计算%3之后的1就是只出现1次的元素的位值
func singleNumber2(nums []int) int {
	var res int
	for i := 0; i < 64; i++ {
		sum := 0
		for j := 0; j < len(nums); j++ {
			// 左移动之后最后一位是1
			sum += (nums[j] >> i) & 1
		}
		res |= (sum % 3) << i
	}
	return res
}
