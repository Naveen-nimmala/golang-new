package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {

	x := 10
	y := 10.5
	z := "11"
	b := true

	fmt.Printf("%T : %v\n", x, x)
	fmt.Printf("%T : %v\n", y, y)
	fmt.Printf("%T : %v\n", z, z)
	fmt.Printf("%T : %v\n", b, b)
	itoF := float64(x)
	fmt.Printf("%T : %v\n", itoF, itoF)
	ftoI := int(y)
	fmt.Printf("%T : %v\n", ftoI, ftoI)

	// Float to String

	ftoS := fmt.Sprintf("%f", y)
	fmt.Printf("%T : %q\n", ftoS, ftoS)
	// Int to String
	itoS := fmt.Sprintf("%d", x)
	fmt.Printf("%T : %s\n", itoS, itoS)

	// String to Int

	stoI, err := strconv.Atoi(z)
	if err != nil {
		log.Fatal(err.Error)
	}
	fmt.Printf("%T : %v\n", stoI, stoI)

	// int to String
	itoS1 := strconv.Itoa(x)
	fmt.Printf("%T : %s\n", itoS1, itoS1)

	// converting string to float

	stoF, _ := strconv.ParseFloat(z, 64)
	fmt.Printf("%T : %f\n", stoF, stoF)
}
