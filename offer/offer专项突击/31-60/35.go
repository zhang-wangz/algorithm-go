package _1_60

import (
	"math"
	"sort"
	"strconv"
)

func findMinDifference(timePoints []string) int {
	sort.Strings(timePoints)
	min := math.MaxInt
	for i := 0; i < len(timePoints)-1; i++ {
		time1 := tansferTime(timePoints[i])
		time2 := tansferTime(timePoints[i+1])
		min = int(math.Min(float64(min), float64(timeSub(time2, time1))))
	}
	min = int(math.Min(float64(min),
		float64(timeSub(tansferTime(timePoints[0]), tansferTime(timePoints[len(timePoints)-1])))))
	return min
}

func tansferTime(date string) int {
	h := date[:2]
	m := date[3:]
	hn, _ := strconv.Atoi(h)
	mn, _ := strconv.Atoi(m)
	return (hn*60 + mn) % (24 * 60)
}
func timeSub(d int, d1 int) int {
	return (d - d1 + 24*60) % (24 * 60)
}
