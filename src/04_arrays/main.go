package main

import (
	"fmt"
	"strings"
)

func main() {
	//declaring arrays
	var numbers [4]int

	fmt.Printf("%v\n", numbers)
	fmt.Printf("%#v\n", numbers)

	var a1 = [4]float64{}
	fmt.Printf("%#v\n", a1)

	var a2 = [3]int{-1, 5, -100}
	fmt.Printf("%#v\n", a2)

	var a3 = [4]string{"A", "B"}
	fmt.Printf("%#v\n", a3)

	// ... -> ellipsis operator -- length değerini otomatik atar
	var a4 = [...]int{1, 2, 5, 78, 9, -2}
	fmt.Printf("%#v\n", a4)

	fmt.Println(strings.Repeat("#", 40))

	for i, v := range a4 {
		fmt.Println("Index:", i, "Value:", v)
	}

	fmt.Println(strings.Repeat("#", 40))

	for i := 0; i < len(a4); i++ {
		fmt.Println("Index:", i, "Value:", a4[i])
	}

	fmt.Println(strings.Repeat("#", 40))

	//multi-dimensional arrays
	var mda = [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println("Multi-Dimensional Array :", mda)

	fmt.Println(strings.Repeat("#", 40))

	//array equality
	m := [3]int{1,2,3}
	n := m
	fmt.Println("n is equal to m : ", n == m)
	n[0] = -1
	fmt.Println("after changing -> n is equal to m : ", n == m)

	fmt.Println(strings.Repeat("#", 40))

	//keyed elements
	var k = [...]int{6:9}
	fmt.Printf("%#v\n", k)

}
