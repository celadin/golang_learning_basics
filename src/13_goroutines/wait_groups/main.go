package main

import (
	"fmt"
	"sync"
	"time"
)

func f1(wg *sync.WaitGroup) {
	fmt.Println("f1() execution is started")

	for i := 0; i < 3; i++ {
		fmt.Println(i)
		time.Sleep(time.Second)
	}
	fmt.Println("f1() execution is completed")

	wg.Done()
}

func f2() {
	fmt.Println("f2() execution is started")

	for i := 4; i < 7; i++ {
		fmt.Println(i)
	}

	fmt.Println("f2() execution is completed")
}

func main() {

	var wg sync.WaitGroup

	wg.Add(1)

	go f1(&wg)

	f2()

	wg.Wait()
	fmt.Println("main execution is completed")
}
