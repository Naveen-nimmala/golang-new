package main

import "fmt"

func main() {

	nums := [...]int{1, 2, 3}
	nums1 := nums

	nums1[1] = 44
	fmt.Println(nums, nums1)
	fmt.Println(&nums, &nums1)

}
