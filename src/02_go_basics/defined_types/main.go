package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World!")

	type age int
	type oldAge age

	var child age = 10
	var adult oldAge = 80

	var x int = int(child)
	_ = x

	fmt.Println("Age =>", child)
	fmt.Println("Old =>", adult)

	//aliases
	//byte bir int8 aliasıdır
	//run bir int32 aliastır

	type second = uint

	var hour second = 3600

	fmt.Println("Minutes in an hour : ", hour/60)
}
