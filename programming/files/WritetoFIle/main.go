package main

import (
	"io/ioutil"
	"log"
	"os"
)

func main() {
	file, err := os.OpenFile(
		"hello.txt",
		os.O_WRONLY|os.O_TRUNC|os.O_CREATE,
		0644,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	// write some daya into file
	byteSlice := []byte("I Love India") // conver the string to byte slice
	byteWritten, err := file.Write(byteSlice)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Bytes Written: %d\n", byteWritten)

	// File operations using Ioutil
	bs := []byte("Naveen Nimmala")
	err = ioutil.WriteFile("naveen.txt", bs, 0644) // it will create a file if it doesnt exists, it will truncate the file if it exists
	if err != nil {
		log.Fatal(err)
	}
}
