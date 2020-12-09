package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	file, err := os.OpenFile("hello.txt", os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	bs := []byte{97, 98, 99}
	bufferWritter := bufio.NewWriter(file)
	bytesWritten, err := bufferWritter.Write(bs)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Bytes written to buffer(not file) %d\n", bytesWritten)

	bytesAvibalbe := bufferWritter.Available()
	log.Printf("Bytes Avilabe in buffer %d\n", bytesAvibalbe)

	bytesWritten, err = bufferWritter.WriteString("\nHello World")
	if err != nil {
		log.Fatal(err)
	}
	unflushBufferSize := bufferWritter.Buffered()
	log.Printf("Bytes buffred: %d\n", unflushBufferSize)

	bufferWritter.Flush()
}
