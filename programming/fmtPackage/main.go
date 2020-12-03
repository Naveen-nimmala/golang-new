package main

import "fmt"

const value = 10

func main() {

	name := "Naveen"
	age := 28
	height := 5.1023456789
	job := true

	fmt.Printf("This is Decimal: Age : %d \n", age)
	fmt.Printf("This is Float: height : %f \n", height)
	fmt.Printf("This is String: Name: %s \n", name)
	fmt.Printf("This is Bool: Job: %t \n", job)
	fmt.Println()

	fmt.Printf("This is Value all items accepts: Name: %v , Age: %v, Height: %v, job: %v \n", name, age, height, job)

	fmt.Printf("This is float with points value: %.4f \n", height)

	fmt.Printf("This is quoted string %q \n", name)

	fmt.Printf("This is an expression: %f \n ", float64(value*age))

	fmt.Printf("This is Type: %T \n ", job)
	fmt.Printf("This is Byte: %b \n ", value)
	fmt.Printf("This is Byte: %08b \n ", value)
	fmt.Printf("This is Type: %x \n ", value)

}
