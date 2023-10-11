package zenrows

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func validateFullURL(targetURL string) error {
	u, err := url.Parse(targetURL)
	if err != nil {
		return fmt.Errorf("failed to parse target url, please provide a valid url; %w", err)
	}

	if u.Scheme == "" || u.Host == "" {
		return fmt.Errorf("provided string is not a full URL, it should include both a scheme and a host")
	}

	return nil
}

// Scrape fetches content from the specified targetURL using the ZenRows API.
//
// The function constructs the API URL based on the provided targetURL and optional ScrapeOptions.
// It then sends a GET request to the ZenRows API and returns the scraped content as a string.
//
// The function validates the provided targetURL to ensure it's a full URL with both a scheme and a host.
// It also checks if the 'js_instructions' parameter is set without enabling 'js_render', and returns an error if so.
//
// # For now only supports GET method
//
// Parameters:
// - ctx: Context
// - targetURL: The URL of the website you want to scrape.
// - params: Optional parameters to customize the scraping process. Refer to ScrapeOptions for available options.
//
// Returns:
// - A string containing the scraped content.
// - An error if there's any issue during the scraping process, such as invalid URLs, failed requests, or reading issues.
//
// Example usage:
//
//	content, err := client.Scrape(context.Background(), "https://example.com", zenrows.WithJSRender(true))
//	if err != nil {
//	    log.Fatalf("Failed to scrape the target: %v", err)
//	}
//	fmt.Println("Scraped Content:", content)
//
// For more details and examples, refer to the https://pkg.go.dev/github.com/renatoaraujo/go-zenrows and the example provided in the repository https://github.com/renatoaraujo/go-zenrows/blob/main/examples/example.go.
func (c *Client) Scrape(ctx context.Context, targetURL string, params ...ScrapeOptions) (string, error) {
	select {
	case <-ctx.Done():
		return "", ctx.Err()
	default:
	}

	if err := validateFullURL(targetURL); err != nil {
		return "", fmt.Errorf("failed to parse target url: %w", err)
	}

	apiURL, err := c.constructAPIURL(targetURL, params...)
	if err != nil {
		return "", err
	}

	return c.fetchContent(ctx, apiURL)
}

func (c *Client) constructAPIURL(targetURL string, params ...ScrapeOptions) (*url.URL, error) {
	baseURL, err := url.Parse(c.config.BaseURL)
	if err != nil {
		return nil, fmt.Errorf("failed to parse base zenrows url: %w", err)
	}

	addTokenParams := func(values url.Values) {
		values.Add("apikey", c.config.key)
		values.Add("url", targetURL)
	}

	allParams := append([]ScrapeOptions{addTokenParams}, params...)
	return ApplyParameters(baseURL, allParams...), nil
}

func (c *Client) fetchContent(ctx context.Context, apiURL *url.URL) (string, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, apiURL.String(), nil)
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}

	return string(body), nil
}
