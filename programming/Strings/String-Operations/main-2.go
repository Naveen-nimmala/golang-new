package main

import (
	"fmt"
	"strings"
)

func main() {
	p := fmt.Println
	pf := fmt.Printf
	mynameRepeate := strings.Repeat("Naveen", 5)
	p(mynameRepeate)

	rePlace := strings.Replace("192.168.0.1", ".", ":", 2)
	p(rePlace) // output 192:168:0.1

	rePlace = strings.Replace("192.168.0.1", ".", ":", -1)
	p(rePlace) // output 192:168:0:1

	rePlace = strings.ReplaceAll("192.168.0.1", ".", ":")
	p(rePlace) // output 192:168:0:1

	s := strings.Split("N,a,v,e,e,n", ",")
	pf("%T\n", s)

	s = strings.Split("Naveen Nimmala Kumar", "")
	pf("%#v\n", s) // Output: []string{"N", "a", "v", "e", "e", "n", " ", "N", "i", "m", "m", "a", "l", "a", " ", "K", "u", "m", "a", "r"}

	s = strings.Split("Naveen Nimmala Kumar", " ")
	pf("%#v\n", s) //Output: []string{"Naveen", "Nimmala", "Kumar"}

	s = []string{"naveen", "Nimmala", "kumar"}
	mystr := strings.Join(s, "-")
	p(mystr) // Output: naveen-Nimmala-kumar

	mystr = "Orange Greean Bue yelllow"
	fields := strings.Fields(mystr)
	pf("%T\n", fields)  // Output: []string
	pf("%#v\n", fields) //output: []string{"Orange", "Greean", "Bue", "yelllow"}
	p(fields)           // output: [Orange Greean Bue yelllow]

	ss1 := strings.TrimSpace("\t	Naveen		Nimmala			Kumar  !")
	pf("%q\n", ss1)

	ss2 := strings.Trim("...Naveen Nimmala Kumar!!!", ".!")
	pf("%q\n", ss2)

}
