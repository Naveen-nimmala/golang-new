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
	total(marksall)

}

func total(marksall []int) {
	total := 0
	fail := 0
	for _, v := range marksall {
		if v > 100 {
			log.Fatal("Marks Should be under 100")
		}
		if v <= 35 {
			fail++
		}
		total += v

	}
	if fail != 0 {
		fmt.Printf("You have failed in %v subject[s]\n", fail)
		fmt.Println("Total Marks", total)
	} else {
		fmt.Println("Total Marks", total)
		grade(total)
	}

}

func grade(total int) {
	switch {
	case total < 300:
		fmt.Println(" Your Grade is Third Class")
	case total <= 360:
		fmt.Println(" Your Grade is Second Class")
	case total > 360:
		fmt.Println(" Your Grade is First Class")
	}
}
