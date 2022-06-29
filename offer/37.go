package offer

func asteroidCollision(asteroids []int) []int {
	stack := make([]int, 0)
	for i := 0; i < len(asteroids); i++ {
		if asteroids[i] > 0 {
			stack = append(stack, asteroids[i])
		} else {
			if len(stack) == 0 || stack[len(stack)-1] < 0 {
				stack = append(stack, asteroids[i])
				continue
			}
			for len(stack) != 0 && stack[len(stack)-1] > 0 {
				num := stack[len(stack)-1]
				// 如果撞过来的行星大
				if cmp(num, asteroids[i]) < 0 {
					stack = stack[:len(stack)-1]
					if len(stack) != 0 && stack[len(stack)-1] < 0 || len(stack) == 0 {
						stack = append(stack, asteroids[i])
					}
				} else if cmp(num, asteroids[i]) == 0 {
					stack = stack[:len(stack)-1]
					break
				} else {
					break
				}
			}
		}
	}
	return stack
}

func cmp(a, b int) int {
	if a < 0 {
		a = -a
	}
	if b < 0 {
		b = -b
	}
	return a - b
}
