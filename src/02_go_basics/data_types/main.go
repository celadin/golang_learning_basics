package main

import "fmt"

func main() {

	//RUNE TYPE
	var r rune = 'a'
	fmt.Printf("%T\n", r)
	fmt.Printf("%x\n", r)

	//BOOL
	var b bool = true
	fmt.Printf("%T\n", b)

	//STRING
	var s string = "Hello World!"
	fmt.Printf("%T\n", s)

	//ARRAY
	var arr = [4]int{1, 2, 3, 4}
	fmt.Printf("%T\n", arr)

	//SLICE
	var cities = []string{"London", "Batman", "İstanbul"}
	fmt.Printf("%T\n", cities)

	//MAP
	currencies := map[string]float32{
		"USD": 17.8,
		"EUR": 18.2,
	}
	fmt.Printf("%T\n", currencies)
	fmt.Printf("%f\n", currencies["USD"])
	fmt.Printf("%f\n", currencies["EUR"])

	//STRUCT
	type Person struct {
		name string
		age  int
	}

	var you Person
	fmt.Printf("%T\n", you)

	//POINTER
	var x int = 1
	ptr := &x
	fmt.Printf("ptr type is %T with the value of %v\n", ptr, ptr)

	//FUNC
	fmt.Printf("%T\n", f)

}

func f(){

}
