go-zenrows
===

[![codecov](https://codecov.io/gh/renatoaraujo/go-zenrows/graph/badge.svg?token=ORVP7TXY4A)](https://codecov.io/gh/renatoaraujo/go-zenrows)

`go-zenrows` is a Go client for the ZenRows API, allowing users to easily scrape web content.

## Features

- **Scrape Web Content**: Easily scrape content from any website using the ZenRows API.
- **Flexible Configuration**: Comes with a default configuration but allows for customization.
- **Various Scrape Options**: Customize your scraping with options like JS rendering, custom headers, session ID, and more.
- **Examples Included**: A basic example is provided to help you get started quickly.

## Usage

Here's a basic example to get you started:

```go
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

	// add JS rendering
	r2, err := client.Scrape("https://httpbin.org", zenrows.WithJSRender(true))
	if err != nil {
		log.Fatalf("Failed to scrape the target: %v", err)
	}

	fmt.Println("Scraped Content:", r2)
}
```

[View the full example here](examples/example.go).

## Documentation

For a detailed list of all available functions and scrape options, refer to the official documentation:
- [ZenRows docs website](https://www.zenrows.com/docs)

## Credits

* [Renato Araujo](https://www.linkedin.com/in/renatoraraujo/)

## License

The MIT License (MIT) - see [`LICENSE`](LICENSE) for more details