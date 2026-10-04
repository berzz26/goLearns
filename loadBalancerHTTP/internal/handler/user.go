package handler

import (
	"io"
	"log"
	"net/http"
	"sync/atomic"
)

type Proxy struct {
	client *http.Client
}

type Server struct {
	Address string
	Weight  int
}

func NewProxy(client *http.Client) *Proxy {
	return &Proxy{
		client: client,
	}
}

var reqCount int64

func (p *Proxy) FwdUser(w http.ResponseWriter, r *http.Request) {
	// rr balancing -select the server based on the
	// request count

	servers := []Server{
		{Address: "http://localhost:8080", Weight: 1},
		{Address: "http://localhost:8081", Weight: 2},
		{Address: "http://localhost:8082", Weight: 2},
	}

	log.Println(r.Method)
	log.Println(r.URL)
	log.Println(r.Header)

	if r.Method != http.MethodGet && r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var path string
	//atmoic increment of reqCount to ensure thread safety
	// during concurrent requests.

	//weigheted RR
	serverIndex := getWeightedServer(servers)

	//simple RR
	// serverIndex := getNextServer(servers)

	switch r.Method {
	case http.MethodGet:
		path = servers[serverIndex].Address + "/users"
	case http.MethodPost:
		path = servers[serverIndex].Address + "/addUser"
	}

	log.Println("selected server: ", servers[serverIndex])

	newReq, err := http.NewRequest(r.Method, path, r.Body)
	if err != nil {
		http.Error(w, "Failed to create request", http.StatusInternalServerError)
		return
	}

	// Forward request headers.
	for key, values := range r.Header {
		for _, value := range values {
			newReq.Header.Add(key, value)
		}
	}

	resp, err := p.client.Do(newReq)
	if err != nil {
		log.Println("failed to forward request:", err)
		http.Error(w, "Failed to forward request", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Forward response headers.
	for key, values := range resp.Header {
		for _, value := range values {
			w.Header().Add(key, value)
		}
	}

	// Forward status.
	w.WriteHeader(resp.StatusCode)

	// Stream response body back to client.
	_, err = io.Copy(w, resp.Body)
	if err != nil {
		log.Println("failed to stream response:", err)
	}

	atomic.AddInt64(&reqCount, 1)
}

func (p *Proxy) SlowReqDemo(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method)
	log.Println(r.URL)
	log.Println(r.Header)

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	path := "http://localhost:8080/slow"

	newReq, err := http.NewRequest(r.Method, path, r.Body)
	if err != nil {
		http.Error(w, "Failed to create request", http.StatusInternalServerError)
		return
	}

	// Forward request headers.
	for key, values := range r.Header {
		for _, value := range values {
			newReq.Header.Add(key, value)
		}
	}

	resp, err := p.client.Do(newReq)
	if err != nil {
		log.Println("forwarding error: ", err)
		http.Error(w, "Failed to forward request", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Forward response headers.
	for key, values := range resp.Header {
		for _, value := range values {
			w.Header().Add(key, value)
		}
	}

	w.WriteHeader(resp.StatusCode)

	// Stream response body back to client.
	_, err = io.Copy(w, resp.Body)
	if err != nil {
		log.Println("failed to stream response:", err)
	}
}

func getWeightedServer(servers []Server) int {
	totalWeight := 0

	for _, server := range servers {
		totalWeight += server.Weight
	}

	count := atomic.AddInt64(&reqCount, 1) - 1
	slot := int(count % int64(totalWeight))

	for i, server := range servers {
		if slot < server.Weight {
			return i
		}

		slot -= server.Weight
	}

	return 0
}
func getNextServer(servers []Server) int {
	count := atomic.AddInt64(&reqCount, 1) - 1
	return int(count % int64(len(servers)))
}