package main

import (
	"fmt"
	"strings"
)

type student struct {
	name string
	mark int
}

type college struct {
	name string
}

type alldata interface {
	getInfo() string
}

func (s student) getInfo() string {
	name := fmt.Sprint("Name is ", s.name)
	//marks, _ := fmt.Println("Your Marks is :", s.marks)

	return name + "STUDENT DAAAA"
}

func (c college) getInfo() string {
	clgName := fmt.Sprint("Clg Name is", c.name)
	return clgName

}

func extraStudent() {
	fmt.Println("This is extra student")
}

func allInfo(a alldata) {
	fmt.Println()
	fmt.Printf("%T\n", a)
	fmt.Println(strings.Repeat("#", 20))
	switch value := a.(type) {
	case college:
		fmt.Printf("This type %T clg %v\n", value, value)
		fmt.Println(a.getInfo())
	case student:
		fmt.Println("This is Student", value)
		fmt.Println(a.getInfo())
		extraStudent()
	}

}

func (s student) volume() string {
	return s.name
}

func main() {
	naveen := student{name: "Naveen Nimmala", mark: 256}
	clgName := college{name: "Audi"}
	allInfo(naveen)
	allInfo(clgName)

	//type assertions
	var t alldata = student{name: "Tom"}
	fmt.Printf("%T\n", t)
	valum, ok := t.(student)
	if ok == true {
		fmt.Println("New Name", valum.volume())
		fmt.Println(valum.getInfo())
	}
	// t = college{name: "CBIT"} // changing the t interface type, now t type is college
	// switch value := t.(type) {
	// case college:
	// 	fmt.Println("This clg", value)
	// case student:
	// 	fmt.Println("This is Student", value)
	// }
}

/* Output

main.student
####################
This is Student {Naveen Nimmala 256}
Name is Naveen NimmalaSTUDENT DAAAA
This is extra student

main.college
####################
This type main.college clg {Audi}
Clg Name isAudi
main.student
New Name Tom
Name is TomSTUDENT DAAAA

*/
