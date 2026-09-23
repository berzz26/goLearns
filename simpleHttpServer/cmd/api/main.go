package main

import (
	"fmt"
	"httpServer/internal/handler"
	"log"
	"net"
	"net/http"
	"time"
)

func main() {
	//register routes

	http.HandleFunc("/", healthRoute)
	http.HandleFunc("/users", handler.GetUserData)
	http.HandleFunc("/addUser", handler.AddUserData)
	http.HandleFunc("/slow", handler.SlowReqDemo)

	log.Println("Server running on :8080")

	// ConnState specifies an optional callback function that is
	// called when a client connection changes state
	server := &http.Server{
		Addr: ":8080",
		IdleTimeout: 10 * time.Second,

		ConnState: func(conn net.Conn, state http.ConnState) {
			log.Println(conn.RemoteAddr(), state)
		},
	}
	server.ListenAndServe()
}

func healthRoute(w http.ResponseWriter, r *http.Request) {

	fmt.Print("Method: ", r.Method)
	fmt.Print("Path : ", r.URL.Path)

	w.Write([]byte("server is up"))

}
