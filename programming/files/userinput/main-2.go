package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	inp := bufio.NewScanner(os.Stdin)
	inp.Scan()
	fmt.Println("Entered ", inp.Text())

	for inp.Scan() {
		fmt.Printf("Enter marks of %s subject", inp.Text())
		if err := inp.Err(); err != nil {
			log.Fatal(err.Error)

		}
	}
}
