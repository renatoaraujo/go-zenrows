package zenrows

import (
	"net/url"
	"strconv"
)

// ScrapeOptions Options to be passed to ZenRows api as query string
type ScrapeOptions func(values url.Values)

// WithJSRender Render the JavaScript on the page with a headless browser (5 credits/request)
func WithJSRender(value bool) ScrapeOptions {
	return func(values url.Values) {
		values.Add("js_render", strconv.FormatBool(value))
	}
}

// WithCustomHeaders Enable custom headers to be passed to the request.
func WithCustomHeaders(value bool) ScrapeOptions {
	return func(values url.Values) {
		values.Add("custom_headers", strconv.FormatBool(value))
	}
}

// WithPremiumProxy Use premium proxies to make the request harder to detect (10-25 credits/request)
func WithPremiumProxy(value bool) ScrapeOptions {
	return func(values url.Values) {
		values.Add("premium_proxy", strconv.FormatBool(value))
	}
}

// WithProxyCountry Geolocation of the IP used to make the request. Only for Premium Proxies.
func WithProxyCountry(value string) ScrapeOptions {
	return func(values url.Values) {
		values.Add("proxy_country", value)
	}
}

// WithSessionID Send a Session ID number to use the same IP for each API Request for up to 10 minutes.
func WithSessionID(value int) ScrapeOptions {
	return func(values url.Values) {
		values.Add("session_id", strconv.Itoa(value))
	}
}

// WithDevice Use either desktop or mobile user agents in the headers.
func WithDevice(value string) ScrapeOptions {
	return func(values url.Values) {
		values.Add("device", value)
	}
}

// WithOriginalStatus Returns the status code provided by the website.
func WithOriginalStatus(value bool) ScrapeOptions {
	return func(values url.Values) {
		values.Add("original_status", strconv.FormatBool(value))
	}
}

// WithWaitFor Wait for a given CSS Selector to load in the DOM before returning the content.
func WithWaitFor(value string) ScrapeOptions {
	return func(values url.Values) {
		values.Add("wait_for", value)
	}
}

// WithWait Wait a fixed amount of time before returning the content.
func WithWait(value int) ScrapeOptions {
	return func(values url.Values) {
		values.Add("wait", strconv.Itoa(value))
	}
}

// WithBlockResources Block specific resources from loading using this parameter.
func WithBlockResources(value string) ScrapeOptions {
	return func(values url.Values) {
		values.Add("block_resources", value)
	}
}

// WithJSONResponse Get content in JSON including XHR or Fetch requests.
func WithJSONResponse(value bool) ScrapeOptions {
	return func(values url.Values) {
		values.Add("json_response", strconv.FormatBool(value))
	}
}

// WithWindowWidth Set browser's window width.
func WithWindowWidth(value int) ScrapeOptions {
	return func(values url.Values) {
		values.Add("window_width", strconv.Itoa(value))
	}
}

// WithWindowHeight Set browser's window height.
func WithWindowHeight(value int) ScrapeOptions {
	return func(values url.Values) {
		values.Add("window_height", strconv.Itoa(value))
	}
}

// WithCSSExtractor Define CSS Selectors to extract data from the HTML.
func WithCSSExtractor(value string) ScrapeOptions {
	return func(values url.Values) {
		values.Add("css_extractor", value)
	}
}

// WithAutoparse Use our auto parser algorithm to automatically extract data.
func WithAutoparse(value bool) ScrapeOptions {
	return func(values url.Values) {
		values.Add("autoparse", strconv.FormatBool(value))
	}
}

// WithResolveCaptcha Use a CAPTCHA solver integration to automatically solve interactive CAPTCHAs in the page.
func WithResolveCaptcha(value bool) ScrapeOptions {
	return func(values url.Values) {
		values.Add("resolve_captcha", strconv.FormatBool(value))
	}
}

// ApplyParameters Applies the api parameters to the query string
func ApplyParameters(u *url.URL, params ...ScrapeOptions) *url.URL {
	values := u.Query()
	for _, param := range params {
		param(values)
	}
	u.RawQuery = values.Encode()
	return u
}
