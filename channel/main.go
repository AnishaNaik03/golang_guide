package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}
	c := make(chan string)
	for _, link := range links {
		go checkLink(link, c)
	}
	// fmt.Println(<-c)
	// fmt.Println(<-c)

	// for i := 0; i < len(links); i++ {
	// 	// fmt.Println(<-c)
	// }
	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	//time.Sleep(5 * time.Second)
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- "Might be down"
		return
	}
	fmt.Println(link, "is up!")
	c <- "its up!"
}
