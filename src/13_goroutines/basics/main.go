package main

import (
	"fmt"
	"runtime"
	"time"
)

func f1() {
	fmt.Println("f1() execution is started")

	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}

	fmt.Println("f1() execution is completed")
}

func f2() {
	fmt.Println("f2() execution is started")

	for i := 4; i < 7; i++ {
		fmt.Println(i)
	}

	fmt.Println("f2() execution is completed")
}

func main() {
	fmt.Println("Number of CPU's :", runtime.NumCPU())
	fmt.Println("Number of Goroutines :", runtime.NumGoroutine())
	fmt.Println("GOMAXPROCS :", runtime.GOMAXPROCS(0))
	fmt.Println("OS :", runtime.GOOS)
	fmt.Println("Arch :", runtime.GOARCH)

	go f1() //first goroutine

	f2()

	time.Sleep(time.Second * 3) // in order to see f1() execution results

	fmt.Println("main execution is completed")
}
