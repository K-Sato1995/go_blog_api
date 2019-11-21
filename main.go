package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const baseURL string = "https://k-blog0130.herokuapp.com/api/v2/"

// Client represents api clients
type Client struct {
	baseURL string
}

func (c *Client) getCategories() string {
	url := fmt.Sprintf(c.baseURL + "categories")

	res, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}

	return string(body)
}

func (c *Client) getTags() string {
	url := fmt.Sprintf(c.baseURL + "tags")
	res, err := http.Get(url)

	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Fatalln(err)
	}

	return string(body)
}

func main() {
	client := Client{baseURL: "https://k-blog0130.herokuapp.com/api/v2/"}
	var categories string = client.getCategories()
	fmt.Println(categories)

	var tags string = client.getTags()
	fmt.Println(tags)

}
