package main

import (
	"fmt"
	"math"
	"strings"
)

//empty interface
type empty interface {}

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

func (c circle) volume() float64 {
	return 4 / 3 * math.Pi * math.Pow(c.radius, 3)
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

func (r rectangle) getColor() string {
	return "black"
}

func printShapeDetails(s shape) {
	fmt.Println(strings.Repeat("#", 25))
	fmt.Printf("Shape : %T\n", s)
	fmt.Printf("Area : %v\n", s.area())
	fmt.Printf("Perimeter : %v\n", s.perimeter())
}

//embedded interfaces
type geometry interface {
	shape
	getColor() string
}

func printGeometryDetails(g geometry) {
	fmt.Println(strings.Repeat("#", 25))
	fmt.Printf("Shape : %T\n", g)
	fmt.Printf("Area : %v\n", g.area())
	fmt.Printf("Perimeter : %v\n", g.perimeter())
	fmt.Printf("Color : %v\n", g.getColor())
}

func main() {
	c := circle{radius: 2.}
	printShapeDetails(c)

	r := rectangle{height: 2.0, width: 3.0}
	printShapeDetails(r)

	fmt.Println(strings.Repeat("#", 25))

	var s shape = circle{radius: 2.5}

	//type assertion
	ball, ok := s.(circle)
	if ok {
		fmt.Println("Ball Volume : ", ball.volume())
	}

	//type switches
	switch value := s.(type) {
	case circle:
		fmt.Printf("Type of s is circle %#v", value)
	case rectangle:
		fmt.Printf("Type of s is rectangle %#v", value)
	}

	fmt.Println(strings.Repeat("#", 25))

	cube := rectangle{width: 3, height: 4}
	printGeometryDetails(cube)

}
