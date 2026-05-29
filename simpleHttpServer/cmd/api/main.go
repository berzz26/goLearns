package main

import (
	"fmt"
	"httpServer/internal/handler"
	"log"
	"net/http"
)

func main() {
	//register routes
	http.HandleFunc("/", healthRoute)
	http.HandleFunc("/users", handler.GetUserData)
	http.HandleFunc("/addUser", handler.AddUserData)

	log.Println("Server running on :8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func healthRoute(w http.ResponseWriter, r *http.Request) {

	fmt.Print("Method: ", r.Method)
	fmt.Print("Path : ", r.URL.Path)

	w.Write([]byte("HEY"))

}
