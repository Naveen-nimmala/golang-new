package main

import "fmt"

func main() {

	a := map[string]string{"a": "A"}
	b := map[string]string{"a": "A"}

	//mt.Println(a == b) we cannot compare, map can only compare with nil maps
	s1 := fmt.Sprintf("%s", a)
	s2 := fmt.Sprintf("%s", b)
	fmt.Println(s1, s2)

	if s1 == s2 {
		fmt.Println("Maps are equal")
	} else {
		fmt.Println("maps are not equal")
	}

}
