package handler

import (
	"io"
	"log"
	"net/http"
)

type Proxy struct {
	client *http.Client
}

func NewProxy(client *http.Client) *Proxy {
	return &Proxy{
		client: client,
	}
}

func (p *Proxy) FwdUser(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method)
	log.Println(r.URL)
	log.Println(r.Header)

	if r.Method != http.MethodGet && r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var path string

	switch r.Method {
	case http.MethodGet:
		path = "http://localhost:8080/users"
	case http.MethodPost:
		path = "http://localhost:8080/addUser"
	}

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
