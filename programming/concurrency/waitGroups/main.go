package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

func fun1(wg *sync.WaitGroup) {
	fmt.Println("fun1 excution started")
	for i := 0; i < 3; i++ {
		fmt.Println("FROM fun1", i)
		time.Sleep(time.Second)
	}
	fmt.Println("fun1 excution stopped")
	wg.Done()
}

func fun2() {
	fmt.Println("fun2 excution started")
	for i := 0; i < 3; i++ {
		fmt.Println("FROM fun2", i)

	}
	fmt.Println("fun2 excution stopped")
}
func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	fmt.Println(strings.Repeat("#", 20))
	go fun1(&wg)
	fun2()
	wg.Wait()
	fmt.Println("Main function finished")

}
