package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}

type rect struct {
	width, height float64
}

type shapes interface {
	calArea() float64
	calPeram() float64
	//greeting()
}

func (c circle) calArea() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}
func (r rect) calArea() float64 {
	return r.height * r.width
}

func (c circle) calPeram() float64 {
	return 2 * math.Pi * c.radius
}
func (r rect) calPeram() float64 {
	return 2 * (r.height + r.width)
}

func greeting() {
	fmt.Println("Hello")
}

func printing(s shapes) {
	fmt.Printf("shape type: %T\n", s)
	fmt.Printf("Area %v\n", s.calArea())
	fmt.Printf("Perm type: %v\n", s.calPeram())

}

func main() {
	cr1 := circle{radius: 1}
	rt1 := rect{height: 1.1, width: 2.2}
	printing(cr1)
	printing(rt1)
}
