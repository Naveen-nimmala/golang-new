package main

import "fmt"

func main() {

	relatives := []string{"NN", "Naveen", "Nimmala", "Sahi", "India", "abhi"}
	friends := []string{"friend1", "Sahi", "abhi"}

outer:
	for _, name := range relatives {
		for _, nme := range friends {
			if name == nme {
				fmt.Printf("Found %q as your relative and friend\n", name)
				break outer
			}
		}
	}
	fmt.Println("Go to nex step")
}
