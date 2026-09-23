package main
import (
	"fmt"
	"reverseProxy/internal/handler"
	"log"
	"net/http"
)

func main(){

	http.HandleFunc("/", healthRoute)
	http.HandleFunc("/users", handler.FwdUser)

	log.Println("proxy running on :8081")

	log.Fatal(http.ListenAndServe(":8081", nil))
}

func healthRoute(w http.ResponseWriter, r *http.Request) {

	fmt.Print("Method: ", r.Method)
	fmt.Print("Path : ", r.URL.Path)

	w.Write([]byte("proxy is up"))

}



