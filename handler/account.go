package handler

import (
	"github.com/wahyunurdian26/gateway/handler/dto"
	"github.com/wahyunurdian26/gateway/kit"
	uLog "github.com/wahyunurdian26/util/logger"
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
