package main

import "fmt"

type Point struct {
	X int
	Y int
}

type Circle struct {
	Point
	Radius int
}

type Wheel struct {
	Circle
	Spokes int		// 轮辐
}

func main(){
	wheel := Wheel{Circle{Point{8, 8}, 5}, 20}

	fmt.Printf("%#v\n", wheel)

	wheel = Wheel{
		Circle: Circle{
			Point: Point{X: 5, Y: 5},
			Radius: 8,
		},
		Spokes: 30,
	}

	wheel.X = 42
	fmt.Printf("X: %d\n", wheel.Circle.Point.X)


	fmt.Printf("X: %d\n", wheel.X)

	fmt.Printf("%#v\n", wheel)

}