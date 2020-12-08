package main

import "fmt"

func main() {

	var nums []int

	fmt.Printf("%#v \n", nums)
	fmt.Printf("Length: %d Capacity %d \n", len(nums), cap(nums)) // Len 0 , cap 0
	/* when we append new elements to slice it will increse the
	capacity like 2 --> 4 --> 8 --> 16 --> 32 etc. this operation
	might avoid the slice copy everytime when we add new elements */
	nums = append(nums, 1, 2)
	fmt.Printf("Length: %d Capacity %d \n", len(nums), cap(nums)) // Len 2 , cap 2
	nums = append(nums, 30)
	fmt.Printf("Length: %d Capacity %d \n", len(nums), cap(nums)) // Len 3 , cap 4
	nums = append(nums, 10)
	fmt.Printf("Length: %d Capacity %d \n", len(nums), cap(nums)) // Len 4 , cap 4
	nums = append(nums, 100)
	fmt.Printf("Length: %d Capacity %d \n", len(nums), cap(nums)) // Len 5 , cap 8
	nums = append(nums, 100)
	fmt.Printf("Length: %d Capacity %d \n", len(nums), cap(nums)) // Len 5 , cap 8
	nums = append(nums, 22, 33, 44, 55, 66)
	fmt.Printf("Length: %d Capacity %d \n", len(nums), cap(nums)) // Len 11 , cap 16

	values := []int{1, 2, 3, 4, 5, 6, 7, 8}

	values = append(values[0:2], 236, 256)
	fmt.Println(values, len(values), cap(values))
	/* values : 1,2,236,256 , lenght: 4 capacity: 8
	in backing array 5,6,7,8 values still will be exists
	if you want to access those elements you should use the starting index lenght is 0 to 4. like values[1:5], values[3:6] etc.
	because values varible length is 4. if you try to use index value directly like values[6] it will though error */

	// fmt.Println(values[5]) // we cannot use this, because Value varible doesnt know the index 5 value

	fmt.Println(values[2:6]) // this will print the output because staring index value varible knows

}
