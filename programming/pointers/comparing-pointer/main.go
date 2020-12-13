package main

import "fmt"

func main() {
	a := 5.5
	p1 := &a
	pp1 := &p1
	fmt.Printf("Value of a: %v and Address of a: %v\n", a, &a)
	fmt.Printf("Value of p1: %v and Address of p1: %v\n", p1, &p1)
	fmt.Printf("Value of pp1: %v and Address of pp1: %v\n", pp1, &pp1)

	/* OUTPUT
	Value of a: 5.5 and Address of a: 0xc0000b4008
	Value of p1: 0xc0000b4008 and Address of p1: 0xc0000ae018
	Value of pp1: 0xc0000ae018 and Address of pp1: 0xc0000ae020 */
}
