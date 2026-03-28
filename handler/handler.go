package handler

import (
	"github.com/wahyunurdian26/gateway/kit"
	pbaccount "github.com/wahyunurdian26/client/account"
	pbtransaction "github.com/wahyunurdian26/client/transaction"
)

type GatewayHandler struct {
	accountClient     pbaccount.AccountServiceClient
	transactionClient pbtransaction.TransactionServiceClient
	auditBaseUrl      string
}

func NewGatewayHandler(accountClient pbaccount.AccountServiceClient, transactionClient pbtransaction.TransactionServiceClient, auditBaseUrl string) *GatewayHandler {
	return &GatewayHandler{
		accountClient:     accountClient,
		transactionClient: transactionClient,
		auditBaseUrl:      auditBaseUrl,
	}
}

func (h *GatewayHandler) RegisterRoutes(r *kit.Router) {
	r.Get("/v1/accounts/{id}/balance", h.serveGetAccountBalance)
	r.Post("/v1/payments", h.serveCreatePayment)
	r.Get("/v1/audits", h.serveGetListAuditLog)
}
