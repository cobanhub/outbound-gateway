package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/cobanhub/lib/response"
	"github.com/cobanhub/lib/router"
)

func (a *API) HandleOutbound(ctx *router.Ctx) *response.JSONResponse {
	integrationName := ctx.Params("integration")
	coreRequest := make(map[string]interface{})
	response, err := a.interactor.HandleJson(integrationName, coreRequest, ctx.Request.Header)
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
