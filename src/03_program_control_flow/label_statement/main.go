package main

import "fmt"

func main() {

	people := [5]string{"ali", "veli", "ahmet", "mehmet", "zahmet"}
	friends := [2]string{"kemal", "ahmet"}

outer:
	for index, name := range people {
		for _, friend := range friends {
			if name == friend {
				fmt.Printf("this friend is found  %q in this index %d\n", name, index)
				break outer
			}
		}
	}

	fmt.Println("First instruction after the label")

}
