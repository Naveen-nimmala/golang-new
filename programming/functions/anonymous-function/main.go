package main

import "fmt"

func main() {

	// basic example
	func(msg string) {
		fmt.Println(msg)
	}("Im anonymous function")

	a := increment(10)
	fmt.Printf("%T\n", a) // OUtPut: func() int -->> this is func type which return the func int value
	fmt.Println(a())      // output: 11
	fmt.Println(a())      /// output: 12

}

func increment(x int) func() int {
	return func() int {
		x++
		return x
	}
}
