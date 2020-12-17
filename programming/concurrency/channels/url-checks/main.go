package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

func checkAndSavePage(url string, c chan string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		fmt.Printf("%s is down\n", url)
		c <- "Site is down"
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
			c <- "Site is up"
		}
	}
}
func main() {
	urls := []string{"https://www.google.com", "https://devopsismyworld.com", "https://golang.org", "https://facebook.com"}

	c := make(chan string)
	fmt.Println(time.Now())
	for _, url := range urls {
		go checkAndSavePage(url, c)
		fmt.Println(strings.Repeat("#", 20))
	}
	// fmt.Println(<-c) this channel will wait for only one out put comes, once it comes programe main will be exitied
	for i := 0; i < len(urls); i++ {
		fmt.Println(<-c) // this will wait for the 5 go routines to repond , once five routines will respond ,main will be exitied
	}

}
