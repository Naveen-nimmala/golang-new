package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	const gr = 100

	var wg sync.WaitGroup

	wg.Add(gr * 2)
	var n int = 0

	for i := 0; i < gr; i++ {
		go func() {
			time.Sleep(time.Second / 10)
			n++
			wg.Done()
		}()
	}

	for i := 0; i < gr; i++ {
		go func() {
			time.Sleep(time.Second / 10)
			n--
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Final value of n", n)
	/* this out put will n increase upto 100 and it has to decrese to 0 but here the output comes 0,1,2,3 it keep changes,
	because n value will try to change the both loops at a time it could be confilt
	we can detect this by using race detecter
	go run -race main.go */

}
