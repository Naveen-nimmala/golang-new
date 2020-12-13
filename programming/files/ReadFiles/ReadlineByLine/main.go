package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("hello.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// if you want to scan word by word or line by line or rune by rune or byte by byte
	scanner.Split(bufio.ScanWords)
	// scanner.Split(bufio.ScanRunes)
	// scanner.Split(bufio.ScanLines)
	// scanner.Split(bufio.ScanBytes)
	success := scanner.Scan()

	if success == false {
		err = scanner.Err()
		if err == nil {
			log.Println("Scan was completed EOF ")
		} else {
			log.Fatal(err)
		}
	}
	fmt.Println("First line found", scanner.Text()) // scanner.Text() may functions we can use

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
