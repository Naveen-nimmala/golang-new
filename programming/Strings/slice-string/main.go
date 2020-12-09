package main

import "fmt"

func main() {
	s1 := "Naveen"
	fmt.Println(s1[2:5]) // slicing a sting it return bytes not runes

	s2 := "నవీన్"
	fmt.Println(s2[0:2]) //  it will not print the proper output

	rs := []rune(s2) /* to print sting, we need to conver this as a rune,
	when we conver to new rune, it will craete new backing array, it not efficent */
	fmt.Printf("%T\n", rs)

	fmt.Println(string(rs[0:2])) // now it will print the first two characters of string
}
