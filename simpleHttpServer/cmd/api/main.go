package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"time"
	"httpServer/internal/handler"
)

func main() {
	portRange := 3
	startPort := 8080
	//run multiple instances of the server on 
	// different ports concurrently using goroutines
	for i := 0; i < portRange; i++ {
		go runServer(startPort + i)
	}

	select {}
}

func runServer(port int) {
	mux := http.NewServeMux()

	// register routes
	mux.HandleFunc("/", healthRoute(port))
	mux.HandleFunc("/users", handler.GetUserData)
	mux.HandleFunc("/addUser", handler.AddUserData)
	mux.HandleFunc("/slow", handler.SlowReqDemo)


	server := &http.Server{
		Addr:        fmt.Sprintf(":%d", port),
		Handler:     mux,
		IdleTimeout: 10 * time.Second,

		ConnState: func(conn net.Conn, state http.ConnState) {
			log.Println(conn.RemoteAddr(), state)
		},
	}

	log.Printf("Server running on :%d\n", port)

	err := server.ListenAndServe()
	if err != nil {
		log.Printf("Server on :%d stopped: %v\n", port, err)
	}
}

func healthRoute(port int) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Method:", r.Method)
		fmt.Println("Path:", r.URL.Path)

		fmt.Fprintf(w, "server is up on port %d\n", port)
	}
}