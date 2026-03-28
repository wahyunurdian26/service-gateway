package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	kitlog "github.com/go-kit/log"
	"github.com/gorilla/mux"

	"microservice/gateway/config"
	"microservice/gateway/handler"
	"microservice/gateway/kit"
	"microservice/gateway/util"
	"microservice/util/middleware"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	var logger kitlog.Logger
	logger = kitlog.NewLogfmtLogger(os.Stderr)
	logger = kitlog.With(logger, "ts", kitlog.DefaultTimestampUTC, "caller", kitlog.DefaultCaller)

	cfg := config.LoadConfigs()

	httpPort := cfg.HttpPort
	accountBaseUrl := cfg.AccountBaseUrl
	transactionBaseUrl := cfg.TransactionBaseUrl

	accountClient, accountConn, err := util.NewAccountServiceClient(accountBaseUrl)
	if err != nil {
		logger.Log("err", err, "msg", "Failed to connect to account service")
		os.Exit(1)
	}
	defer accountConn.Close()

	transactionClient, transactionConn, err := util.NewTransactionServiceClient(transactionBaseUrl)
	if err != nil {
		logger.Log("err", err, "msg", "Failed to connect to transaction service")
		os.Exit(1)
	}
	defer transactionConn.Close()

	r := mux.NewRouter()
	kitRouter := kit.NewRouter(r)

	gatewayHandler := handler.NewGatewayHandler(accountClient, transactionClient, cfg.AuditBaseUrl)
	gatewayHandler.RegisterRoutes(kitRouter)

	errs := make(chan error)
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	go func() {
		logger.Log("transport", "HTTP", "addr", httpPort)
		errs <- http.ListenAndServe(":"+httpPort, middleware.DefaultHTTPHandler(r))
	}()

	logger.Log("exit", <-errs)
}
