package handler

import (
	// "encoding/json"
	"encoding/json"
	// "log"
	"net/http"
	"os"
)

//create a simple userData struct with fields
//`json:"username"` is a "struct tag" basically metadata attached to that field so basically, the json package knows that this should
// be encoded to "username " and not "Username"

// struct tags can also define more metadata like required fields, nullable fields, default values etc
type UserData struct {
	Username string `json:"username"`
	Age      int    `json:"age"`
}

func GetUserData(w http.ResponseWriter, r *http.Request) {

	data, err := os.ReadFile("internal/handler/users.json")
	if err != nil {
		panic(err)
	}

	var users []UserData

	err = json.Unmarshal(data, &users)
	if err != nil {
		panic(err)
	}
	

	//set the headers to send a json response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
	// "w" is the http response steram
	// so anything written to "w" is sent to the client

}
