package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleOutbound(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	integrationName := vars["integration"]
	var coreRequest map[string]interface{}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(coreResponse)
	log.Println("Success outbound for integration:", integrationName)
}
func UploadConfigHandler(w http.ResponseWriter, r *http.Request) {
	// Retrieve the YAML file from the request
	file, _, err := r.FormFile("config_file")
	if err != nil {
		http.Error(w, "Failed to read file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Return success response
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Config uploaded successfully"))
}
