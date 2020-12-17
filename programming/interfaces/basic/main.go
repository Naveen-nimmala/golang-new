package main

import "fmt"

type telugu struct{}
type english struct{}

type print interface {
	getGreeting() string
}

func (t telugu) getGreeting() string {
	return "Namastee"
}

func (e english) getGreeting() string {
	return "Hello"
}

func printing(p print) {
	fmt.Println(p.getGreeting())
}
func main() {

	tel1 := telugu{}
	eng1 := english{}

	printing(tel1)
	printing(eng1)

}
