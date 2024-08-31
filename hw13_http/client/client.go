package client

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
)

func Run(url string) {

	// GET
	bodyGet, statusGet := MakeGetRequest(url)
	if statusGet != -1 {
		fmt.Println("Result GET body:", bodyGet, "Status:", statusGet)
	}

	// POST
	bodyPost, statusPost := MakeGPostRequest(url)
	if statusPost != -1 {
		fmt.Println("Result POST body:", bodyPost, "Status:", statusPost)
	}
}

func MakeGetRequest(url string) (string, int) {
	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, url, nil)
	if err != nil {
		fmt.Println(fmt.Errorf("GET NewRequest error: %w", err))
		os.Exit(1)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(fmt.Errorf("GET request error: %w", err))
		os.Exit(1)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(fmt.Errorf("GET error: %w", err))
		return "", -1
	}

	return string(body), resp.StatusCode
}

func MakeGPostRequest(url string) (string, int) {
	bodyreq := bytes.NewReader([]byte(`{"name": "Ivan"}`))
	req, err := http.NewRequestWithContext(context.Background(), http.MethodPost, url, bodyreq)
	if err != nil {
		fmt.Println(fmt.Errorf("GET NewRequest error: %w", err))
		os.Exit(1)
	}

	respPost, err := http.DefaultClient.Do(req)
	// respPost, err := http.Post(url, "", bytes.NewReader([]byte(`{"name": "Ivan"}`)))
	if err != nil {
		fmt.Println(fmt.Errorf("POST request error: %w", err))
		os.Exit(1)
	}
	defer respPost.Body.Close()

	bodyPost, err := io.ReadAll(respPost.Body)
	if err != nil {
		fmt.Println(fmt.Errorf("POST error: %w", err))
		return "", -1
	}

	return string(bodyPost), respPost.StatusCode
}
