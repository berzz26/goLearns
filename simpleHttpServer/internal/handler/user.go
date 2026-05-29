package handler

import (
	// "encoding/json"
	"encoding/json"
	// "log"
	"httpServer/internal/models"
	"net/http"
	"os"
)

//create a simple userData struct with fields
//`json:"username"` is a "struct tag" basically metadata attached to that field so basically, the json package knows that this should
// be encoded to "username " and not "Username"

// struct tags can also define more metadata like required fields, nullable fields, default values etc

func GetUserData(w http.ResponseWriter, r *http.Request) {

	data, err := os.ReadFile("internal/handler/users.json")
	if err != nil {
		panic(err)
	}

	var users []models.UserData

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
func AddUserData(w http.ResponseWriter, r *http.Request) {
	var user models.UserData
	var users []models.UserData

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	fileData, err := os.ReadFile("internal/handler/users.json")
	if err == nil {
		err = json.Unmarshal(fileData, &users)
		if err != nil {
			http.Error(w, "Failed to parse users file", http.StatusInternalServerError)
			return
		}
	}

	users = append(users, user)

	updatedData, err := json.MarshalIndent(users, "", "  ")
	if err != nil {
		http.Error(w, "Failed to encode users", http.StatusInternalServerError)
		return
	}

	err = os.WriteFile("internal/handler/users.json", updatedData, 0644)
	if err != nil {
		http.Error(w, "Failed to save users", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(user)
}