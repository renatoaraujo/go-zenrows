package zenrows_test

import (
	"net/http"
	"testing"

	"renatoaraujo/go-zenrows"

	"github.com/stretchr/testify/assert"
)

func TestNewClient(t *testing.T) {
	client := zenrows.NewClient(&http.Client{})
	assert.NotNil(t, client, "Client should not be nil")
	assert.IsType(t, &zenrows.Client{}, client)
}
