package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("os.Args", os.Args)

	var result, err = strconv.ParseFloat(os.Args[1], 64)
	if err == nil {
		fmt.Printf("%.2f", result)
	}
}
