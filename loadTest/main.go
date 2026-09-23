package main

import (
	"log"
	"net/http"
	"sync"
	"time"
)

func main() {
	client := http.Client{
		//comment out to keep disable connection reusing and allowing the client to make a new connection
		// for ever request.
		// Transport: &http.Transport{
		// 	DisableKeepAlives: true,
		// },
	}
	
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			makeRequest(client)
		}()
	}
	wg.Wait()

	time.Sleep(2 * time.Second)

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			makeRequest(client)
		}()
	}
	wg.Wait()
}

func makeRequest(client http.Client) {
	resp, err := client.Get("http://localhost:8080/slow")
	if err != nil {
		log.Println("Error making request:", err)
		return
	}
	defer resp.Body.Close()

	log.Println("Response status:", resp.Status)
}
