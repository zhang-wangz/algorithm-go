package solution

import "math"

func oneEditAway(first string, second string) bool {
	m, n := len(first), len(second)
	if math.Abs(float64(m-n)) > 1 {
		return false
	}
	i, j := 0, 0
	var flag bool
	for i < m && j < n {
		if first[i] == second[j] {
			i++
			j++
		} else if flag {
			return false
		} else {
			// m比n长，所以不一样的地方在m处
			if m > n {
				i++
			} else if m < n {
				j++
			} else {
				i++
				j++
			}
		}
	}
	return true
}
