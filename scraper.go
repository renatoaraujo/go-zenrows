package zenrows

import (
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

func (c *Client) Scrape(targetURL string, params ...ScrapeOptions) (string, error) {
	err := validateFullURL(targetURL)
	if err != nil {
		return "", fmt.Errorf("failed to parse target url, please provide a valid url; %w", err)
	}

	api, err := url.Parse(c.config.BaseURL)
	if err != nil {
		return "", fmt.Errorf("failed to parse base zenrows url; %w", err)
	}

	token := func(values url.Values) {
		values.Add("apikey", c.config.key)
		values.Add("url", targetURL)
	}

	allParams := append([]ScrapeOptions{token}, params...)
	api = ApplyParameters(api, allParams...)

	req, err := http.NewRequest(http.MethodGet, api.String(), nil)
	if err != nil {
		return "", fmt.Errorf("failed to create request; %w", err)
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to make request; %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed read response body; %w", err)
	}

	return string(body), nil
}
