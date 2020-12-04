package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {

	i, err := strconv.Atoi("45a")
	if err != nil {
		//log.Fatal(err.Error)
		fmt.Println(err)
	} else {
		fmt.Println(i)
	}

	if val, err := strconv.Atoi("30"); err == nil {
		fmt.Printf("No error print value %v\n", val)
	} else {
		fmt.Println(err)
	}
	if x, y := 10, "naveen"; y == "naveen" && x == 10 {
		fmt.Println("x and y satisfied")
	} else {
		fmt.Println("Condition is not satsfied")
	}

	if args := os.Args[1:]; len(args) != 1 {
		fmt.Println("Please enter the 1 arguments")
		// } else if kms := reflect.ValueOf(args[0]).Kind(); kms == reflect.Float64 && kms == reflect.Float32 {
		// 	fmt.Printf("%d KMs in Miles is %v\n", kms, float64(kms)*0.621) //type for switches
	} else if km, err := strconv.Atoi(args[0]); err != nil {
		fmt.Println("You Passed non-interger.. checking Float or not....")
		time.Sleep(4 * time.Second)
		if stoF, err := strconv.ParseFloat(args[0], 64); err == nil {
			fmt.Println("You have passed Float value.. calclulating KMs in Miles.. wait a moment...")
			time.Sleep(4 * time.Second)
			fmt.Printf("%v KMs in Miles is %v\n", stoF, float64(stoF)*0.621)
		} else {
			fmt.Println("Input is not interger or FLoat. Please pass valid input", err)
		}
	} else {
		fmt.Printf("%d KMs in Miles is %v\n", km, float64(km)*0.621)
	}
}
