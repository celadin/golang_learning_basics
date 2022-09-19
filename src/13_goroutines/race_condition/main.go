package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	gr := 100

	wg.Add(gr * 2)

	n := 0

	for i := 0; i < gr; i++ {

		go func() {
			time.Sleep(time.Second / 10)
			n++
			wg.Done()
		}()

		go func() {
			time.Sleep(time.Second / 10)
			n--
			wg.Done()
		}()

		
	}

	wg.Wait()

	fmt.Printf("Last value of n : %d\n", n)

	fmt.Println("DONE")
}
