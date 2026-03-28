package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/wahyunurdian26/gateway/kit"
	uLog "github.com/wahyunurdian26/util/logger"
)

type auditResponse struct {
	Meta interface{} `json:"meta"`
	Data interface{} `json:"data"`
}

func (h *GatewayHandler) serveGetListAuditLog(ctx kit.Context) (interface{}, error) {
	uLog.LogRequest(ctx, "serveGetListAuditLog", ctx.Request().URL.Query())
	
	url := fmt.Sprintf("%s/v1/audits?%s", h.auditBaseUrl, ctx.Request().URL.RawQuery)
	resp, err := http.Get(url)
	if err != nil {
		uLog.LogError(ctx, "http.Get(audit-service)", err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		err := fmt.Errorf("audit-service returned status code: %d", resp.StatusCode)
		uLog.LogError(ctx, "audit-service.Response", err)
		return nil, err
	}

	var result struct {
		Result auditResponse `json:"result"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		uLog.LogError(ctx, "json.NewDecoder(audit-service)", err)
		return nil, err
	}

	return kit.RawResponse{Data: result.Result}, nil
}
