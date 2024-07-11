package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/sethvargo/go-envconfig"
)

type Config struct {
	ServerHost string `env:"ECHO_SERVER_HOST,default=0.0.0.0"`
	ServerPort int    `env:"ECHO_SERVER_PORT,default=18080"`
}

type RequestInfo struct {
	Method      string              `json:"method"`
	URL         string              `json:"url"`
	Headers     map[string][]string `json:"headers"`
	QueryParams map[string][]string `json:"query_params"`
	Body        string              `json:"body"`
}

func echoHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	info := RequestInfo{
		Method:      r.Method,
		URL:         r.URL.Path,
		Headers:     r.Header,
		QueryParams: r.URL.Query(),
		Body:        string(body),
	}

	jsonResponse, err := json.MarshalIndent(info, "", "  ")
	if err != nil {
		http.Error(w, "Error creating JSON response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func main() {
	var config Config
	ctx := context.Background()

	if err := envconfig.Process(ctx, &config); err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", echoHandler)

	serverAddr := fmt.Sprintf("%s:%d", config.ServerHost, config.ServerPort)
	fmt.Printf("Server is running on http://%s\n", serverAddr)
	log.Fatal(http.ListenAndServe(serverAddr, nil))
}
