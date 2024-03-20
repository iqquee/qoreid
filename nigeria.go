package qoreid

import (
	"fmt"
	"net/http"
)

// VerifyNinWithNin Verifies a customer's identity using their National Identity Number (NIN) through the National Identity Management Commission (NIMC).
func (n *Nigeria) VerifyNinWithNin(req VerifyNinWithNinReq) (*VerifyNinWithNinRes, error) {
	url := fmt.Sprintf("ng/identities/nin/%s", req.IdNumber)
	method := http.MethodPost

	if err := validate.Struct(&req); err != nil {
		return nil, err
	}

	var response VerifyNinWithNinRes
	if err := n.config.newRequest(method, url, req, response); err != nil {
		return nil, err
	}

	return &response, nil
}
