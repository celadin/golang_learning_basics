package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	i, _ := strconv.Atoi(os.Args[1])
	
	switch i {
	case 0:
		fmt.Println(0)
	case 1:
		fmt.Println(1)
	case 2:
		fmt.Println(2)
	default:
		fmt.Println("default")
	}

	fmt.Println("after switch")
}
