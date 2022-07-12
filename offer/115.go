package main

func sequenceReconstruction(nums []int, sequences [][]int) bool {
	zeroIn := make([]int, 0)
	// int->set
	graph := make(map[int]map[int]bool, 0)
	in := make([]int, len(nums))
	for _, v := range sequences {
		for i := 0; i < len(v)-1; i++ {
			w1 := v[i] - 1
			w2 := v[i+1] - 1
			if graph[w1] == nil {
				graph[w1] = map[int]bool{}
			}
			if !graph[w1][w2] {
				in[w2]++
			}
			graph[w1][w2] = true
		}
	}
	for i, z := range in {
		if z == 0 {
			zeroIn = append(zeroIn, i)
		}
	}
	if len(zeroIn) > 1 {
		return false
	}
	for len(zeroIn) != 0 {
		node := zeroIn[0]
		zeroIn = zeroIn[1:]
		// nei
		for k, _ := range graph[node] {
			in[k]--
			if in[k] == 0 {
				zeroIn = append(zeroIn, k)
			}
		}
		if len(zeroIn) > 1 {
			return false
		}
	}
	return true
}
