package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const baseURL string = "https://k-blog0130.herokuapp.com/api/v2/"

func main() {
	var categories string = GetCategories()
	fmt.Println(categories)

	var tags string = GetTags()
	fmt.Println(tags)
}

// GetCategories gets categories
func GetCategories() string {
	url := fmt.Sprintf(baseURL + "categories")
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

// GetTags gets tags
func GetTags() string {
	url := fmt.Sprintf(baseURL + "tags")
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

// https://medium.com/@marcus.olsson/writing-a-go-client-for-your-restful-api-c193a2f4998c
// https://www.scaledrone.com/blog/creating-an-api-client-in-go/
