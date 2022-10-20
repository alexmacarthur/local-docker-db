package main

import (
	"log"
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
}
