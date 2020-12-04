package main

import "fmt"

func main() {

	for i := 0; i <= 4; i++ {
		fmt.Println("Naveen", i)
	}

	for i := 0; i < 5; {
		fmt.Println("Nimmala", i)
		i++
	}
	i := 0
	for i < 5 {
		fmt.Println("Sahi", i)

	}

	// sum := 0
	// for {
	// 	sum++
	// }
	// fmt.Println(sum) // this is Infinite loop
}
