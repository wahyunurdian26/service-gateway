package dto

import (
	pbtransaction "microservice/cp-proto/transaction"
)

type PaymentResponse struct {
	TransactionID string `json:"transaction_id"`
	Status        string `json:"status"`
	Message       string `json:"message"`
}

func (req ParseRequest) ToPaymentRequest() *pbtransaction.PaymentRequest {
	return &pbtransaction.PaymentRequest{
		AccountId:    req.AccountID,
		Amount:       req.Amount,
		MerchantName: req.MerchantName,
		Description:  req.Description,
	}
}

func MapPaymentResponse(resp *pbtransaction.PaymentResponse) PaymentResponse {
	return PaymentResponse{
		TransactionID: resp.TransactionId,
		Status:        resp.Status,
		Message:       resp.Message,
	}
}
