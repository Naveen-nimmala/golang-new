package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	fmt.Printf("Please enter someting ")
	input.Scan()

	text := input.Text()
	bytess := input.Bytes()
	fmt.Println("Input Text:", text)
	fmt.Println("Input Bytes", bytess)

	for input.Scan() { // scaning coninously
		text = input.Text()
		fmt.Println("You Entered", text)
		if text == "stop" {
			fmt.Println("Existing scanning")
			break
		}
	}
	if err := input.Err(); err != nil {
		fmt.Println("Something wrong")
	}
}
