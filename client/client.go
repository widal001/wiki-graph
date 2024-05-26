package client

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type Client interface {
	Get(baseUrl string, queryParams map[string]string) ([]byte, error)
}

type HttpClient struct {
	client *http.Client
}

// Get a new Wikipedia API client
func New() *HttpClient {
	return &HttpClient{client: &http.Client{}}
}

func (c HttpClient) Get(baseURL string, queryParams map[string]string) ([]byte, error) {

	// Prepare the base request
	request, err := http.NewRequest("GET", baseURL, nil)
	if err != nil {
		err = fmt.Errorf("error preparing the request: %s", err)
		return []byte{}, err
	}

	// Add the query params to the request
	params := url.Values{}
	for param, value := range queryParams {
		params.Add(param, value)
	}
	request.URL.RawQuery = params.Encode()

	// Make the GET request
	resp, err := c.client.Do(request)
	if err != nil {
		err = fmt.Errorf("error sending request: %s", err)
		return []byte{}, err
	}
	defer resp.Body.Close()

	// Read and return the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		err = fmt.Errorf("error reading response body: %s", err)
		return []byte{}, err
	}
	return body, nil
}
