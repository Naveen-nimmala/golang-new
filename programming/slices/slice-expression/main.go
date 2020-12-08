package main

import "fmt"

func main() {

	a := [6]int{1, 2, 3, 4, 5, 6} // a is an array with fixed length

	b := a[1:3] // b is an slice

	fmt.Printf("%T %v and %T %v\n", a, a, b, b)

	s1 := []int{1, 2, 3, 4, 5, 6}
	s2 := s1[1:3]

	s2[1] = 1000 // if you change the slice s2, it will change s1 also, because both are accessing same backing array

	fmt.Println(s1, s2)

	// if you want slice should be with own backing array We Should use append method

	sl1 := []int{1, 2, 3, 4}

	sl2 := append(sl1[1:3])
	fmt.Println(sl1, sl2)

	// adding more elements
	sl2 = append(sl1[1:3], 100, 200)
	fmt.Println(sl2)

}
