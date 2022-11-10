package main

// 枚举
func flipLights(n int, presses int) int {
	if n == 0 {
		return 1
	}
	vis := map[int]struct{}{}
	preArr := make([]int, 4)
	for i := 0; i < 1<<4; i++ {
		sum := 0
		for j := 0; j < 4; j++ {
			preArr[j] = (i >> j) & 1
			sum += preArr[j]
		}
		if sum%2 == presses%2 && sum <= presses {
			status := preArr[0] ^ preArr[2] ^ preArr[3]
			if n >= 2 {
				status |= (preArr[0] ^ preArr[1]) << 1
			}
			if n >= 3 {
				status |= (preArr[0] ^ preArr[2]) << 2
			}
			vis[status] = struct{}{}
		}
	}
	return len(vis)
}
