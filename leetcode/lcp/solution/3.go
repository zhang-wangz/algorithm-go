package solution

import "strings"

// 计算x和y
func robot2(command string, obstacles [][]int, x int, y int) bool {
	if !isValidPath(command, x, y) {
		return false
	}
	for _, ob := range obstacles {
		if x+y > ob[0]+ob[1] && isValidPath(command, ob[0], ob[1]) {
			return false
		}
	}
	return true
}

// 可以抽象成x+y这么长的一条路，障碍只要在x+y里面，就是可能
func isValidPath(command string, x, y int) bool {
	y1 := strings.Count(command, "U")*((x+y)/len(command)) + strings.Count(command[0:(x+y)%len(command)], "U")
	x1 := strings.Count(command, "R")*((x+y)/len(command)) + strings.Count(command[0:(x+y)%len(command)], "R")
	if x == x1 && y == y1 {
		return true
	}
	return false
}

// 直接移动硬算模拟
func robot(command string, obstacles [][]int, x1 int, y1 int) bool {
	ob := map[int]map[int]int{}
	for i := 0; i < len(obstacles); i++ {
		i, j := obstacles[i][0], obstacles[i][1]
		if ob[i] == nil {
			ob[i] = map[int]int{}
		}
		ob[i][j] = 1
	}
	x, y := 0, 0
	i := 0
	for true {
		i = i % len(command)
		if command[i] == 'U' {
			y++
		} else if command[i] == 'R' {
			x++
		}

		if x == x1 && y == y1 {
			return true
		}
		if ob[x] != nil && ob[x][y] == 1 {
			return false
		}
		if x > x1 || y > y1 {
			return false
		}
		i++
	}
	return false
}
