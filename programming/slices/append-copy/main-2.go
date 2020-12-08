package main

import "fmt"

func main() {

	slc1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	var slc2 []int
	slc3 := make([]int, 5)
	slc4 := []int{11, 12, 13, 14, 15, 16}
	slc5 := make([]int, len(slc1))

	// slc1 to slc2 copy --> this will copy nothing because slc2 length is zero
	cp1 := copy(slc2, slc1)
	fmt.Println(slc2)

	// slc1 to slc3 copy --> this will copy only five elements form the slc1 to slc3
	cp2 := copy(slc3, slc1)
	fmt.Println(slc3)

	// slc1 to slc4  --> this will copy 6 elements(slc4 length is 6) from slc1 to slc4 and exiting values will be overwride
	cp3 := copy(slc4, slc1)
	fmt.Println(slc4)

	//slc1 to slc5
	cp4 := copy(slc5, slc1) // this will copy entire slc1 elements to slc5, because both lenght of slice is same
	fmt.Println(slc5)
	_, _, _, _ = cp3, cp1, cp2, cp4 // there are the list of number of copy elements

	fmt.Println("Hello World")

}
