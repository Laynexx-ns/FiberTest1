package jsonEncod

import (
	"FiberTest1/DataBase"
	"encoding/json"
	"fmt"
	"net/http"
)

type Message struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func JsonHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var msg Message
		err := json.NewDecoder(r.Body).Decode(&msg)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		DataBase.AddBaseJsonNameText(msg.Name, msg.Email)
		response := fmt.Sprintf("Received message from %s : %s", msg.Name, msg.Email)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"response": response})

	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}
