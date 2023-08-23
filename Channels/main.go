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
		"http://amazon.com",
		"http://flipkart.com",
		"http://golang.org",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link,c) 

	}
	// for {
	// 	go checkLink(<-c,c)
	// }

	// for l := range c{
	// 	fmt.Println(l,c)
	// }

	for l:= range c{
		go func(link string){
			time.Sleep(3*time.Second)
			checkLink(link,c)
		}(l)
	}

}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link
		return
	}
	fmt.Println(link, "is Up and Running!")
	c <- link
}
