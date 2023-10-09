package zenrows_test

import (
	"bytes"
	"errors"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"testing"

	"renatoaraujo/go-zenrows"
	mocks "renatoaraujo/go-zenrows/mocks"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestScrape(t *testing.T) {
	tests := []struct {
		name            string
		url             string
		httpClientSetup func(client *mocks.HttpClient)
		result          string
		expectError     bool
	}{
		{
			name: "Success scraping data from website",
			url:  "http://example.com",
			httpClientSetup: func(s *mocks.HttpClient) {
				s.On("Do", mock.Anything).
					Once().
					Return(&http.Response{
						StatusCode: 200,
						Body:       io.NopCloser(bytes.NewReader([]byte("some content"))),
					}, nil)
			},
			result:      "some content",
			expectError: false,
		},
		{
			name: "Failed to make the request",
			url:  "http://example.com",
			httpClientSetup: func(s *mocks.HttpClient) {
				s.On("Do", mock.Anything).
					Once().
					Return(nil, errors.New("failed to make the request"))
			},
			expectError: true,
		},
		{
			name:        "Failed to scrape with valid url",
			url:         "invalid",
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			httpClientMock := mocks.NewHttpClient(t)

			if tt.httpClientSetup != nil {
				tt.httpClientSetup(httpClientMock)
			}

			client := zenrows.NewClient(httpClientMock).
				WithApiKey("key")
			content, err := client.Scrape(tt.url)

			if tt.expectError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}

			assert.Equal(t, tt.result, content)
		})
	}
}
