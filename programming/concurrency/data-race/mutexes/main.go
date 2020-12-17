package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	const gr = 100

	var wg sync.WaitGroup
	var m sync.Mutex

	wg.Add(gr * 2)
	var n int = 0

	for i := 0; i < gr; i++ {
		go func() {
			time.Sleep(time.Second / 10)
			m.Lock()
			n++
			m.Unlock()
			wg.Done()
		}()
	}

	for i := 0; i < gr; i++ {
		go func() {
			time.Sleep(time.Second / 10)
			m.Lock()
			n--
			m.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Final value of n", n)
	/* this out put will change, because n value will try to chnage the both loops at a time it could be confilt
	we can detect this by using race detecter
	go run -race main.go

	we cam resolve this by using mutexes , using Mutexes we can really lock and unlock now the n output will get zero */

}
