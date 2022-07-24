package main

func zeroFilledSubarray(nums []int) int64 {
	cnt := 0
	cnt0 := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			cnt0++
		} else {
			cnt += (1 + cnt0) * cnt0 / 2
			cnt0 = 0
		}
	}
	if cnt0 != 0 {
		cnt += (1 + cnt0) * cnt0 / 2
	}
	return int64(cnt)
}
