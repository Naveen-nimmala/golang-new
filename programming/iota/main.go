package main

import "fmt"

func main() {
	const (
		c1 = iota // default 0
		c2 = iota // 1
		c3 = iota // 2
	)

	fmt.Println(c1, c2, c3)

	const ( // Its same as above
		cc1 = iota // 0
		cc2        // 1
		cc3        // 2
	)

	fmt.Println(c1, c2, c3)

	const ( //
		ccc1 = iota * 2 // 0 * 2 = 0
		ccc2            // 1 * 2 = 2
		ccc3            // 2 * 2 = 4
	)

	fmt.Println(ccc1, ccc2, ccc3)

	const ( // required values x = -2, y = -4 , z = -5 ( using black identifiers)
		x = -(iota + 2) // -2
		_
		y
		z
	)
	fmt.Println(x, y, z)
}
