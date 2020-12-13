package main

import "fmt"

func main() {
	name := "Naveen"
	fmt.Println(&name) // address of varible

	var x int = 10
	ptr := &x
	fmt.Printf("ptr is of type %T with a value of %v\n", ptr, ptr) //output: ptr is of type *int with a value of 0xc000016048
	fmt.Printf("Address of x is %p\n", &x)                         //Address of x is 0xc000016048
	fmt.Printf("Address of ptr is %p\n", &ptr)                     //Address of ptr is 0xc0000ae020

	var ptr1 *float64 // declearing pointer type to that varibale
	_ = ptr1

	a := 10
	b := &a
	*b = 200           // this is equal to a = 200 --> referencing operator , that we want the value which pointer pointing to
	fmt.Println(a, *b) // output: 200 200

	*b = 400    // a = 400
	*b = *b / 2 // a = a/2
	fmt.Println(a)
	// &Value ==> Pointer (address)
	// *pointer ==> Value

}
