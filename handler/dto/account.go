package dto

import (
	pbaccount "microservice/cp-proto/account"
)

type AccountBalanceResponse struct {
	AccountID string  `json:"account_id"`
	Balance   float64 `json:"balance"`
	Message   string  `json:"message"`
}

func MapAccountBalanceResponse(resp *pbaccount.AccountResponse) AccountBalanceResponse {
	return AccountBalanceResponse{
		AccountID: resp.AccountId,
		Balance:   resp.Balance,
		Message:   resp.Message,
	}
}
