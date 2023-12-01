package main

import (
	"fmt"
	"io"
	"net/http"
)

const (
	okStatus = http.StatusOK
)

func fetchData(url string) ([]byte, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to make GET request to %s: %v", url, err)
	}
	defer response.Body.Close()

	if response.StatusCode != okStatus {
		return nil, fmt.Errorf("error: status code %d for URL: %s", response.StatusCode, url)
	}

	return io.ReadAll(response.Body)
}
