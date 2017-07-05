package main

import (
	"fmt"
	"net/http"
)

var urls = []string{
	"https://www.google.com/",
	"https://golang.org/",
	"https://blog.golang.org/",
}

func main() {
	for _, url := range urls {
		resp, err := http.Head(url)
		if err != nil {
			fmt.Println("Error:", url, err)
		}
		fmt.Println(url, ": ", resp.Status)
	}
}
