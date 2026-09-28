package main

import (
	"fmt"
	"log"
	"net/http"
	"reverseProxy/internal/handler"
	"time"
)

func main() {
	client := &http.Client{
		Timeout: 5 * time.Second,

		// defines the maximum number of connections per host.
		// if the limit is reached, the client will wait for a connection
		// to be available before making a new request
		Transport: &http.Transport{
			MaxConnsPerHost: 2,
		},
	}

	proxy := handler.NewProxy(client)

	http.HandleFunc("/", healthRoute)
	http.HandleFunc("/users", proxy.FwdUser)
	http.HandleFunc("/slow", proxy.SlowReqDemo)

	log.Println("proxy running on :8081")

	log.Fatal(http.ListenAndServe(":8081", nil))
}

func healthRoute(w http.ResponseWriter, r *http.Request) {

	fmt.Print("Method: ", r.Method)
	fmt.Print("Path : ", r.URL.Path)

	w.Write([]byte("proxy is up"))

}
