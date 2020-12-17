package main

import "fmt"

type telugu struct {
	name string
}
type english struct {
	name string
}

type print interface {
	getGreeting() string
}

func (t telugu) getGreeting() string {
	str := fmt.Sprintln("Namaste", t.name)
	return str
}

func (e english) getGreeting() string {
	str1 := fmt.Sprintln("Hello", e.name)
	return str1

}

func printing(p print) {
	fmt.Println(p.getGreeting())
}
func main() {

	tel1 := telugu{name: "Naveen"}
	eng1 := english{name: "nimmala"}

	printing(tel1)
	printing(eng1)

}
