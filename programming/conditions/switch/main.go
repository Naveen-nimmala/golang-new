package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {

	language := "golanwwwg"

	num := os.Args[1:]
	if len(num) != 1 {
		fmt.Println("Please pass the integer Value")
	}
	n, _ := strconv.Atoi(num[0])

	switch true {
	case n%2 == 0:
		fmt.Printf("%d is an even number\n", n)
	case n%2 != 0:
		fmt.Printf("%d is an odd number\n", n)
	}
	switch language {
	case "Python":
		fmt.Println("This is Python")
	case "java":
		fmt.Println("This is Java")
	case "golang", "Go":
		fmt.Println("This is golang")
	default:
		fmt.Println("Start Learning any one of programming language ")
	}

	hour := time.Now().Hour()
	fmt.Println(hour)

	switch {
	case hour < 12:
		fmt.Println("Good Morning")
	case hour < 17:
		fmt.Println("Good Afternoon")
	default:
		fmt.Println("good Evening")
	}
}
