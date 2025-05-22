package server

// func Start() {
// 	r := mux.NewRouter()
// 	middleware := middleware.NewMiddleware(middleware.MiddlewareOptions{
// 		Timeout: 5 * time.Second,
// 	})
// 	r.Use(middleware.RecoveryMiddleware)
// 	r.Handle("/outbound/{integration}", middleware.RecoveryMiddleware(http.HandlerFunc(controller.HandleOutbound))).Methods("POST")
// 	r.HandleFunc("/upload-config", controller.UploadConfigHandler).Methods("POST")

// 	srv := &http.Server{
// 		Addr:    ":8080",
// 		Handler: r,
// 	}
// 	log.Println("Starting server on :8080")
// 	srv.ListenAndServe()
// }
