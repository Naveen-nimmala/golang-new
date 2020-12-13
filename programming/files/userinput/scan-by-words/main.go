package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	input := bufio.NewScanner(os.Stdin)
	input.Split(bufio.ScanWords)
	marks := make([]string, 0)
	count := 0
	for input.Scan() {
		fmt.Println(input.Text())
		marks = append(marks, input.Text())
		count++
		if count > 5 {
			break
		}

	}
	if err := input.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(len(marks))
}
