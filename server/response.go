package server

import (
	"fmt"
	"strings"
)

type Response struct {
	StatusCode int
	StatusText string
	Headers    []string
	Body       string
}

func BuildResponse(res Response) string {
	var lines []string
	lines = append(lines, fmt.Sprintf("HTTP/1.1 %d %s", res.StatusCode, res.StatusText))
	for _, v := range res.Headers {
		lines = append(lines, v)
	}
	lines = append(lines, fmt.Sprintf("Content-length: %d", len(res.Body)))
	lines = append(lines, "")
	lines = append(lines, res.Body)
	response := strings.Join(lines, "\r\n")
	fmt.Printf("response %s\n", response)
	return response
}
