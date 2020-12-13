package geometry

import "math"

// Point 创建 二维图 一个点 的结构体
type Point struct {
	X, Y float64
}

// Hypot returns Sqrt(p*p + q*q), taking care to avoid unnecessary overflow and underflow
// Hypot返回Sqrt（p *p + q *q），请注意避免不必要的上溢和下溢

// 两个点 距离的方法
func (p Point) Distence(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// 定量扩大坐标
func (p *Point) Scale(factor float64) {
	p.X *= factor
	p.Y *= factor
}

// count 多个点依次连接的距离
type Path []Point

func (path Path) Distence() float64 {
	sum := 0.0

	for i := range path {
		if i > 0 {
			sum += path[i-1].Distence(path[i])
		}
	}

	return sum
}
