package handler

import (
	// "encoding/json"
	// "encoding/json"
	"io"
	"log"
	// "httpServer/internal/models"
	"net/http"
	// "os"
)
// currently, this doesnt stream the body. the http.NewRequest uses the intermediate body and passes it to the downstream
// route making a new http connection.
// this should not happen. imagine if body is 5gb. that is way to large to be stored in proxy and then passed down again
// proxy should not care what the body is. it should act as a pipe. client -> |proxy| -> backend
// basically, stream the body
func FwdUser(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method)
	log.Println(r.URL)
	log.Println(r.Header)
	log.Println("Raw body: ", r.Body)
	log.Printf("Raw body type: %T", r.Body)

	if( r.Method != http.MethodGet && r.Method != http.MethodPost) {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if(r.Method == http.MethodGet) {
		path := "http://localhost:8080/users"
		// fwd get request
		log.Println("fwding request to the server")
		
		newReq,err := http.NewRequest(r.Method, path, nil)
		if err != nil {
			http.Error(w, "Failed to create request", http.StatusInternalServerError)
			return
		}

		resp,err := http.DefaultClient.Do(newReq)
		
		if err != nil {
			http.Error(w, "Failed to forward request", http.StatusInternalServerError)
			return
		}
		
		defer resp.Body.Close()

		io.Copy(w,resp.Body)

	}

	if(r.Method == http.MethodPost){
		path := "http://localhost:8080/addUser"

		newReq,err := http.NewRequest(r.Method, path, r.Body)
		if err != nil {
			http.Error(w, "Failed to create request", http.StatusInternalServerError)
			return
		}

		resp,err := http.DefaultClient.Do(newReq)
		
		if err != nil {
			http.Error(w, "Failed to fwd request", http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		io.Copy(w, resp.Body)

	}
}
