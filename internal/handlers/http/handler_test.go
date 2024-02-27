package http

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vadimpk/card-validator-golang/config"
	"github.com/vadimpk/card-validator-golang/internal/services"
	"github.com/vadimpk/card-validator-golang/pkg/logging"
)

func TestNewHandler(t *testing.T) {
	h := NewHandler(Options{
		Services: services.Services{
			CardValidatorService: services.NewCardValidatorService(),
		},
		Logger: logging.New("debug"),
		Config: config.Get(),
	})

	ts := httptest.NewServer(h)
	defer ts.Close()

	testCases := []struct {
		name               string
		endpoint           string
		body               io.Reader
		expectedStatusCode int
	}{
		{
			name:     "positive: successful request for card validation",
			endpoint: "/api/validate",
			body: bytes.NewReader([]byte(`{
					"number": "4242424242424242",
					"expMonth": "02",
					"expYear": "2025"
				}`,
			)),
			expectedStatusCode: http.StatusOK,
		},
		{
			name:     "negative: invalid body",
			endpoint: "/api/validate",
			body: bytes.NewReader([]byte(`{
					"number": "4242424242424242",
					"expMonth": "02",
					"expYear": "2025"
				`,
			)),
			expectedStatusCode: http.StatusBadRequest,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res, err := http.Post(fmt.Sprint(ts.URL, tc.endpoint), "application/json", tc.body)

			assert.NoError(t, err)
			assert.Equal(t, res.StatusCode, tc.expectedStatusCode)
		})
	}
}
