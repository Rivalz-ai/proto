package balance

import (
	grpcClient "github.com/Rivalz-ai/proto/client/go/common"
	"github.com/Rivalz-ai/proto/pb/go/balance"
)

var SVC balance.BalanceServiceClient

func NewBalanceServiceClient(grpc_server_address string) error {
	conn, err := grpcClient.NewGRPCClientConn(grpc_server_address)
	if err != nil {
		return err
	}
	SVC = balance.NewBalanceServiceClient(conn)
	return nil
}
