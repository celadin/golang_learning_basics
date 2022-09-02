package main

import "fmt"

func main() {

	name := "Celalettin"
	fmt.Println("Address : ", &name)

	var x int = 10
	xPtr := &x

	fmt.Printf("X Pointer Type %T Value %v\n", xPtr, xPtr)

	*xPtr = 90
	fmt.Println("x : ", x , "ptr:", *xPtr)

	fmt.Println("===========================================")

	a := 0
	change(&a)
	fmt.Println("After changing a  value", a)

}

func change(a *int){
	*a = 100
}