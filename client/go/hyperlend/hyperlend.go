package hyperlend

import (
	grpcClient "github.com/Rivalz-ai/proto/client/go/common"
	"github.com/Rivalz-ai/proto/pb/go/hyperlend"
)

var SVC hyperlend.HyperLendServiceClient

func NewHyperlendServiceClient(grpc_server_address string) error {
	conn, err := grpcClient.NewGRPCClientConn(grpc_server_address)
	if err != nil {
		return err
	}
	SVC = hyperlend.NewHyperLendServiceClient(conn)
	return nil
}
