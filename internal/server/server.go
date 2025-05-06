package server

import (
	"log"
	"net/http"

	"github.com/cobanhub/outbound-gateway/internal/controller"
	"github.com/cobanhub/outbound-gateway/internal/middleware"

	"github.com/gorilla/mux"
)

func Start() {
	r := mux.NewRouter()
	r.HandleFunc("/outbound/{integration}", middleware.RecoveryMiddleware(controller.HandleOutbound())).Methods("POST")
	r.HandleFunc("/upload-config", controller.UploadConfigHandler).Methods("POST")

	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	log.Println("Starting server on :8080")
	srv.ListenAndServe()
}
