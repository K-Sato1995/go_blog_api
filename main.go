package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const baseURL string = "https://k-blog0130.herokuapp.com/api/v2/"

func main() {
	GetCategories()
}

func GetCategories() {
	url := fmt.Sprintf(baseURL + "categories")
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(string(body))
}

// https://medium.com/@marcus.olsson/writing-a-go-client-for-your-restful-api-c193a2f4998c
// https://www.scaledrone.com/blog/creating-an-api-client-in-go/
