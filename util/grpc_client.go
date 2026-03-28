package util

import (
	"fmt"

	"google.golang.org/grpc"

	"github.com/wahyunurdian26/client"
	pbaccount "github.com/wahyunurdian26/client/account"
	pbtransaction "github.com/wahyunurdian26/client/transaction"
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
