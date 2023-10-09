package zenrows

import (
	"net/url"
	"strconv"
)

type ScrapeOptions func(values url.Values)

func WithJSRender(value bool) ScrapeOptions {
	return func(values url.Values) {
		values.Add("js_render", strconv.FormatBool(value))
	}
}

func WithCustomHeaders(value bool) ScrapeOptions {
	return func(values url.Values) {
		values.Add("custom_headers", strconv.FormatBool(value))
	}
}

func WithSessionID(value int) ScrapeOptions {
	return func(values url.Values) {
		values.Add("session_id", strconv.Itoa(value))
	}
}

func WithWait(value int) ScrapeOptions {
	return func(values url.Values) {
		values.Add("wait", strconv.Itoa(value))
	}
}

func WithPremiumProxy(value bool) ScrapeOptions {
	return func(values url.Values) {
		values.Add("premium_proxy", strconv.FormatBool(value))
	}
}

func WithProxyCountry(value string) ScrapeOptions {
	return func(values url.Values) {
		values.Add("proxy_country", value)
	}
}

func WithDevice(value string) ScrapeOptions {
	return func(values url.Values) {
		values.Add("device", value)
	}
}

func WithOriginalStatus(value bool) ScrapeOptions {
	return func(values url.Values) {
		values.Add("original_status", strconv.FormatBool(value))
	}
}

func WithWaitFor(value string) ScrapeOptions {
	return func(values url.Values) {
		values.Add("wait_for", value)
	}
}

func WithBlockResources(value string) ScrapeOptions {
	return func(values url.Values) {
		values.Add("block_resources", value)
	}
}

func WithJSONResponse(value bool) ScrapeOptions {
	return func(values url.Values) {
		values.Add("json_response", strconv.FormatBool(value))
	}
}

func WithWindowWidth(value int) ScrapeOptions {
	return func(values url.Values) {
		values.Add("window_width", strconv.Itoa(value))
	}
}

func WithWindowHeight(value int) ScrapeOptions {
	return func(values url.Values) {
		values.Add("window_height", strconv.Itoa(value))
	}
}

func WithCSSExtractor(value string) ScrapeOptions {
	return func(values url.Values) {
		values.Add("css_extractor", value)
	}
}

func WithAutoparse(value bool) ScrapeOptions {
	return func(values url.Values) {
		values.Add("autoparse", strconv.FormatBool(value))
	}
}

func WithResolveCaptcha(value bool) ScrapeOptions {
	return func(values url.Values) {
		values.Add("resolve_captcha", strconv.FormatBool(value))
	}
}

func ApplyParameters(u *url.URL, params ...ScrapeOptions) *url.URL {
	values := u.Query()
	for _, param := range params {
		param(values)
	}
	u.RawQuery = values.Encode()
	return u
}
