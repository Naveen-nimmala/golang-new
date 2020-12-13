package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	var all []int
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter the marks with \",\" ")
	scanner.Scan()
	in := scanner.Text()
	marksall := strings.Split(in, ",")
	if len(marksall) > 7 {
		log.Printf("Please enter 6 subjects marks, You have Entered %d", len(marksall))
	}

	for _, v := range marksall {
		mark, err := strconv.Atoi(v)
		if err != nil {
			log.Fatal("Please pass int values", err)
		}
		all = append(all, mark)
	}
	total(all)

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
