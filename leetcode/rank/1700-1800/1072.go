package main

func maxEqualRowsAfterFlips(matrix [][]int) (ans int) {
	// m, n := len(matrix), len(matrix[0])
	cnt := map[string]int{}
	for i := range matrix {
		sb := []byte{}
		for j := range matrix[i] {
			sb = append(sb, byte(matrix[i][j]+'0'))
		}
		if sb[0] != '0' {
			for j := 0; j < len(sb); j++ {
				sb[j] = ('1' - sb[j]) + '0'
			}
		}
		cnt[string(sb)]++
	}
	for _, v := range cnt {
		if v > ans {
			ans = v
		}
	}
	return
}
