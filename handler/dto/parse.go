package dto

import (
	pbaccount "github.com/wahyunurdian26/client/account"
	"github.com/wahyunurdian26/gateway/kit"
)

type ParseRequest struct {
	ID           string  `json:"id"`
	AccountID    string  `json:"account_id"`
	Amount       float64 `json:"amount"`
	MerchantName string  `json:"merchant_name"`
	Description  string  `json:"description"`
}

func ParseDefaultRequest(ctx kit.Context) (ParseRequest, error) {
	var req ParseRequest
	
	// Handle path variables
	if id := ctx.GetPathVariable("id"); id != "" {
		req.ID = id
	}

	// Handle body for POST requests
	if ctx.Request().Method == "POST" {
		if err := ctx.BindJSON(&req); err != nil {
			return req, err
		}
	}

	return req, nil
}

func (req ParseRequest) ToAccountRequest() *pbaccount.AccountRequest {
	return &pbaccount.AccountRequest{
		AccountId: req.ID,
	}
}
