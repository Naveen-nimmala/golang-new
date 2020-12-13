package main

import (
	"fmt"
	"strings"
)

func fun1(a ...int) { // when we dont know the arguments length, we use variadic paraments with three dots ...
	fmt.Printf("%T\n", a) //[]int
	fmt.Printf("%#v\n", a)
}

func fun2(a ...int) {
	a[0] = 256
}

func sumAndProd(a ...float64) (float64, float64) {
	sum := 0.0
	prod := 1.0
	for _, v := range a {
		sum += v
		prod *= v
	}
	return sum, prod
}

func persionInfo(age int, names ...string) string {
	fullName := strings.Join(names, "@")
	returnString := fmt.Sprintf("Age: %d, fullName: %s\n", age, fullName)
	return returnString
}
func main() {
	fun1(1, 2, 3, 4) //[]int{1, 2, 3, 4}
	fun1()           //[]int(nil)

	nums := []int{1, 2}
	nums = append(nums, 3, 4) // this is variadic function
	fun1(nums...)             //[]int{1, 2, 3, 4}

	fun2(nums...)
	fmt.Println(nums) //[256 2 3 4] it chanes the value of num

	s, p := sumAndProd(1.0, 2.0, 3.0, 4., 5.5, 6.5, 45.6)
	fmt.Println(s, p)

	info := persionInfo(30, "Naveen", "Kumar", "Nimmala")
	fmt.Println(info) //Age: 30, fullName: Naveen@Kumar@Nimmala
}
