package httpClient

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const server = "https://jsonplaceholder.typicode.com"

func FetchList(what string, data interface{}) error {
	url := fmt.Sprintf("%s/posts", server)
	return fetch(url, data)
}

func fetch(url string, data interface{}) error {
	response, err := doFetch(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	return responseToJson(response, data)
}

func doFetch(url string) (*http.Response, error) {
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("Failed to create request: %s", err)
	}
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, fmt.Errorf("Failed to execute request: %s", err)
	}
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Request return an error code: %s", err)
	}
	return response, nil
}

func responseToJson(response *http.Response, data interface{}) error {
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return fmt.Errorf("Failed to read response %s", err)
	}
	if err := json.Unmarshal(body, data); err != nil {
		return fmt.Errorf("Failed to unmarshal response %s", err)
	}
	return nil
}
