package main

import (
	"fmt"
	"strings"
)

type animals []string

func (a animals) Print() {
	for _, v := range a {
		fmt.Println(v)
	}
}

type car struct {
	brand string
	price int
}

func (c *car) changeCar(newBrand string, newPrice int) {
	(*c).brand = newBrand
	(*c).price = newPrice
}

func main() {

	var pets animals

	pets = []string{"dog", "cat", "hen"}

	pets.Print()

	fmt.Println(strings.Repeat("#", 20))

	audi := car{"Audi", 40}
	fmt.Println(audi)
	audi.changeCar("Bmw", 50)
	fmt.Println(audi)
}
