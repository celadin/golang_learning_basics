package main

import (
	"fmt"
	"strings"
)

func main() {

	var cities []string

	fmt.Printf("cities : %#v\n", cities)

	if cities == nil {
		fmt.Println("Cities is nil")
	}

	var numbers = []int{1, 2, 3, 4}
	fmt.Printf("Numbers : %#v\n", numbers)

	//make built-in function
	var nums = make([]int, 3)
	fmt.Printf("Nums : %#v\n", nums)

	//type
	type names []string
	friends := names{"Sonya", "Maria"}
	fmt.Printf("Friends : %#v\n", friends)

	//comparing
	a, b := []int{1, 2, 3}, []int{1, 2, 3, 4}
	//fmt.Println("a == b", a == b) // error

	var eq bool = true

	if len(a) == len(b) {

		for i, valueA := range a {
			if valueA != b[i] {
				eq = false
				break
			}
		}

	} else {
		eq = false
	}

	if eq {
		fmt.Println("a and b slices are equal")
	} else {
		fmt.Println("a and b slices are not equal")
	}

	var x = []int{1, 2, 3}
	x = append(x, 10, 20, 30)
	fmt.Println("x :", x)

	y := []int{100, 200}
	x = append(x, y...)
	fmt.Println("x :", x)

	//copy
	src := []int{10, 20, 30}
	dst := make([]int, len(src))
	nn := copy(dst, src)
	fmt.Printf("src : %v, dst : %v, nn : %v\n", src, dst, nn)

	fmt.Println(strings.Repeat("#", 30))

	//slicing
	s := []int{1, 2, 3, 4, 5}
	t := s[1:3]
	fmt.Printf("sliced t : %#v\n", t)

	fmt.Println("\ns 	=>", s)
	fmt.Println("s[:2]  =>", s[:2])
	fmt.Println("s[3:]  =>", s[3:])
	fmt.Println("s[2:4] =>", s[2:4])
	fmt.Println("append(s[:3], 100) =>", append(s[:3], 100))

}
