package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"./printers"
)

const baseURL string = "https://k-blog0130.herokuapp.com/api/v2/"

// Client represents api clients
type Client struct {
	baseURL string
}

// GetCategories gets categories and return it as a string.
func (c *Client) GetCategories() string {
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

// GetTags gets tags and return it as a stringnd return it as a string.
func (c *Client) GetTags() string {
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
	var categories string = client.GetCategories()
	printers.Display(categories, "CATEGORIES")

	var tags string = client.GetTags()
	printers.Display(tags, "TAGS")
}
