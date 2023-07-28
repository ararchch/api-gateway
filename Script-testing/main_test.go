package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

type Response struct {
	Product string `json:"Product"`
	Quotient string `json:"Quotient"`
    Sum string `json:"Sum"`
}

func TestMultiplyEndpoint(t *testing.T) {
	testEndpoint(t, "http://127.0.0.1:8080/multiply", "18", "multiply")
}

func TestDivideEndpoint(t *testing.T) {
	testEndpoint(t, "http://127.0.0.1:8080/divide", "2", "divide")
}

func TestAddEndpoint(t *testing.T) {
	testEndpoint(t, "http://127.0.0.1:8080/add", "9", "add")
}

func testEndpoint(t *testing.T, url string, expectedProduct string, operation string) {
	method := "POST"
	payload := strings.NewReader(`{ "FirstNum": "6", "SecondNum": "3" }`)

	client := &http.Client{}
	req, _ := http.NewRequest(method, url, payload)
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("Failed to read response body: %v", err)
	}

	var resp Response
	err = json.Unmarshal(body, &resp)
	if err != nil {
		t.Fatalf("Failed to unmarshal response body: %v", err)
	}

	var result string
    if operation == "multiply" {
        result = resp.Product
    } else if operation == "divide" {
        result = resp.Quotient
    } else if operation == "add" {
        result = resp.Sum
    }

	if result != expectedProduct {
		t.Fatalf("Expected '%v', got '%v'", expectedProduct, resp.Product)
	}
}
