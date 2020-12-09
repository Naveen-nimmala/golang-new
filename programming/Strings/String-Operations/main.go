package main

import (
	"fmt"
	"strings"
)

func main() {

	p := fmt.Println
	//pf := fmt.Printf
	result := strings.Contains("I Love India!", "Love") // it will check the weather Love is exists in the string or not
	p(result)                                           // output: true

	result = strings.ContainsAny("I Love India!", "xz") // it will check x or y characters is exists or not
	p(result)                                           // output: false

	result = strings.ContainsRune("Naveen", 'n')
	p(result) // output: false

	result1 := strings.Count("Naveen", "n") // checking how many times n has been occured in the given string
	p(result1)                              // output: false

	p(strings.ToUpper("heLlo wORld"))
	p(strings.ToLower("heLlo wORld"))

	p("go" == "go")                                   // Output: true
	p("Go" == "go")                                   // Output: false  this not an efficent
	p(strings.ToUpper("Go") == strings.ToUpper("go")) // Output: true this is  an efficent

	p(strings.EqualFold("GO", "go")) // Output : true  This is the best way to comapre two strings

}
