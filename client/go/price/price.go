package price

import (
	grpcClient "github.com/Rivalz-ai/proto/client/go/common"
	"github.com/Rivalz-ai/proto/pb/go/price"
)

var SVC price.PriceServiceClient

func NewPriceServiceClient(grpc_server_address string) error {
	conn, err := grpcClient.NewGRPCClientConn(grpc_server_address)
	if err != nil {
		return err
	}
	SVC = price.NewPriceServiceClient(conn)
	return nil
}
