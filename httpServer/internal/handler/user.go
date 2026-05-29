package handler

import (
	"encoding/json"
	"net/http"
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
	//intialise userData struct with dummy values
	response := UserData{
		Username: "Raju",
		Age:      22,
	}
	//set the headers to send a json response
	w.Header().Set("Content-Type", "application/json")

	// "w" is the http response steram
	// so anything written to "w" is sent to the client
	// newEncoder( ) creates an obj that knows how to convert go values to JSON
	encoder := json.NewEncoder(w)
	// ENcode() looks aat the map (the struct) and converts it to json
	encoder.Encode(response)
}
