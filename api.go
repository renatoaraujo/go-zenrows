// Package zenrows provides utility functions to set scraping options
// for the ZenRows API. These functions help in configuring the request parameters
// for various scraping features and requirements.
package zenrows

import (
	"net/url"
	"strconv"
	"strings"
)

// ScrapeOptions defines functions that modify URL query values
// based on the chosen scraping options.
type ScrapeOptions func(values url.Values)

// WithJSRender enables JavaScript rendering for the scrape request.
// Consumes 5 credits per request.
func WithJSRender() ScrapeOptions {
	return func(values url.Values) {
		values.Set("js_render", "true")
	}
}

// WithJSInstructions provides JavaScript instructions for the scrape request.
// It automatically enables WithJSRender to ensure the correct execution of JavaScript instructions.
//
// value: A JSON string representing the JavaScript instructions.
func WithJSInstructions(value string) ScrapeOptions {
	condensed := strings.Join(strings.Fields(value), "")
	return func(values url.Values) {
		values.Set("js_render", "true")
		values.Set("js_instructions", condensed)
	}
}

// WithCustomHeaders allows custom headers to be added to the request.
//
// value: A boolean indicating if custom headers are to be included.
func WithCustomHeaders(value bool) ScrapeOptions {
	return func(values url.Values) {
		values.Set("custom_headers", strconv.FormatBool(value))
	}
}

// WithPremiumProxy enables the use of premium proxies for the request.
// This makes the request less detectable and consumes 10-25 credits per request.
func WithPremiumProxy() ScrapeOptions {
	return func(values url.Values) {
		values.Set("premium_proxy", "true")
	}
}

// WithProxyCountry specifies the geolocation of the IP for the request.
// Note: Only applicable for Premium Proxies.
//
// value: The desired country code for the proxy.
func WithProxyCountry(value string) ScrapeOptions {
	return func(values url.Values) {
		values.Set("premium_proxy", "true")
		values.Set("proxy_country", value)
	}
}

// WithBlockResources prevents specific resources from loading during the scrape request.
//
// value: The types of resources to block.
func WithBlockResources(value string) ScrapeOptions {
	return func(values url.Values) {
		values.Set("js_render", "true")
		values.Set("block_resources", value)
	}
}

// WithJSONResponse configures the request to return content in JSON format,
// including any XHR or Fetch requests made.
//
// value: A boolean to determine if the response should be in JSON format.
func WithJSONResponse(value bool) ScrapeOptions {
	return func(values url.Values) {
		values.Set("js_render", "true")
		values.Set("json_response", strconv.FormatBool(value))
	}
}

// WithWindowWidth defines the browser window width for the request.
//
// value: The desired window width in pixels.
func WithWindowWidth(value int) ScrapeOptions {
	return func(values url.Values) {
		values.Set("js_render", "true")
		values.Set("window_width", strconv.Itoa(value))
	}
}

// WithWindowHeight defines the browser window height for the request.
//
// value: The desired window height in pixels.
func WithWindowHeight(value int) ScrapeOptions {
	return func(values url.Values) {
		values.Set("js_render", "true")
		values.Set("window_height", strconv.Itoa(value))
	}
}

// WithCSSExtractor sets CSS Selectors to extract specific data from the HTML.
//
// value: The desired CSS selectors.
func WithCSSExtractor(value string) ScrapeOptions {
	return func(values url.Values) {
		values.Set("css_extractor", value)
	}
}

// WithAutoparse employs the auto-parser algorithm for the request,
// which extracts data from the page automatically.
//
// value: A boolean to determine if the auto parser should be used.
func WithAutoparse(value bool) ScrapeOptions {
	return func(values url.Values) {
		values.Set("autoparse", strconv.FormatBool(value))
	}
}

// WithResolveCaptcha integrates a CAPTCHA solver for the request,
// enabling automatic solving of CAPTCHAs on the page.
//
// value: A boolean to determine if the CAPTCHA solver should be used.
func WithResolveCaptcha(value bool) ScrapeOptions {
	return func(values url.Values) {
		values.Set("js_render", "true")
		values.Set("resolve_captcha", strconv.FormatBool(value))
	}
}

// WithDevice sets the user agent type (either desktop or mobile) for the request.
//
// value: A string specifying the device type ("desktop" or "mobile").
func WithDevice(value string) ScrapeOptions {
	return func(values url.Values) {
		values.Set("device", value)
	}
}

// WithOriginalStatus configures the request to return the status code as provided by the website.
//
// value: A boolean determining if the original status code should be returned.
func WithOriginalStatus(value bool) ScrapeOptions {
	return func(values url.Values) {
		values.Set("original_status", strconv.FormatBool(value))
	}
}

// WithWaitFor delays the request until a specific CSS Selector is loaded in the DOM.
//
// value: A string specifying the CSS Selector to wait for.
func WithWaitFor(value string) ScrapeOptions {
	return func(values url.Values) {
		values.Set("js_render", "true")
		values.Set("wait_for", value)
	}
}

// WithWait introduces a fixed delay before the content is returned.
//
// value: An integer specifying the wait time in milliseconds.
func WithWait(value int) ScrapeOptions {
	return func(values url.Values) {
		values.Set("js_render", "true")
		values.Set("wait", strconv.Itoa(value))
	}
}

// WithSessionID sets the Session ID number for the scrape request.
// This allows the use of the same IP for each API Request for up to 10 minutes.
//
// sessionID: An integer representing the Session ID.
func WithSessionID(sessionID int) ScrapeOptions {
	return func(values url.Values) {
		values.Set("session_id", strconv.Itoa(sessionID))
	}
}

// WithAIAntiBot sets the anti-bot
// Some websites protect their content with anti-bot solutions such as Cloudfare, Akamai, or Datadome. Enable Anti-bot to bypass them easily without any hassle.
func WithAIAntiBot() ScrapeOptions {
	return func(values url.Values) {
		values.Set("js_render", "true")
		values.Set("antibot", "true")
	}
}

// ApplyParameters applies the chosen scraping options to a URL.
// It modifies the URL's query string based on the provided scraping options.
//
// u: The target URL.
// params: The ScrapeOptions to be applied to the URL.
func ApplyParameters(u *url.URL, params ...ScrapeOptions) *url.URL {
	values := u.Query()
	for _, param := range params {
		param(values)
	}
	u.RawQuery = values.Encode()
	return u
}
