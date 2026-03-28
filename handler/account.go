package handler

import (
	"microservice/gateway/handler/dto"
	"microservice/gateway/kit"
	uLog "microservice/util/logger"
)

func (h *GatewayHandler) serveGetAccountBalance(ctx kit.Context) (interface{}, error) {
	req, err := dto.ParseDefaultRequest(ctx)
	if err != nil {
		return nil, err
	}
	
	uLog.LogRequest(ctx, "serveGetAccountBalance", req)

	resp, err := h.accountClient.GetAccountBalance(ctx, req.ToAccountRequest())
	if err != nil {
		uLog.LogError(ctx, "h.accountClient.GetAccountBalance", err)
		return nil, err
	}

	return dto.MapAccountBalanceResponse(resp), nil
}
