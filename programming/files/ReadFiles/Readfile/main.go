package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	file, err := os.Open("hello.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	byteSlice := make([]byte, 2)

	numberByetsRead, err := io.ReadFull(file, byteSlice)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("No of Bytes read: %d\n", numberByetsRead)
	log.Printf("Data read: %s\n", byteSlice)

	file, err = os.Open("main.go")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// this is antoher way read the file
	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Data as String : %s\n", data)
	fmt.Printf("Length of data: %d\n", len(data))

	// another way to read the fule
	data, err = ioutil.ReadFile("hello.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("data: %s\n", data)

}
