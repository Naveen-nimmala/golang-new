package main

import "fmt"

func main() {
	s1 := "Naveen"
	fmt.Println(s1)
	fmt.Println("\"Naveen\"") // to print in doublecouts
	fmt.Println(`"Naveen"`)   // to print in doublecouts its called raw string

	s2 := `Nimmala\n`
	fmt.Println(s2)

	fmt.Println("Hello World\nNaveen")

	fmt.Println(`
Hello World
Naveen`)

	fmt.Println("C:\\Users\\naveen")
	fmt.Println(`C:\Users\naveen`)

	var str3 = "I Love " + "Go " + "Programming"
	fmt.Println(str3 + "!")

	fmt.Println("Element at index 0 :", str3[0]) // It will print 73 which is decimal value of Character eg: I

	//str3[5]="N" // you cannot asign like this, Strings are immutable in Golang

	fmt.Printf("%s\n", str3) // Output : I Love Go Programming
	fmt.Printf("%q\n", str3) // Output : "I Love Go Programming"
}
