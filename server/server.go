package server

import (
	"fmt"
	"net/http"
)

const PORT = ":8080"

type Server interface {
	Serve(filePath string) error
}

type NormalServer struct{}

func (ns *NormalServer) Serve(filePath string) error {
	http.HandleFunc("/", serveFileFromPath(filePath))

	err := http.ListenAndServe(PORT, nil)
	if err != nil {
		return err
	}
	fmt.Printf("Listening on port %s\n", PORT)
	return nil
}

func serveFileFromPath(filePath string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, filePath)
	}
}
