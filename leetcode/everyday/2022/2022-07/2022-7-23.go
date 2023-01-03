package _022_07

func sequenceReconstruction(nums []int, sequences [][]int) bool {
	graph := map[int]map[int]bool{}
	in := make([]int, len(nums))
	for _, v := range sequences {
		for i := 0; i < len(v)-1; i++ {
			i, j := v[i]-1, v[i+1]-1
			if graph[i] == nil {
				graph[i] = map[int]bool{}
			}
			if !graph[i][j] {
				in[j]++
			}
			graph[i][j] = true
		}
	}
	zeroIn := make([]int, 0)
	for i := range in {
		if in[i] == 0 {
			zeroIn = append(zeroIn, i)
		}
	}
	if len(zeroIn) > 1 {
		return false
	}

	for len(zeroIn) != 0 {
		zero := zeroIn[0]
		zeroIn = zeroIn[1:]
		for i := range graph[zero] {
			in[i]--
			if in[i] == 0 {
				zeroIn = append(zeroIn, i)
			}
		}
		if len(zeroIn) > 1 {
			return false
		}
	}
	return true
}
