package main

type Client struct {
	BaseURL    *url.URL
	UserAgent  string
	httpClient *http.Client
}

func main() {

}

// https://medium.com/@marcus.olsson/writing-a-go-client-for-your-restful-api-c193a2f4998c
