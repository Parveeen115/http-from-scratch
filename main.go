package main

import (
	"github.com/Parveeen115/http-from-scratch/server"
)

func main() {
	//routes and shit
	server.HandleFunc("GET /", func(req server.Request) server.Response {
		return server.Response{StatusCode: 200, StatusText: "ok", Headers: []string{"content-type: text/plain"}, Body: "welcome to home"}
	})

	server.HandleFunc("GET /users/:id", func(req server.Request) server.Response {
		id := server.Param("id", req)
		return server.Response{StatusCode: 200, StatusText: "OK", Headers: []string{"Content-Type: text/plain"}, Body: "user id is " + id}
	})

	server.HandleFunc("GET /users/:id/posts", func(req server.Request) server.Response {
		id := server.Param("id", req)
		return server.Response{StatusCode: 200, StatusText: "OK", Headers: []string{"Content-Type: text/plain"}, Body: "posts for user " + id}
	})

	server.HandleFunc("POST /users", func(req server.Request) server.Response {
		return server.Response{StatusCode: 201, StatusText: "Created", Headers: []string{"Content-Type: text/plain"}, Body: "user created"}
	})

	server.HandleFunc("GET /search", func(req server.Request) server.Response {
		params := server.Query(req)
		return server.Response{StatusCode: 200, StatusText: "OK", Headers: []string{"Content-Type: text/plain"}, Body: "searching for " + params["q"]}
	})
	//start server
	server.Listen(8000)
}
