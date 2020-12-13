package main

import "fmt"

func sum(a int, b int) {
	fmt.Println(a + b)
}

func sum1(a, b, c int, d, e float64, s string) {
	fmt.Println(a, b, c, d, e, s)

}

func sum3(a, b int) int {
	return a + b
}

func sum4(a, b int) (int, int) { // unmaed retuns
	return a + b, a * b
}

func sum6(a, b int) (dd int, cc int) { // named returns
	dd = a + b
	cc = a * b
	return
}

func main() {
	sum(1, 2)

	sum1(1, 2, 3, 3.4, 4.5, "Naveen")

	x := sum3(1, 2)
	fmt.Println(x)

	n, m := sum4(1, 2)
	fmt.Println(n, m)

	dd, cc := sum6(1, 2)
	fmt.Println(dd, cc)
}
