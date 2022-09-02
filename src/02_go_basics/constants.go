package main

import "fmt"

func main2() {
	const a int = 1
	const b = 2

	const c, d int = 6, 7

	const (
		e = 10
		f
		g
	)

	const aa = 1
	const bb float64 = aa

	const (
		c1 = iota
		c2 = iota
		c3 = iota
	)

	const (
		c4 = 1 << iota
		c5 = 1 << iota
		c6 = iota
		c7 = 1 << iota
	)

	const x = 7
	const y float32 = 3.1
	fmt.Println(x * y)

}
