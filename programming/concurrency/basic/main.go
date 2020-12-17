package main

import (
	"fmt"
	"runtime"
	"time"
)

func fun1() {
	fmt.Println("fun1 excution started")
	for i := 0; i < 3; i++ {
		fmt.Println("FROM fun1", i)
	}
	fmt.Println("fun1 excution stopped")
}

func fun2() {
	fmt.Println("fun2 excution started")
	for i := 0; i < 3; i++ {
		fmt.Println("FROM fun2", i)
	}
	fmt.Println("fun2 excution stopped")
}
func main() {
	// NumCPU returns the number of logical CPUs or cores usable by the current process.
	fmt.Println("No. of CPUs:", runtime.NumCPU()) // => No. of CPUs: 4

	// NumGoroutine returns the number of goroutines that currently exists.
	fmt.Println("No. of Goroutines:", runtime.NumGoroutine()) // => No. of Goroutines: 1

	// GOOS and GOARCH are environment variables
	fmt.Println("OS:", runtime.GOOS)     // => OS: linux
	fmt.Println("Arch:", runtime.GOARCH) // => Arch: amd64

	//  GOMAXPROCS sets the maximum number of CPUs that can be executing simultaneously and returns
	//  the previous setting.
	fmt.Println("GOMAXPROCS:", runtime.GOMAXPROCS(0)) // => GOMAXPROCS: 4
	// If n < 1, it does not change the current setting.

	go fun1()
	fun2()
	time.Sleep(time.Second * 2)

}
