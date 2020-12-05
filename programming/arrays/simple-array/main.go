package main

import (
	"fmt"
	"strings"
)

func main() {
	// simple decleration
	var nums [4]int
	fmt.Printf("%v\n", nums)
	fmt.Printf("%#v\n", nums)
	// declare and assign with nil values
	var vals = [4]float64{}
	fmt.Printf("%#v\n", vals)
	// declare with values
	var val1 = [3]int{1, 2, 3}
	fmt.Printf("%#v\n", val1)
	//declare with short declreation
	str := [2]string{"Naveen", "Nimmala"}
	fmt.Printf("%#v\n", str)
	// declare and assign the few values of array, rest os of values will be nil values
	str1 := [5]string{"hello", "naveen"}
	fmt.Printf("%#v\n", str1)

	// declare and assign value with dynamic length
	val2 := [...]int{1, 2, 3, 4, 5}
	fmt.Printf("%#v  Length of Array : %v\n", str, len(val2))

	for i, v := range val2 {
		fmt.Printf("Index: %v and Value: %v\n", i, v)
	}
	fmt.Println(strings.Repeat("N", 20))
	for i := 0; i < len(val2); i++ {
		fmt.Printf("Index: %v and Value: %v\n", i, val2[i])
	}
}
