package main

import "fmt"

func main() {
	a := 5.5
	p1 := &a
	pp1 := &p1
	fmt.Printf("Value of a: %v and Address of a: %v\n", a, &a)
	fmt.Printf("Value of p1: %v and Address of p1: %v\n", p1, &p1)
	fmt.Printf("Value of pp1: %v and Address of pp1: %v\n", pp1, &pp1)
	fmt.Printf("Value of a: using pp1 %v:\n", **pp1)

	/* OUTPUT
	Value of a: 5.5 and Address of a: 0xc0000b4008
	Value of p1: 0xc0000b4008 and Address of p1: 0xc0000ae018
	Value of pp1: 0xc0000ae018 and Address of pp1: 0xc0000ae020
	Value of a: using pp1 5.5	*/

	// Comparing pointers

	x := 10
	y := &x
	xx := 10
	yy := &xx
	z := &x
	fmt.Println("y is equal to yy ", y == yy) // Output: y is equal to yy  false, when pointers pointing to different addresses, it shows false
	fmt.Println("z is equal to y ", y == z)   // Output: z is equal to y  true, both y and z are pointing to &x , result become true

}
