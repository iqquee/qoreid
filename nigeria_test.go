package qoreid

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CreateService(t *testing.T) {
	httpClient := http.DefaultClient
	token := ""

	client := New(httpClient, token)
	assert.NotNil(t, client)
	assert.NotNil(t, client.Nigeria)

	payload := VerifyNinWithNinReq{
		IdNumber:  "000000000",
		Firstname: "jane",
		Lastname:  "doe",
	}

	response, err := client.Nigeria.VerifyNinWithNin(payload)
	if err != nil {
		assert.Error(t, err)
	}

	// TODO cannot be nil
	assert.Nil(t, response)
}

// TODO mock newRequest to accept custom request object and return custom request response
