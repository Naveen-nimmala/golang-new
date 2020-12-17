package main

import "fmt"

func main() {
	var ch chan int
	fmt.Println(ch) // nil
	ch = make(chan int)
	fmt.Println(ch) // 0xc00008c060 channels like pointers

	c := make(chan int)

	// <- channel operators

	//SEND
	c <- 10

	//REVICE
	num := <-c
	_ = num
	fmt.Println(<-c)
	close(c)
}
