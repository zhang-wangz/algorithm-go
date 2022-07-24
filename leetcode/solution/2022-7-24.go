package solution

import "math"

// https://leetcode.cn/problems/distance-between-bus-stops/

func distanceBetweenBusStops(distance []int, start int, destination int) int {
	sum1, sum2 := 0, 0
	if start > destination {
		start, destination = destination, start
	}
	for i := range distance {
		if start <= i && i < destination {
			sum1 += distance[i]
		} else {
			sum2 += distance[i]
		}
	}
	return int(math.Min(float64(sum1), float64(sum2)))
}
