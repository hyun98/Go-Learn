package main

import (
	"fmt"
	"math"
)

type Figure interface {
	surfaceArea() float64
	volume() float64
}

type Cylinder struct {
	radius float64
	height float64
}


type Rect struct {
	side [3]float64
}

func (c Cylinder) surfaceArea() float64 {
	return 2 * math.Pi * c.radius * (c.radius + c.height)
}

func (c Cylinder) volume() float64 {
	return math.Pi * c.radius * c.radius * c.height
}

func (r Rect) surfaceArea() float64 {
	return 2 * (r.side[0] * r.side[1] + r.side[0] * r.side[2] + r.side[1] * r.side[2])
}

func (r Rect) volume() float64 {
	return r.side[0] * r.side[1] * r.side[2]
}

func main() {
	cy1 := Cylinder{radius: 10, height: 10}
	cy2 := Cylinder{radius: 4.2, height: 15.6}
	cu1 := Rect{side: [3]float64{10.5, 20.2, 20}}
	cu2 := Rect{side: [3]float64{4, 10, 23}}

	printMeasure(cy1, cy2, cu1, cu2)
}

func printMeasure(figures ...Figure) {
	for _, f := range figures {
		fmt.Printf("%.2f, %.2f\n", f.surfaceArea(), f.volume())
	}
}