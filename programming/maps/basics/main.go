package main

import "fmt"

func main() {

	pf := fmt.Printf
	var surNames map[string]string
	pf("%#v\n", surNames)
	pf("No of pairs: %d\n", len(surNames))                       // defail empty map value is nil
	pf("The value for %q is %q\n", "Naveen", surNames["Naveen"]) // when we search key which doesn't exits, it will show the empty tring

	var accounts map[string]float64
	pf("%#v\n", accounts)

	//var m1 map[[]string]int // we cannot use slice as map type
	var m1 map[[5]string]int // But we can use array as map type
	_ = m1

	/* surNames["Naveen"] = "nimmala"  this is not possiable, when you declare a map and assigning the values is not possiable
	we need to initialize the map to add the key vales into that and we can use make function to create a map */

	people := map[string]float64{} // this is not a nil map, this is declared and initialize, if you initialize map you can assign the values

	people["Naveen"] = 172.2 // wen can assing the values
	people["sahithi"] = 176.6
	fmt.Println(people) //

	//another type to create map
	map1 := make(map[int]int)
	map1[1] = 10
	map1[2] = 20
	fmt.Println(map1)

	//another method to declare and initialize map
	balance := map[string]float64{
		"INR": 100.2,
		"USD": 150.5,
		"MYR": 17.2,
	}
	fmt.Println(balance)
	balance["INR"] = 1
	balance["SGD"] = 53
	fmt.Println(balance)
	fmt.Println(balance["EURO"]) /* this is not existing key, it will return zero
	zero 0 return means we dont really know weather EURO key has zero value or it doesnot exist we cannot confirm that
	to know that "COMMA OK IDIOM" notation we follow like below */

	val, ok := balance["EURO"]
	if ok {
		fmt.Println("The value is", val)
	} else {
		fmt.Println("THis keyh doesnt exits")
	}

	//loop

	for k, v := range balance {
		fmt.Printf("Key: %#v, Value %#v\n", k, v)
	}

	//delete

	delete(balance, "USD")
	fmt.Println(balance)

	m11 := map[int]int{1: 1, 2: 2, 3: 3}
	_ = m11

}
