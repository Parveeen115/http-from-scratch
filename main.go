package main

import (
	"github.com/Parveeen115/http-from-scratch/server"
)

func main() {
	server.HandleFunc("Get /", func(req server.Request) server.Response {
		return server.Response{StatusCode: 200, StatusText: "OK", Headers: []string{"Content-Type: text/plain"}, Body: "happy birthday"}
	})
	server.Listen(8000)
}
