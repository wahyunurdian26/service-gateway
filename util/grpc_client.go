package util

import (
	"fmt"

	"google.golang.org/grpc"

	"microservice/util/client"
	pbaccount "microservice/cp-proto/account"
	pbtransaction "microservice/cp-proto/transaction"
)

func NewAccountServiceClient(serverAddr string) (pbaccount.AccountServiceClient, *grpc.ClientConn, error) {
	conn, err := client.NewGRPCConn(serverAddr)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to dial account service: %w", err)
	}
	return pbaccount.NewAccountServiceClient(conn), conn, nil
}

func NewTransactionServiceClient(serverAddr string) (pbtransaction.TransactionServiceClient, *grpc.ClientConn, error) {
	conn, err := client.NewGRPCConn(serverAddr)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to dial transaction service: %w", err)
	}
	return pbtransaction.NewTransactionServiceClient(conn), conn, nil
}
