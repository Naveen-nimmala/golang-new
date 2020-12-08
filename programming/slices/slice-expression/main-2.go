package main

import (
	"fmt"
	"unsafe"
)

func main() {

	a := []int{1, 2, 3, 4, 5, 6}
	sl1, sl2 := a[:2], a[1:4]
	sl2[1] = 256
	fmt.Println(sl1, a, sl2) // 256 value will change for all slices due to same backing array and this is same for Arrays as well

	// if we want copy slice with different backing arrays We Should use append function
	cars := []string{"Ford", "Honda", "Audi", "Volvo"}
	newCars := []string{}

	newCars = append(newCars, cars[0:2]...) // we should give three dots if you want append the slice
	newCars[0] = "Benz"
	fmt.Println(cars)
	fmt.Println(newCars) // Here cars and newCars is using different backing arrays

	// length and capacity

	slc1 := []int{1, 2, 3, 4, 5, 6}
	slc2 := slc1[:2]
	fmt.Println(len(slc2), cap(slc2)) // length is 2 and capacity is 6

	slc2 = slc1[2:4]
	fmt.Println(len(slc2), cap(slc2))
	/*
	 Here slice slc2 length is 2 and capacity is 4
	 Reason : slc2 starting element is 2 , so it will start capacity count from 2nd element to last element.
	 in above case starting element is 2 and ending element is 6, so capacity is 4
	*/

	/* If we Check the pointer of both slices, it should be different because
	of position of elements are different but same backing array
	*/
	fmt.Printf("%p \n ", &slc1)
	fmt.Printf("%p \n ", &slc2)
	// Slice operation is cheper than arrays

	ar := [5]int{1, 2, 3, 4, 5}
	slcc := []int{1, 2, 3, 4, 5}

	fmt.Printf("Array size %d\n", unsafe.Sizeof(ar))
	fmt.Printf("Slice size %d\n", unsafe.Sizeof(slcc))

}
