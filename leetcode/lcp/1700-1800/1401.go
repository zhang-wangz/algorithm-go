package main

// 1401
func checkOverlap(radius int, xCenter int, yCenter int, x1 int, y1 int, x2 int, y2 int) bool {
	// 求最近的一个点有没有被包含
	x, y := -1, -1
	if x1 > xCenter {
		x = x1 - xCenter
	} else {
		if x2 < xCenter {
			x = xCenter - x2
		} else {
			x = 0
		}
	}
	if y1 > yCenter {
		y = y1 - yCenter
	} else {
		if y2 < yCenter {
			y = yCenter - y2
		} else {
			y = 0
		}
	}
	return x*x+y*y <= radius*radius
}
