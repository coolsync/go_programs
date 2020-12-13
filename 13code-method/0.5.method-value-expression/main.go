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
	q := Point{4, 6}

	// method value
	distanceFromP := p.Distance

	fmt.Println(distanceFromP(q))
	fmt.Printf("%T\n", distanceFromP)

	var origin Point	// {0, 0}

	fmt.Println(distanceFromP(origin))

	scaleByP := p.scaleBy

	scaleByP(2)
	scaleByP(4)
	scaleByP(8)
	fmt.Println(p)

	// method expression
	distance := Point.Distance

	fmt.Printf("%T\n", distance)

	scale := (*Point).scaleBy

	scale(&p, 2)

	fmt.Println(p)

	fmt.Printf("%T\n", scale)
}
