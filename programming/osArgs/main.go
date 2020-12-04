package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("os args", os.Args)
	fmt.Println("os args 0 position", os.Args[0])
	fmt.Println("os args 1 postion ", os.Args[1])

	values := os.Args[1:]

	fmt.Println(values)
	fmt.Println(len(values))

	result, _ := strconv.ParseFloat(os.Args[1], 64)
	fmt.Printf("%T : %f \n", result, result)
}
