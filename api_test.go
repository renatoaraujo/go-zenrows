package zenrows_test

import (
	"net/url"
	"testing"

	"github.com/renatoaraujo/go-zenrows"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestScrapeOptions(t *testing.T) {
	tests := []struct {
		name     string
		option   zenrows.ScrapeOptions
		expected url.Values
	}{
		{
			"WithJSRender",
			zenrows.WithJSRender(),
			url.Values{"js_render": []string{"true"}},
		},
		{
			"WithJSInstructions",
			zenrows.WithJSInstructions(`[{"click": ".selector"}, {"wait": 500}]`),
			url.Values{"js_render": []string{"true"}, "js_instructions": []string{`[{"click":".selector"},{"wait":500}]`}},
		},
		{
			"WithCustomHeaders",
			zenrows.WithCustomHeaders(true),
			url.Values{"custom_headers": []string{"true"}},
		},
		{
			"WithPremiumProxy",
			zenrows.WithPremiumProxy(),
			url.Values{"premium_proxy": []string{"true"}},
		},
		{
			"WithProxyCountry",
			zenrows.WithProxyCountry("US"),
			url.Values{"premium_proxy": []string{"true"}, "proxy_country": []string{"US"}},
		},
		{
			"WithBlockResources",
			zenrows.WithBlockResources("image"),
			url.Values{"js_render": []string{"true"}, "block_resources": []string{"image"}},
		},
		{
			"WithJSONResponse",
			zenrows.WithJSONResponse(true),
			url.Values{"js_render": []string{"true"}, "json_response": []string{"true"}},
		},
		{
			"WithWindowWidth",
			zenrows.WithWindowWidth(1920),
			url.Values{"js_render": []string{"true"}, "window_width": []string{"1920"}},
		},
		{
			"WithWindowHeight",
			zenrows.WithWindowHeight(1080),
			url.Values{"js_render": []string{"true"}, "window_height": []string{"1080"}},
		},
		{
			"WithCSSExtractor",
			zenrows.WithCSSExtractor(".content"),
			url.Values{"css_extractor": []string{".content"}},
		},
		{
			"WithAutoparse",
			zenrows.WithAutoparse(true),
			url.Values{"autoparse": []string{"true"}},
		},
		{
			"WithResolveCaptcha",
			zenrows.WithResolveCaptcha(true),
			url.Values{"js_render": []string{"true"}, "resolve_captcha": []string{"true"}},
		},
		{
			"WithDevice",
			zenrows.WithDevice("desktop"),
			url.Values{"device": []string{"desktop"}},
		},
		{
			"WithOriginalStatus",
			zenrows.WithOriginalStatus(true),
			url.Values{"original_status": []string{"true"}},
		},
		{
			"WithWaitFor",
			zenrows.WithWaitFor(".selector"),
			url.Values{"js_render": []string{"true"}, "wait_for": []string{".selector"}},
		},
		{
			"WithWait",
			zenrows.WithWait(500),
			url.Values{"js_render": []string{"true"}, "wait": []string{"500"}},
		},
		{
			"WithSessionID",
			zenrows.WithSessionID(12345),
			url.Values{"session_id": []string{"12345"}},
		},
		{
			"WithAIAntiBot",
			zenrows.WithAIAntiBot(),
			url.Values{"js_render": []string{"true"}, "antibot": []string{"true"}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			values := url.Values{}
			tt.option(values)
			assert.Equal(t, tt.expected, values)
		})
	}
}

func TestApplyParameters(t *testing.T) {
	tests := []struct {
		name     string
		options  []zenrows.ScrapeOptions
		expected url.Values
	}{
		{
			"Single Option - WithJSRender",
			[]zenrows.ScrapeOptions{zenrows.WithJSRender()},
			url.Values{"js_render": []string{"true"}},
		},
		{
			"Multiple Options - WithJSInstructions and WithCustomHeaders",
			[]zenrows.ScrapeOptions{
				zenrows.WithJSInstructions(`[{"click": ".selector"}, {"wait": 500}]`),
				zenrows.WithCustomHeaders(true),
			},
			url.Values{"js_render": []string{"true"}, "js_instructions": []string{`[{"click":".selector"},{"wait":500}]`}, "custom_headers": []string{"true"}},
		},
		{
			"Multiple Options - WithPremiumProxy and WithProxyCountry",
			[]zenrows.ScrapeOptions{
				zenrows.WithPremiumProxy(),
				zenrows.WithProxyCountry("US"),
			},
			url.Values{"premium_proxy": []string{"true"}, "proxy_country": []string{"US"}},
		},
		{
			"Multiple Options - WithBlockResources and WithJSONResponse",
			[]zenrows.ScrapeOptions{
				zenrows.WithBlockResources("image"),
				zenrows.WithJSONResponse(true),
			},
			url.Values{"js_render": []string{"true"}, "block_resources": []string{"image"}, "json_response": []string{"true"}},
		},
		{
			"Multiple Options - WithWindowWidth and WithWindowHeight",
			[]zenrows.ScrapeOptions{
				zenrows.WithWindowWidth(1920),
				zenrows.WithWindowHeight(1080),
			},
			url.Values{"js_render": []string{"true"}, "window_width": []string{"1920"}, "window_height": []string{"1080"}},
		},
		{
			"Multiple Options - WithCSSExtractor and WithAutoparse",
			[]zenrows.ScrapeOptions{
				zenrows.WithCSSExtractor(".content"),
				zenrows.WithAutoparse(true),
			},
			url.Values{"autoparse": []string{"true"}, "css_extractor": []string{".content"}},
		},
		{
			"Multiple Options - WithResolveCaptcha and WithJSInstructions",
			[]zenrows.ScrapeOptions{
				zenrows.WithResolveCaptcha(true),
				zenrows.WithJSInstructions(`[{"click": ".button"}]`),
			},
			url.Values{"js_render": []string{"true"}, "js_instructions": []string{`[{"click":".button"}]`}, "resolve_captcha": []string{"true"}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			baseURL, err := url.Parse("https://example.com")
			require.NoError(t, err)
			resultURL := zenrows.ApplyParameters(baseURL, tt.options...)
			assert.Equal(t, tt.expected, resultURL.Query())
		})
	}
}
