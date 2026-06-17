package main

import (
	"fmt"
	"strings"
)

type Response struct {
	StatusCode int
	StatusText string
	Headers    map[string]string
	Body       string
}

func BuildResponse(res Response) (string, error) {
	var lines []string
	lines = append(lines, fmt.Sprintf("HTTP/1.1 %d %s", res.StatusCode, res.StatusText))
	response := strings.Join(lines, "\r\n")
	fmt.Printf("response %s\n", response)
	return response, nil
}
