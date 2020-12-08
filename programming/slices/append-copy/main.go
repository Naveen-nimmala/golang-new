package main

import "fmt"

func main() {

	// append to exiting slice
	nums := []int{1, 2}
	nums = append(nums, 20)
	fmt.Println(nums)

	nums = append(nums, 10, 20, 30, 40)
	fmt.Println(nums)

	// append  new slice

	n := []int{100, 200}
	nums = append(nums, n...)
	fmt.Println(nums)

	// copy slice from src to dest

	src := []int{10, 20, 30}
	dst := make([]int, len(src))
	nn := copy(dst, src)
	fmt.Println(src, dst, nn)

}
