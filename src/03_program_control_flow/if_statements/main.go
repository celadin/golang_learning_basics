package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if i, err := strconv.Atoi("45"); err == nil {
		fmt.Printf("i = %d\n", i)
	}

	if args := os.Args; len(args) > 0 {
		fmt.Printf("Args = %v\n", args)
	}
}
