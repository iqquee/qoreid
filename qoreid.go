package qoreid

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog"
)

var (
	// validate runs a validation on the incoming json payload
	validate = validator.New(validator.WithRequiredStructEnabled())
	// Logger error logger
	Logger = zerolog.New(os.Stderr).With().Timestamp().Logger()
	// errErrorOccured for an error of error occured
	errErrorOccured = errors.New("an error occured")
)

type (
	// Client is an object for the objects implemeting different methods
	Client struct {
		Nigeria Nigeria
	}

	// Nigeria is an object for methods under the nigeria
	Nigeria struct {
		config Config
	}

	// Config is an object for the configs
	Config struct {
		Http    *http.Client
		BaseUrl string
		Token   string
	}
)

// New is the qoreid config initializer
func New(h *http.Client, token string) *Client {
	config := Config{
		Http:    h,
		BaseUrl: "",
		Token:   token,
	}

	nigeria := Nigeria{
		config: config,
	}

	return &Client{
		Nigeria: nigeria,
	}
}

// newRequest makes a http request to the render server and decodes the server response into the reqBody parameter passed into the newRequest method
func (c *Config) newRequest(method, reqURL string, reqBody, resp interface{}) error {
	newURL := c.BaseUrl + reqURL
	var body io.Reader

	if reqBody != nil {
		bb, err := json.Marshal(reqBody)
		if err != nil {
			return err
		}
		body = bytes.NewReader(bb)
	}

	req, err := http.NewRequest(method, newURL, body)
	if reqBody != nil {
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", c.Token))
	}

	if err != nil {
		return err
	}

	res, err := c.Http.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if err := c.error(res, resp); err != nil {
		return err
	}

	return nil
}

// error checks for errors in respect to the status code and return response payload accordingly and log error if any...
func (c *Config) error(httpRes *http.Response, response interface{}) error {
	statusCode := fmt.Sprint(httpRes.StatusCode)
	if strings.HasPrefix(statusCode, "4") || strings.HasPrefix(statusCode, "5") {
		var errorRes map[string]interface{}
		if err := json.NewDecoder(httpRes.Body).Decode(&errorRes); err != nil {
			return err
		}

		Logger.Err(errErrorOccured).Msgf("Code: %d ::: Error: [ %v ]", httpRes.StatusCode, errorRes)
	}

	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return err
	}

	return nil
}
