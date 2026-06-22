package server

import "strings"

type Handler struct {
	Route  string
	Handle func(Request) Response
}

var batch []Handler

func HandleFunc(route string, handle func(Request) Response) {
	batch = append(batch, Handler{route, handle})
}

func CallFunc(req Request) Response {
	for _, v := range batch {
		routeMethod := strings.Split(v.Route, " ")
		if routeMethod[1] != req.Path {
			continue
		}
		return v.Handle(req)
	}
	return Response{}
}
