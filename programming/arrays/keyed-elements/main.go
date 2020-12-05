package main

import "fmt"

func main() {
	// it will show the output in order as per index
	grades := [3]int{
		1: 10,
		0: 20,
		2: 30,
	}
	fmt.Println(grades)

	// You no need to specify the all the elements of the array, it will take default zero/nil/false values
	accouts := [3]int{
		2: 400,
	}
	fmt.Println(accouts)

	// you declare  dynamic array
	names := [...]string{
		5: "Naveen",
	}
	fmt.Println(names, len(names))

	/// if you dont specify the index, it will be added at the end if the array
	countries := [...]string{
		5: "India",
		"US",
		3: "Malaysia",
	}
	fmt.Printf("%#v\n", countries)
}
