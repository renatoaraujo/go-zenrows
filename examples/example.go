package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/renatoaraujo/go-zenrows"
)

func main() {
	hc := &http.Client{
		Timeout: time.Duration(60) * time.Second,
	}
	client := zenrows.NewClient(hc).WithApiKey("ZENROWS_API_KEY")

	jsInstructions := `[
  {"click": ".selector"},
  {"wait": 500},
  {"fill": [".input", "value"]},
  {"wait_for": ".slow_selector"}
]
`
	// add options, e.g.: add JS instructions; or just call Scrape("http://...")
	result, err := client.Scrape(context.Background(), "https://httpbin.org", zenrows.WithJSInstructions(jsInstructions))
	if err != nil {
		log.Fatalf("Failed to scrape the target: %v", err)
	}

	fmt.Println("Scraped Content:", result)
}
