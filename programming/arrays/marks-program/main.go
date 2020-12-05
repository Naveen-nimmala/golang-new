package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	var marksall []int
	marks := os.Args[1:]
	if len(marks) != 6 {
		fmt.Println("Please pass 6 subject marks with space")
		return
	}
	for _, v := range marks {
		mark, err := strconv.Atoi(v)
		if err != nil {
			log.Fatal("Please pass int values", err)
		}
		marksall = append(marksall, mark)
	}
	total := 0
	for _, v := range marksall {
		if v > 100 {
			log.Fatal("Marks Should be under 100")
		}
		total += v

	}
	fmt.Println("total Marks", total)
}
