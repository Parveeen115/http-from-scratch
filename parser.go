package main

import (
	"fmt"
	"strings"
)

type Request struct {
	Method  string
	Path    string
	Version string
	Headers map[string]string
	Body    string
}

func ParseRequest(raw string) Request {
	var req Request
	sep := strings.Split(raw, "\r\n")
	fmt.Printf("seperate %q\n", sep)
	//method n shit
	top := strings.Split(sep[0], " ")
	fmt.Printf("top seperate %q\n", top)
	req.Method = top[0]
	req.Path = top[1]
	req.Version = top[2]
	//headers
	req.Headers = make(map[string]string)
	for i, line := range sep[1:] {
		if line == "" {
			req.Body = sep[i+2]
			break
		}
		parts := strings.SplitN(line, ": ", 2)
		if len(parts) == 2 {
			req.Headers[parts[0]] = parts[1]
		}
	}
	fmt.Printf("header : %q\n", req)
	return req
}
