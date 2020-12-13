package main

import "fmt"

func fun1() {
	fmt.Println("THis is fun1")
}
func fun2() {
	fmt.Println("THis is fun2")
}
func fun3() {
	fmt.Println("THis is fun3")
}
func main() {

	defer fun1()
	defer fun2()
	fun3()
	/* Output
	THis is fun3
	THis is fun2
	THis is fun1
	*/

}
