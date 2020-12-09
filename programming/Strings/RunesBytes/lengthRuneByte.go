package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	ss1 := "Naveen"
	fmt.Println(len(ss1)) // lenght is 6

	name := "˜å√´´˜"
	teluguName := "˜ˆµµå¬å"
	fmt.Println(len(name))       // here total characters are 6 but length is 14
	fmt.Println(len(teluguName)) // here total characters are 3 but length is 15
	fmt.Println(teluguName)
	for _, v := range teluguName {
		fmt.Printf("%c\n", v)
	}

	for i := 0; i < len(teluguName); {
		value, size := utf8.DecodeRuneInString(teluguName[i:])
		fmt.Printf("%c and size is %d\n", value, size)
		i += size
	}
	n := utf8.RuneCountInString(name) // its just count the no of characters in rune with string
	tn := utf8.RuneCountInString(teluguName)
	tnn := utf8.RuneCountInString("😇")
	fmt.Println(n, tn, tnn) // output is 6 , 7, 1

}
