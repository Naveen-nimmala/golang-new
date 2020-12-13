package main

import "fmt"

func main() {

	type Contact struct {
		email, address string
		phone          int
	}
	type Employee struct {
		name        string
		salary      int
		ContactInfo Contact
	}

	Naveen := Employee{
		name:   "Naveen Nimmala",
		salary: 10000,
		ContactInfo: Contact{
			email:   "naveen@gmail.com",
			address: "India",
			phone:   1234567,
		},
	}
	fmt.Printf("%#v\n", Naveen)
	fmt.Println(Naveen.name)
	fmt.Println(Naveen.ContactInfo.phone)
	Naveen.ContactInfo.address = "Malaysia"
	fmt.Println(Naveen.ContactInfo.address)
}
