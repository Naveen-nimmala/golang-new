package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	var newFile *os.File
	fmt.Printf("%T\n", newFile)

	newFile, err := os.Create("hello.txt")
	if err != nil {
		//fmt.Println(err)
		// os.Exit(1)
		log.Fatal(err) //this will exit the programe and throw the error
	}

	err = os.Truncate("hello.txt", 50) // make file with 0 or more bytes.. if you have data already in that, it will destroy
	if err != nil {
		//fmt.Println(err)
		// os.Exit(1)
		log.Fatal(err) //this will exit the programe and throw the error
	}
	newFile.Close()

	file, err := os.Open("hello.txt")
	if err != nil {
		log.Fatal(err)
	}
	file.Close()

	file1, err1 := os.OpenFile("hello.txt", os.O_CREATE|os.O_APPEND, 0644) /* this will create a file if doent exists,
	it will append the data if file already exists */
	if err != nil {
		log.Fatal(err1)
	}
	file1.Close()
	// File information
	var fileInfo os.FileInfo
	fileInfo, err = os.Stat("hello.txt")
	fmt.Println("Filename:", fileInfo.Name())
	fmt.Println("FIleSize:", fileInfo.Size())
	fmt.Println("Modification time:", fileInfo.ModTime())
	fmt.Println("Directory?:", fileInfo.IsDir())
	fmt.Println("Permissions:", fileInfo.Mode())

	// check file exists or not
	fileInfo, err = os.Stat("sample.txt")
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("File doesnt exists")
		}
	}

	// rename the file
	oldPath := "hello.txt"
	newPath := "abc.txt"
	err = os.Rename(oldPath, newPath)
	if err != nil {
		log.Fatal(err1)
	}
	// Remove the file
	err = os.Remove("abc.txt")
	if err != nil {
		log.Fatal(err1)
	}
}
