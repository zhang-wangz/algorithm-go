package number

import (
	"math"
	"math/rand"
)

// https://leetcode.cn/problems/generate-random-point-in-a-circle/
// 随机生成圆内一个点
type Solution struct {
	radius float64
	x, y   float64
}

func Constructor1(radius float64, x_center float64, y_center float64) Solution {
	return Solution{radius: radius, x: x_center, y: y_center}
}

func (this *Solution) RandPoint() []float64 {
	// sina^2 + cosa^2 = 1
	for {
		x := rand.Float64()*2 - 1
		y := rand.Float64()*2 - 1 // -1 1
		if x*x+y*y <= 1 {
			return []float64{this.radius*x + this.x, this.radius*y + this.y}
		}
	}
}

func (s *Solution) RandPoint2() []float64 {
	// 为满足等概率
	s.radius = math.Sqrt(rand.Float64() * s.radius * s.radius)
	sin, cos := math.Sincos(rand.Float64() * 2 * math.Pi)
	return []float64{s.radius*cos + s.x, s.radius*sin + s.y}
}
