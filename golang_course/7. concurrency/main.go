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
		"http://hotstar.com",
		"http://youtube.com",
	}

	for _, link := range links {
		go checkLink(link)
	}
	time.Sleep(1 * time.Second)
}

func checkLink(l string) {
	_, err := http.Get(l)

	if err != nil {
		fmt.Println(l, "is down :(")
	}
	fmt.Println(l, "is up :)")
}
