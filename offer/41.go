package offer

type MovingAverage struct {
	right int
	size  int
	nums  []int
	sum   int
	le    int
}

/** Initialize your data structure here. */
func Constructor4(size int) MovingAverage {
	return MovingAverage{right: 0, size: size, nums: make([]int, 0)}
}

func (this *MovingAverage) Next(val int) float64 {
	if this.right < this.size {
		this.nums = append(this.nums, val)
		this.le++
		this.right++
		this.sum += val
	} else {
		this.nums = append(this.nums, val)
		this.sum -= this.nums[0]
		this.nums = this.nums[1:]
	}
	return float64(this.sum) / float64(this.le)
}
