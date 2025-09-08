package biconomy

import (
	grpcClient "github.com/Rivalz-ai/proto/client/go/common"
	"github.com/Rivalz-ai/proto/pb/go/biconomy"
)

var SVC biconomy.BiconomyServiceClient

func NewBiconomyServiceClient(grpc_server_address string) error {
	conn, err := grpcClient.NewGRPCClientConn(grpc_server_address)
	if err != nil {
		return err
	}
	SVC = biconomy.NewBiconomyServiceClient(conn)
	return nil
}
