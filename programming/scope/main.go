package main

import (
	"fmt"
	f "fmt"
)

// file scope

const val1 int = 1 // package scope

func main() {

	var2 := "Naveen" // Local block scope
	f.Println("Hell", var2)

	const val1 int = 10
	f.Println(val1)

	fun()

}

func fun() {
	fmt.Println("Hello World")
}
