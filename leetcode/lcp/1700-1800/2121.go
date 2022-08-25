package main

// 2121
func getDistances(arr []int) (ans []int64) {
	pos := map[int][]int{}
	for i, v := range arr {
		pos[v] = append(pos[v], i)
	}
	ans = make([]int64, len(arr))
	for _, p := range pos {
		sum := int64(0)
		for i := range p {
			sum += int64(p[i] - p[0])
		}
		ans[p[0]] = sum
		for i := 1; i < len(p); i++ {
			sum += int64(2*i-len(p)) * int64(p[i]-p[i-1])
			ans[p[i]] = sum
		}
	}
	return
}
