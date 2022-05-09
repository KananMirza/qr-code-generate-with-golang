package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"qr-code-generate-with-golang/pkg/handlers"
	"qr-code-generate-with-golang/pkg/helpers"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", handlers.MainHandlers).Methods("GET")
	r.HandleFunc("/", handlers.MainHandlersPost).Methods("POST")

	r.HandleFunc("/qr-code", handlers.QrCodeHandlers).Methods("GET")

	fmt.Println("Server starting...")
	helpers.CheckError(http.ListenAndServe(":9000", r))
}
