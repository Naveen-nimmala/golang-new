package main

import (
	"fmt"
)

func main() {
	const days int = 7
	const pi float64 = 3.14
	const Seconds int = 60

	munits := 60

	fmt.Printf("Total %v \n", Seconds*munits)

	const a = 5
	const b = 0

	// fmt.Println(a / b)  // const varibales will find the errors in compile time
	// in above example 5 cannot be devided by 0, it will show the errors in compile time

	const x, y, z int = 1, 2, 3
	const xx, yy = 1, 2

	const ( // Declaring multiple const variables
		aa = 100
		bb = -100
		cc = 200
	)
	fmt.Println(aa, bb, cc)

	const (
		aaa = 100
		bbb
		ccc
	)
	fmt.Println(aaa, bbb, ccc) // it will print the 100 three times

	// const value = math.Pow(2, 3) // const can not define at runtime --> math functoin is runtime function

	val1 := 10
	_ = val1
	// const val2 = val1 // you cannot assign the value of variable to const

	const l1 = len("naveen") //len functoin is runtime function, we can assign the values

}
