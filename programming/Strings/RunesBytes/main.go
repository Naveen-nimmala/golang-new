package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {

	var1, var2 := 'a', 'b'
	fmt.Printf("%T and %v\n", var1, var2) // int 32 is called as Rune as well , Byte can be called as int8 as well

	str := "😀"       // if its a single byte it should print the length is 1
	str1 := "😀 ß∂ƒ©" // if its a single byte it should print the length is 1

	fmt.Println(len(str)) /* but The output is 4 , because its depends on characher it changes the byte size,
	so when we use length function we can exactly take the lenght of strings */
	fmt.Println(len(str1)) // length is 14

	fmt.Println(str, str1)

	/// if you want to ittrate the character by character(rune by rune) (not byte by byte) use below method

	newstr, sz := utf8.DecodeRuneInString("😀 ß∂ƒ©")
	fmt.Printf("%c\n , %d\n", newstr, sz)

	fmt.Println("\n" + strings.Repeat("*", 20)) //
	strng := "🥰 œ∑´®†"
	for i := 0; i < len(strng); { // this is way to print the runes
		run, size := utf8.DecodeRuneInString(strng[i:])
		fmt.Printf("%c", run)
		i += size
	}
	fmt.Println("\n" + strings.Repeat("*", 20))
	for _, runee := range strng { // this is another way to print the runes
		fmt.Printf("%c", runee)
	}
	fmt.Println("\n" + strings.Repeat("*", 20))

}
