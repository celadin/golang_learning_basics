package main

import "fmt"

type book struct {
	title  string
	author string
	year   int
}

func main() {
	myBook := book{title: "the wind of wind", author: "John Doe", year: 2020}
	fmt.Println(myBook)
	fmt.Printf("Title : %q \nAuthor : %q \nYear : %v\n", myBook.title, myBook.author, myBook.year)

	mySecondBook := myBook

	fmt.Println("Are they equal ? ", myBook == mySecondBook)

	mySecondBook.author = "Celalettin"

	fmt.Printf("%#v\n", myBook)
	fmt.Printf("%#v\n", mySecondBook)
}
