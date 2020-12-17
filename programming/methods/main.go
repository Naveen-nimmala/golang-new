package main

import "fmt"

type names []string

type students struct {
	name   string
	rollNo int
}

func (n names) printing() {
	for i, name := range n {
		fmt.Println(i, name)
	}
}

func (n names) changing() {
	n[0] = "hello"
	n[1] = "world"
}

func (n students) changingStruct() {
	n.name = "NN"
	n.rollNo = 25
}

func (n *students) changingPointerStruct() {
	n.name = "NN"
	n.rollNo = 25
}

func (n *students) changingPointerStructValues(roll int) {
	n.name = "NN"
	n.rollNo = 25 + roll
}

func main() {
	friends := names{"Naveen", "Nimmala"}
	friends.printing()

	friends.changing()   // this is slice, it doesnt require pointer to pass, by deafult it passes the pointer
	fmt.Println(friends) // output: [hello world]

	naveen := students{name: "Naveen Nimmala", rollNo: 50}
	naveen.changingStruct()
	fmt.Println(naveen) // output: {Naveen Nimmala 50}
	naveen.changingPointerStruct()
	fmt.Println(naveen) // output: {NN 25}

	naveen.changingPointerStructValues(10)
	fmt.Println(naveen) //output: {NN 35}

}
