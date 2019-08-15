package httpClient

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func FetchResources(resource Resource, data interface{}) error {
	url := makeUrl(resource)
	return fetch(url, data)
}

func FetchResource(resource Resource, resourceId int, data interface{}) error {
	url := makeUrl(resource, resourceId)
	return fetch(url, data)
}

func FetchSubResources(resource Resource, resourceId int, subResource Resource, data interface{}) error {
	url := makeUrl(resource, resourceId, subResource)
	return fetch(url, data)
}

func FetchSubResource(resource Resource, resourceId int, subResource Resource, subResourceId int, data interface{}) error {
	url := makeUrl(resource, resourceId, subResource, subResourceId)
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
		return nil, fmt.Errorf("Failed to create request: %v", err)
	}
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, fmt.Errorf("Failed to execute request: %v", err)
	}
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Request return an error code: %d", response.StatusCode)
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

func makeUrl(params ...interface{}) string {
	var url strings.Builder
	url.WriteString(server)
	for _, str := range params {
		url.WriteRune('/')
		url.WriteString(fmt.Sprint(str))
	}
	return url.String()
}

const server = "https://jsonplaceholder.typicode.com"

type Resource string
