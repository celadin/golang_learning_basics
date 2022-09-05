package main

import (
	"fmt"
	"strings"
)

type animals []string

func (a animals) Print(){
	for _, v := range a {
		fmt.Println(v)
	}
}

func main() {

	var pets animals

	pets = []string{"dog", "cat", "hen"}

	pets.Print()

	fmt.Println(strings.Repeat("#", 20))
}
