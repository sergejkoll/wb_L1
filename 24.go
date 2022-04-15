package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func newPoint(x, y float64) *Point {
	return &Point{x: x, y: y}
}

func (p Point) Distance(secondPoint *Point) float64 {
	return math.Sqrt(math.Pow(p.x-secondPoint.x, 2) + math.Pow(p.y-secondPoint.y, 2))
}

func main() {
	p1 := newPoint(0, 0)
	p2 := newPoint(3, 4)
	dist := p1.Distance(p2)
	fmt.Println(dist)
}
