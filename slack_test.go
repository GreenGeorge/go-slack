package slack

import (
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	accessTokenStub = "123456789"
	httpClientStub  = &http.Client{Timeout: 5 * time.Second}
)

type newClientTestCase struct {
	accessToken string
	httpClient  *http.Client
	expected    newClientExpected
}

type newClientExpected struct {
	accessToken string
	httpClient  *http.Client
}

func TestNew(t *testing.T) {
	tests := map[string]newClientTestCase{
		"Constructs a valid client": {
			accessToken: accessTokenStub,
			httpClient:  httpClientStub,
			expected: newClientExpected{
				accessToken: accessTokenStub,
				httpClient:  httpClientStub,
			},
		},
		"Constructs a valid client with default http client": {
			accessToken: accessTokenStub,
			httpClient:  nil,
			expected: newClientExpected{
				accessToken: accessTokenStub,
				httpClient:  http.DefaultClient,
			},
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			expected := test.expected
			actual := New(test.accessToken, test.httpClient)
			assert.Equal(t, expected.accessToken, actual.AccessToken())
			assert.Equal(t, expected.httpClient, actual.HttpClient())
		})
	}
}
