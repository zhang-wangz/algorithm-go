package _022_08

func busyStudent(startTime []int, endTime []int, queryTime int) int {
	cnt := 0
	for i := range startTime {
		if startTime[i] <= queryTime && queryTime <= endTime[i] {
			cnt++
		}
	}
	return cnt
}
