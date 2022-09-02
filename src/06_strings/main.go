package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	var c rune = 100
	d := 'b'

	fmt.Printf("%T %d\n", c, c)
	fmt.Printf("%T %d", d, d)

	s := "İstanbul"

	fmt.Println("\n" + strings.Repeat("#", 30))

	fmt.Println("First Char :", s[0])

	fmt.Println("\n" + strings.Repeat("#", 30))

	for i := 0; i < len(s); {
		r, size := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%c", r)
		i = size + i
	}

	fmt.Println("\n" + strings.Repeat("#", 30))

	for _, r := range s {
		fmt.Printf("%c", r)
	}

	fmt.Println("\n" + strings.Repeat("#", 30))

	s2 := "中文维基是世界上"

	fmt.Println("s2", s2)

	fmt.Println("s2[0]", s2[0])
	fmt.Println("s2[0:2]", s2[0:2])

	s2r := []rune(s2)
	fmt.Println("s2r", s2r)
	fmt.Println("s2r[0:2]", s2r[0:2])
	fmt.Println("string(s2r[0])", string(s2r[0]))
	fmt.Println("string(s2r[0:2])", string(s2r[0:2]))
	fmt.Println("string(s2r[0:3])", string(s2r[0:3]))

	s2s := string(s2r)
	fmt.Println("s2s", s2s)


}
