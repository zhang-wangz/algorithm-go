package _1_78

func peakIndexInMountainArray(arr []int) int {
	l, r := 0, len(arr)-1
	var ans int
	for l < r {
		mid := l + (r-l)<<1
		if arr[mid] > arr[mid+1] {
			ans = mid
			r = mid
		} else {
			l = mid + 1
		}
	}
	return ans
}
