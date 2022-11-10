package offer专项突击_1_30

func TwoSum(numbers []int, target int) []int {
	for i := 0; i < len(numbers); i++ {
		t := target - numbers[i]
		l, r := i+1, len(numbers)
		ans := -1
		for l < r {
			mid := l + (r-l)>>1
			if numbers[mid] > t {
				r = mid
			} else if numbers[mid] < t {
				l = mid + 1
			} else {
				ans = mid
				break
			}
		}
		if ans != -1 {
			return []int{i, ans}
		}
	}
	return []int{0, 0}
}
