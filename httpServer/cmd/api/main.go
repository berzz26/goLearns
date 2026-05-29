package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", healthRoute)
	log.Println("Server running on :8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
func healthRoute(w http.ResponseWriter, r *http.Request) {

	fmt.Print("Method: ", r.Method)
	fmt.Print("Path : ", r.URL.Path)

	w.Write([]byte("HEY"))

}
