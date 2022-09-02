package main

import "fmt"

func main() {
	defer helloWorld()
	f4("GoLang", 1, 2, 3, 4, 5)
}

func helloWorld() {
	fmt.Println("Hello World!")
}

func f1(a int, b int) {
	fmt.Println(a, b)
}

func f2(a, b int) int {
	return a + b
}

func f3() (int, int) {
	return 5, 6
}

func f4(b string, a ...int) {
	fmt.Printf("%T\n", a)
	fmt.Printf("%#v\n", a)
}
