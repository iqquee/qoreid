package qoreid

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_New(t *testing.T) {
	httpClient := http.DefaultClient
	token := ""

	client := New(httpClient, token)

	assert.NotNil(t, client)

	assert.NotNil(t, client.Nigeria)

	assert.Equal(t, httpClient, client.Nigeria.config.Http)
	assert.Equal(t, token, client.Nigeria.config.Token)
}

func Test_error(t *testing.T) {
	successResponse := make(map[string]interface{})
	successResponse["id"] = "12345"
	successResponse["message"] = "success"

	errorResponse := make(map[string]interface{})
	errorResponse["id"] = "12345"
	errorResponse["message"] = "an error occured"

	testCases := []struct {
		name            string
		errorResponse   map[string]interface{}
		successResponse map[string]interface{}
		StatusCode      int
	}{
		{
			name:          "error response status 4XX",
			errorResponse: errorResponse,
			StatusCode:    400,
		},
		{
			name:          "error response status 5XX",
			errorResponse: errorResponse,
			StatusCode:    500,
		},
		{
			name:            "success response status 2**",
			successResponse: successResponse,
			StatusCode:      200,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			response := make(map[string]interface{})

			if strings.Contains(testCase.name, "error response") {
				if testCase.StatusCode == 400 {
					mockResponseBytes, _ := json.Marshal(testCase.errorResponse)
					mockHTTPResponse := httptest.NewRecorder()
					mockHTTPResponse.WriteHeader(http.StatusBadRequest)
					mockHTTPResponse.Body.Write(mockResponseBytes)

					config := Config{}
					if err := config.error(mockHTTPResponse.Result(), response); err != nil {
						assert.Error(t, err)
					}

					assert.NotNil(t, response)
				} else if testCase.StatusCode == 500 {
					mockResponseBytes, _ := json.Marshal(testCase.errorResponse)
					mockHTTPResponse := httptest.NewRecorder()
					mockHTTPResponse.WriteHeader(http.StatusInternalServerError)
					mockHTTPResponse.Body.Write(mockResponseBytes)

					config := Config{}
					if err := config.error(mockHTTPResponse.Result(), response); err != nil {
						assert.Error(t, err)
					}

					assert.NotNil(t, response)
				} else {
					mockResponseBytes, _ := json.Marshal(testCase.successResponse)
					mockHTTPResponse := httptest.NewRecorder()
					mockHTTPResponse.WriteHeader(http.StatusOK)
					mockHTTPResponse.Body.Write(mockResponseBytes)

					config := Config{}
					if err := config.error(mockHTTPResponse.Result(), response); err != nil {
						assert.Error(t, err)
					}

					assert.NotNil(t, response)
				}

			}
		})
	}
}

func Test_newRequest(t *testing.T) {
	httpClient := http.DefaultClient
	token := ""
	baseUrl := ""
	config := Config{
		Http:    httpClient,
		Token:   token,
		BaseUrl: baseUrl,
	}

	reqBody := make(map[string]interface{})
	reqBody["data"] = "yay!!!"
	response := make(map[string]interface{})
	if err := config.newRequest(http.MethodPost, "", reqBody, response); err != nil {
		assert.Error(t, err)
	}

	assert.NotNil(t, response)
}
