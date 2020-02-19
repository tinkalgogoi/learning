package main

import (
	"flag"
	"fmt"
	"net/http"
)

var (
	url string
)

type HttpClient interface {
	Get(string) (*http.Response, error)
}

func init() {
	flag.StringVar(&url, "url", "http://google.com", "Which URL do we want to parse?")
	flag.Parse()
}

func main() {
	client := &http.Client{}
	err := send(client, url)

	if err != nil {
		panic(err)
	}
}

func send(client HttpClient, link string) error {
	response, err := client.Get(link)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(response)
	return nil
}
