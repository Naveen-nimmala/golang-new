package main

import "fmt"

func main() {

	title, year, director := "Tenet", 2020, "Nelon"
	title1, year1, director1 := "Avengers", 2012, "Michael"
	fmt.Println("Movie 1: ", title, year, director)
	fmt.Println("Movie 2: ", title1, year1, director1)

	type movie struct {
		title    string
		year     int
		director string
	}
	type student struct {
		name, clgName string
		year          int
	}

	mymovie1 := movie{
		title:    "Tenet",
		year:     2020,
		director: "Nelon",
	}
	fmt.Println(mymovie1)       // print all the fields in the Struct
	fmt.Println(mymovie1.title) // print only title name
	mymovie1.year = 2019        // chaning the year
	fmt.Println(mymovie1.year)  // printing the updated year

	mymovie2 := movie{
		title:    "Avengers",
		year:     2012,
		director: "Naveen",
	}

	fmt.Printf("%#v\n", mymovie2) // it will print the assigned values with nil values

	// comparing two structs
	fmt.Println(mymovie1 == mymovie2) // it will print false

	//copy of structs
	mymovie3 := mymovie2 // when you assing the one struct to another struct, it will copy the struct
	mymovie3.title = "Robo"
	fmt.Println(mymovie3, mymovie2)

}
