package main

import "fmt"

func main() {
	var cities []string
	//cities[0] = 10 we cannot assign the values
	fmt.Println("Cities equal to nil:", cities == nil)
	fmt.Printf("%#v\n", cities)

	nums := []int{1, 2, 3}
	fmt.Println(nums)

	numbers := make([]int, 2)
	fmt.Printf("%#v\n", numbers)

	type names []string

	people := names{"Naveen", "Nimmala"}
	fmt.Println(people)

	name := people[0]
	fmt.Println("Name is", name)

	people[0] = "Naveen Kumar"
	fmt.Println(people)

	for i, v := range people {
		fmt.Printf("Index : %v, Value: %v\n", i, v)
	}

	values := nums

	fmt.Println(values)
}
