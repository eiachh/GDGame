package rest

import (
	"GDGame/gamecontroller"
	"encoding/json"
	"fmt"
	"net/http"
)

/*
	 type Item struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}

	func getItemsHandler(w http.ResponseWriter, r *http.Request) {
		// Convert items to JSON and send it as the response
		json.NewEncoder(w).Encode(items)
	}
*/

func getBasicCommandHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		fmt.Printf("Got GET request")
		json.NewEncoder(w).Encode("ASDDDSDSD222")
	case http.MethodPost:
		// Handle POST request
		var newItem gamecontroller.BasicCommand
		err := json.NewDecoder(r.Body).Decode(&newItem)
		if err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		gamecontroller.HandleBasicCommand(newItem)
		// Print the received item to the server's console
		fmt.Printf("Received Item: %+v\n", newItem)
		w.WriteHeader(http.StatusOK)
	}
}

func getRegisterandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Expected POST", http.StatusBadRequest)
	}

	var registerToken gamecontroller.RegisterCommand
	err := json.NewDecoder(r.Body).Decode(&registerToken)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	gamecontroller.RegisterPlayer(registerToken)
	// Print the received item to the server's console
	fmt.Printf("Received registration request: %+v\n", registerToken)
	w.WriteHeader(http.StatusOK)

}

func StartRestApi(port int) {
	http.HandleFunc("/command/basic", getBasicCommandHandler)
	http.HandleFunc("/command/register", getRegisterandler)

	// Start the HTTP server on port 8080
	address := fmt.Sprintf("127.0.0.1:%d", port)
	http.ListenAndServe(address, nil)
}
