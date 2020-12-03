package main

import "fmt"

func main() {
	const a int = 10 // typed const
	const b = 10.5   // untyped const

	const v = 10 > 5
	fmt.Println(v)

	// const x int = 5
	// const y float64 = 2.5 * x // we cannot multiply the const float with const int value
	const x = 5
	const y = 2.5 * x // we can multiplu the untype const with untype int
	fmt.Printf("%T\n", y)

	var i int = x // an untype const value can be assign to different type
	var j float64 = x
	var p byte = x

	fmt.Println(i, j, p)
	fmt.Printf("%T %T %T\n", i, j, p)

	const r = 5
	var rr = r
	fmt.Printf("%T  %T \n", r, rr)

}
