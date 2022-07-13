package solution

// https://leetcode.cn/problems/asteroid-collision/
// 栈模拟
func asteroidCollision(asteroids []int) []int {
	stack := make([]int, 0)
	for _, aster := range asteroids {
		alive := true
		for alive && aster < 0 && len(stack) > 0 && stack[len(stack)-1] > 0 {
			alive = -aster > stack[len(stack)-1]
			if -aster >= stack[len(stack)-1] {
				stack = stack[:len(stack)-1]
			}
		}
		if alive {
			stack = append(stack, aster)
		}
	}
	return stack
}
