package main

import "fmt"

func main() {

	for i := 0; i < 5; i++ {
		if i%2 != 0 {
			continue
		}
		fmt.Println(i)
	}

	//skipp the 7 in the loop

	for i := 0; i < 10; i++ {
		if i == 7 {
			continue
		}
		fmt.Println(i)
	}
}
