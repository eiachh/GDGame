package rest

import (
	"GDGame/commons"
	"GDGame/gamecontroller"
	"encoding/json"
	"fmt"
	"net/http"
)

func getBasicCommandHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		// Handle POST request
		var newItem commons.BasicCommand
		err := json.NewDecoder(r.Body).Decode(&newItem)
		if err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		response := gamecontroller.HandleBasicCommand(newItem)
		fmt.Printf("Received Item: %+v\n", newItem)
		json.NewEncoder(w).Encode(response)
		w.WriteHeader(http.StatusOK)
	}
}

func getInfoCommandHandler(w http.ResponseWriter, r *http.Request) {
	//TODO CHANGE

	switch r.Method {
	case http.MethodPost:
		// Handle POST request
		var newItem commons.BasicCommand
		err := json.NewDecoder(r.Body).Decode(&newItem)
		if err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		response := gamecontroller.HandleInfoCommand(newItem)
		fmt.Printf("Received Item: %+v\n", newItem)
		json.NewEncoder(w).Encode(response)
		w.WriteHeader(http.StatusOK)
	}
}

func getRegisterandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Expected POST", http.StatusBadRequest)
	}

	var registerToken commons.RegisterCommand
	err := json.NewDecoder(r.Body).Decode(&registerToken)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	fmt.Printf("Received registration request: %+v\n", registerToken)
	success, response, _ := gamecontroller.RegisterPlayer(registerToken)
	json.NewEncoder(w).Encode(response)
	if success {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

func StartRestApi(port int) {
	http.HandleFunc("/command/basic", getBasicCommandHandler)
	http.HandleFunc("/command/register", getRegisterandler)
	http.HandleFunc("/command/info", getInfoCommandHandler)

	// Start the HTTP server on port 8080
	address := fmt.Sprintf("127.0.0.1:%d", port)
	http.ListenAndServe(address, nil)
}
