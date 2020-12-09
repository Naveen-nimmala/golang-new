package main

import "fmt"

func main() {
	friends := map[string]int{"naveen": 28, "sahithi": 26, "abhi": 25}
	neighbors := friends
	friends["naveen"] = 30
	fmt.Println(neighbors) // whe we assing one map to another map, it just referenc it, it doesnt really copy the entire map
	// in the abouve exaple, friends assigned to neighbors, when i change friends key value, neighbors also chnage
	// its not a copy, its just refence it

	// copy or clone map --> we might need to use for loop for that

	people := make(map[string]int)

	for ky, val := range friends {
		people[ky] = val
	}
	fmt.Println(people)
	people["Kumar"] = 60
	fmt.Printf("%#v -----------------  %v#\n", people, friends)
}
