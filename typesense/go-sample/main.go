package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/typesense/typesense-go/typesense"
)

func main() {

	searchUrl, ok := os.LookupEnv("TYPESENSE_HOST")
	if !ok {
		searchUrl = "http://localhost:9108"
	}

	searchApiKey, ok := os.LookupEnv("TYPESENSE_API_KEY")
	if !ok {
		searchApiKey = ""
	}
	fmt.Print(searchApiKey)
	client := typesense.NewClient(
		typesense.WithServer(searchUrl),
		typesense.WithAPIKey(searchApiKey),
	)

	status, err := client.Health(time.Minute)

	if err != nil {
		log.Printf("Error getting health status: %v", err)
	}

	if status {
		log.Printf("connection successfully %v", status)
	} else {
		log.Print("connection failed")
	}

	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
