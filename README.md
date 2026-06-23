# http-from-scratch

Building HTTP/1.1 from raw TCP in Go. No net/http, no frameworks, no external packages.

## What's implemented

- Raw TCP listener and accept loop
- HTTP request parser (request line, headers, body)
- Router with method + path matching
- Route parameters (`/users/:id` → `Param("id", req)`)
- Query string parsing (`/search?q=golang` → `Query(req)`)
- HTTP response builder

## Usage

```go
server.HandleFunc("GET /users/:id", func(req server.Request) server.Response {
    id := server.Param("id", req)
    return server.Response{StatusCode: 200, StatusText: "OK", Body: "user: " + id}
})

server.Listen(8000)
```

## Project structure

```
main.go          — registers routes and starts the server
server/
  parser.go      — request parsing and Query()
  router.go      — HandleFunc, CallFunc, route matching
  response.go    — response builder
  server.go      — TCP listener
```
