package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"renatoaraujo/go-zenrows"
)

func main() {
	hc := &http.Client{
		Timeout: time.Duration(60) * time.Second,
	}
	client := zenrows.NewClient(hc).WithApiKey("YOUR_API_KEY")

	result, err := client.Scrape("https://httpbin.org")
	if err != nil {
		log.Fatalf("Failed to scrape the target: %v", err)
	}

	fmt.Println("Scraped Content:", result)
}
