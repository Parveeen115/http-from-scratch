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
	req.Param = make(map[string]string)
	for _, v := range batch {
		routeMethod := strings.Split(v.Route, " ")
		if routeMethod[0] != req.Method {
			continue
		}
		methods := strings.Split(routeMethod[1], "/")
		reqMethod := strings.Split(req.Path, "/")
		if len(methods) != len(reqMethod) {
			continue
		}
		check := true
		for i := 0; i < len(methods); i++ {
			param := strings.Contains(methods[i], ":")
			if param {
				check = true
				req.Param[methods[i][1:]] = reqMethod[i]
				continue
			}
			if methods[i] != reqMethod[i] {
				check = false
				break
			}
		}
		if check {
			return v.Handle(req)
		} else {
			continue
		}
	}
	return Response{StatusCode: 404, StatusText: "not_found", Headers: []string{"content-type: text/plain"}, Body: "NOT FOUND"}
}
