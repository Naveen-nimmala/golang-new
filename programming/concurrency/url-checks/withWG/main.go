package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"
)

func checkAndSavePage(url string, wg *sync.WaitGroup) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		fmt.Printf("%s is down\n", url)
	} else {
		defer resp.Body.Close()
		fmt.Printf("%s --> Status code is: %v\n", url, resp.StatusCode)
		if resp.StatusCode == 200 {
			byteSlices, err := ioutil.ReadAll(resp.Body)
			file := strings.Split(url, "//")[1] // www.google.com
			file += ".txt"
			fmt.Printf("Writing resp body to %s\n", file)
			err = ioutil.WriteFile(file, byteSlices, 0644)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
	wg.Done()
}
func main() {
	urls := []string{"https://www.google.com", "https://devopsismyworld.com", "https://golang.org", "https://facebook.com"}
	fmt.Println(time.Now())
	var wg sync.WaitGroup
	wg.Add(len(urls))
	for _, url := range urls {
		go checkAndSavePage(url, &wg)
		fmt.Println(strings.Repeat("#", 20))
	}
	wg.Wait()
	fmt.Println(time.Now())
}
