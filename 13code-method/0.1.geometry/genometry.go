package main
import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

type Path []Point

func Distance(p, q Point) float64 {
	return math.Hypot((p.X - q.X), (p.Y - q.Y))
}

func (p Point) Distance(q Point) float64 {
	return math.Hypot((p.X - q.X), (p.Y - q.Y))
}

func (p *Point) scaleBy(fract float64) {
	p.X *= fract
	p.Y *= fract
}

func (path Path) Distance() float64{
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}
	}
	return sum
}

func main() {

	p := Point{1, 2}
	p.scaleBy(5)

	q := Point{3, 4}
	q.scaleBy(5)
	fmt.Println(Distance(p, q))

	fmt.Println(p.Distance(q))

	// 演示一个三角形的周长
	perim := Path {
		{1, 1},
		{5, 2},
		{5, 3},
		{1, 1},
	}
	for i := range perim {
		perim[i].scaleBy(5)
	}
	fmt.Println("perim: ", perim.Distance())
}
