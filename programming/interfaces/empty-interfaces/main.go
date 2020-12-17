package main

import "fmt"

func printAll(i ...interface{}) {
	fmt.Println(i)

}
func main() {
	printAll(10)
	printAll("Naveen")
	printAll(1, 2, 3, 4)
}
