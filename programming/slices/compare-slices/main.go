package main

import "fmt"

func main() {
	var x []int
	fmt.Println(x == nil) // true becasue slice x is nil

	y := []int{}
	fmt.Println(y == nil) // false , becasue  y slice intilazied with {}

	//compare slices with values

	a, b := []int{1, 2, 3}, []int{11, 2, 3}
	// fmt.Println(a == b)
	var equal bool = true
	// before we compare we need to check weather both lenght of slices are equal or not
	if len(a) != len(b) {
		equal = false
	}

	for i, vA := range a {
		if vA != b[i] {
			equal = false
			break
		}
	}

	if equal {
		fmt.Println("a and b  are equal")
	} else {
		fmt.Println("a and b are not equal")
	}
}
