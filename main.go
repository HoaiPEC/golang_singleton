package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

type APIClient struct {
	client *http.Client
}

var instance *APIClient
var once sync.Once

func GetInstance() *APIClient {
	once.Do(func() {
		instance = &APIClient{
			client: &http.Client{},
		}
	})

	return instance
}

func (c *APIClient) MakeGETRequest(url string) (*http.Response, error) {
	resp, err := c.client.Get(url)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func main() {
	client := GetInstance()

	response, err := client.MakeGETRequest("https://jsonplaceholder.typicode.com/posts/1")

	if err != nil {
		fmt.Println("Call API error, error: ", err)

		return
	}
	defer response.Body.Close()

	fmt.Println("Response body: ", response.Status)

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	fmt.Println("response body:", string(body))

}
