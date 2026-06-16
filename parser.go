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

func ParseRequest(raw string) (Request, error) {
	var req Request
	sep := strings.Split(raw, "\r\n")
	fmt.Printf("seperate %q\n", sep)
	top := strings.Split(sep[0], " ")
	fmt.Printf("top seperate %q\n", top)
	req.Method = top[0]
	req.Path = top[1]
	req.Version = top[2]
	return req, nil
}
