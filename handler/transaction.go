package handler

import (
	"github.com/wahyunurdian26/gateway/handler/dto"
	"github.com/wahyunurdian26/gateway/kit"
	uLog "github.com/wahyunurdian26/util/logger"
)

func (h *GatewayHandler) serveCreatePayment(ctx kit.Context) (interface{}, error) {
	req, err := dto.ParseDefaultRequest(ctx)
	if err != nil {
		return nil, err
	}

	uLog.LogRequest(ctx, "serveCreatePayment", req)

	resp, err := h.transactionClient.CreatePayment(ctx, req.ToPaymentRequest())
	if err != nil {
		uLog.LogError(ctx, "h.transactionClient.CreatePayment", err)
		return nil, err
	}

	return dto.MapPaymentResponse(resp), nil
}
