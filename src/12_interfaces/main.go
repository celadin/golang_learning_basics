package main

import (
	"fmt"
	"math"
	"strings"
)

type shape interface {
	area() float64
	perimeter() float64
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

type rectangle struct {
	width  float64
	height float64
}

func (r rectangle) area() float64 {
	return r.height * r.width
}

func (r rectangle) perimeter() float64 {
	return 2 * (r.height + r.width)
}

func printShapeDetails(s shape) {
	fmt.Println(strings.Repeat("#", 25))
	fmt.Printf("Shape : %T\n", s)
	fmt.Printf("Area : %v\n", s.area())
	fmt.Printf("Perimeter : %v\n", s.perimeter())
}

func main() {
	c := circle{radius: 2.}
	printShapeDetails(c)

	r := rectangle{height: 2.0, width: 3.0}
	printShapeDetails(r)
}
