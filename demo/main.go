package main

import (
	"bytes"
	"log"
	"log/slog"
	"net/http"
)

func main() {
	mux := http.NewServeMux()


	mux.HandleFunc("/", handleRoot)
	mux.HandleFunc("/goodbye", handleGoodbye)
	mux.HandleFunc("/hello/", handleParameterized)
	mux.HandleFunc("/responses/{user}/hello/", handleUserResponsesHello)

	log.Fatal(http.ListenAndServe(":8080", mux))
}

func handleRoot(w http.ResponseWriter, _ *http.Request) {
	_, err := w.Write([]byte("Welcome to our homepage"))
	if err != nil {
		slog.Error("Failed to write response", "error", err)
		return
	}
}

func handleGoodbye (w http.ResponseWriter, _ *http.Request) {
	_, err := w.Write([]byte("Goodbye!"))
	if err != nil {
		slog.Error("Failed to write response", "error", err)
		return
	}
}

func handleParameterized(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	userList := params["user"]

	username := "User"
	if len(userList) > 0 {
		username = userList[0]
	}

	var output bytes.Buffer
	output.WriteString("Hello, ")
	output.WriteString(username)
	output.WriteString("!\n")
	_, err := w.Write(output.Bytes())

	if err != nil {
		slog.Error("Error writing response body", "err", err)
		return
	}
}

func handleUserResponsesHello(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Not implemented yet", http.StatusInternalServerError)
}