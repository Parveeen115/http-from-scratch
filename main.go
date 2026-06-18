package main

import (
	"fmt"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Server is listening on port 8000")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error :", err)
			continue
		}

		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()

	buffer := make([]byte, 1024)

	for {
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("error", err)
			return
		}

		fmt.Printf("Recived: %s\n", buffer[:n])
		req := ParseRequest(string(buffer[:n]))

		fmt.Printf("Parsed request: %v\n", req)

		res := Response{StatusCode: 200, StatusText: "OK", Headers: []string{"Content-Type: text/plain"}, Body: "happy birthday"}

		response := BuildResponse(res)
		conn.Write([]byte(response))
	}
}
