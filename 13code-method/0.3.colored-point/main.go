package main

import (
	"fmt"
	"image/color"
	"math"
)

type Point struct {
	X, Y float64
}

type ColoredPoint struct {
	*Point
	Color color.RGBA
}

func (p Point) Distance(q Point) float64 {
	return math.Hypot(p.X-q.X, p.Y-q.Y)
}

func (p *Point) scaleBy(fract float64) {
	p.X *= fract
	p.Y *= fract
} 
func main() {
	red := color.RGBA{255, 0, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}

	p := ColoredPoint{&Point{1, 2}, red}
	q := ColoredPoint{&Point{5, 4}, blue}

	fmt.Println(p.Distance(*q.Point))
	// p.Distance(*q) err	
	// invalid operation: cannot indirect q (variable of type ColoredPoint)compiler

	p.scaleBy(2)
	q.scaleBy(2)
	q.Point = p.Point
	fmt.Println(p.X, p.Y)
	fmt.Println(q.X, q.Y)

	fmt.Println(p.Distance(*q.Point))
}