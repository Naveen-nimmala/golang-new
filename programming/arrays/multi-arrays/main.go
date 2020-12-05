package main

import "fmt"

func main() {
	balances := [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println(balances)

	x := [3]int{1, 2, 3}

	y := x
	fmt.Println("X and Y are Equal", x == y)
	x[0] = -1
	fmt.Println("X and Y are Equal", x == y)

}
