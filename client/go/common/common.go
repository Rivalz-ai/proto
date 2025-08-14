package common

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/Rivalz-ai/proto/client/go/common/utils"
	"github.com/joho/godotenv"
	"go.elastic.co/apm/module/apmgrpc"
	"golang.org/x/net/idna"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
)

type TimeoutCallOption struct {
	grpc.EmptyCallOption

	forcedTimeout time.Duration
}

func NewGRPCClientConn(grpc_server_address string, args ...interface{}) (*grpc.ClientConn, error) {
	LoadENV()
	var cfg map[string]interface{}
	var tls credentials.TransportCredentials
	time_out_cfg := GRPC_MAX_TIMEOUT      //mins
	max_msg_size_cfg := GRPC_MAX_MSG_SIZE //mb
	if len(args) > 0 {
		cfg, _ = utils.ItoDictionary(args[0])

	}
	fw_ctx := context.Background()
	if cfg != nil {
		if val, ok := cfg["max_timeout"]; ok {
			time_out_cfg = utils.ItoInt(val)
			if time_out_cfg <= 0 {
				time_out_cfg = GRPC_MAX_TIMEOUT
			}
		}
		if val, ok := cfg["max_msg_size"]; ok {
			max_msg_size_cfg = utils.ItoInt(val)
			if max_msg_size_cfg <= 0 {
				max_msg_size_cfg = GRPC_MAX_MSG_SIZE
			}
		}
		if val, ok := cfg["tls"]; ok {
			tls_c, ok := val.(credentials.TransportCredentials)
			if ok {
				tls = tls_c
			}
		}
		if val, ok := cfg["ctx"]; ok {
			temp_ctx, ok := val.(context.Context)
			if ok {
				fw_ctx = temp_ctx
			}
		}
	}
	if grpc_server_address == "" {
		return nil, errors.New("RPC Server Address is empty")
	}
	if IsServiceName(grpc_server_address) {
		grpc_server_address = GetEndpoint(grpc_server_address)
	}
	if grpc_server_address == "" {
		return nil, errors.New("Service name invalid")
	}
	arr := utils.Explode(grpc_server_address, ":")
	if len(arr) == 1 {
		grpc_server_address = fmt.Sprintf("%s:%s", grpc_server_address, "30000")
	}
	//
	maxMsgSize := 1024 * 1024 * max_msg_size_cfg //mb
	//
	//time.Duration(time_out_cfg)*60
	transportOption := NewDialOption()
	ctx, _ := context.WithTimeout(fw_ctx, time.Duration(GRPC_MAX_TIMEOUT)*60*time.Second)
	conn, err := grpc.DialContext(ctx, grpc_server_address,
		grpc.WithTransportCredentials(tls),
		grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(maxMsgSize),
			grpc.MaxCallSendMsgSize(maxMsgSize)),
		//grpc.WithDefaultServiceConfig(serviceConfig),
		grpc.WithChainUnaryInterceptor(
			TimeoutInterceptor(time.Duration(time_out_cfg)*60*time.Second),
			apmgrpc.NewUnaryClientInterceptor(),
			UnaryClientInfoInterceptor,
		),
		transportOption)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
func LoadENV() {
	err := godotenv.Load(os.ExpandEnv("/config/.env"))
	if err != nil {
		err := godotenv.Load(os.ExpandEnv(".env"))
		if err != nil {
			panic(err)
		}
	}
}

func IsServiceName(endpoint string) bool {
	if len(endpoint) > 64 {
		panic("Service name is too long")
	}
	_, err := idna.Lookup.ToASCII(endpoint)
	if err != nil {
		return false
	}
	return true
}
func GetEndpoint(service_name string) string {
	env := os.Getenv("CLUSTER")
	if env == "" {
		return ""
	}
	if IsServiceName(service_name) {
		//service_name=strings.ReplaceAll(service_name,"services.","")
		arr := utils.Explode(service_name, ".")
		arr = utils.ReverseStringArray(arr)
		if len(arr) == 2 {
			return strings.Join(arr, ".") + ".svc." + env
		} else if len(arr) > 2 {
			prefix := arr[0 : len(arr)-2]
			suffix := arr[len(arr)-2 : len(arr)]
			//fmt.Println("prefix",prefix)
			//fmt.Println("sufix",sufix)
			return strings.Join(prefix, "--") + "--" + strings.Join(suffix, ".") + ".svc." + env
		}

	} else {
		service_name = strings.ReplaceAll(service_name, "@", "")
	}
	return service_name
}
func NewDialOption() grpc.DialOption {
	return grpc.WithInsecure()
}

func TimeoutInterceptor(t time.Duration) grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		timeout := t
		if v, ok := getForcedTimeout(opts); ok {
			timeout = v
		}

		if timeout <= 0 {
			return invoker(ctx, method, req, reply, cc, opts...)
		}

		ctx, cancel := context.WithTimeout(ctx, timeout)
		defer cancel()

		return invoker(ctx, method, req, reply, cc, opts...)
	}
}
func getForcedTimeout(callOptions []grpc.CallOption) (time.Duration, bool) {
	for _, opt := range callOptions {
		if co, ok := opt.(TimeoutCallOption); ok {
			return co.forcedTimeout, true
		}
	}

	return 0, false
}
func UnaryClientInfoInterceptor(ctx context.Context, method string, req interface{}, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	//add service name to all grpc requests
	fw_client_name := ""
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		if md["client"] != nil {
			if len(md["client"]) > 0 {
				fw_client_name = md["client"][0]
				if fw_client_name != "" {
					fw_client_name += "|"
				}
			}
		}
	}
	fw_client_name += GetClientServiceName()
	ctx = metadata.AppendToOutgoingContext(ctx, "client", fw_client_name)
	return invoker(ctx, method, req, reply, cc, opts...)
}
func GetClientServiceName() string {
	hostname, _ := os.Hostname()
	if hostname == "" {
		fmt.Println("hostname is empty")
		return ""
	}
	arr := utils.Explode(hostname, "-")
	if len(arr) == 0 {
		return hostname
	}
	if utils.StringToInt(arr[len(arr)-1]) < 0 { //deploymentset
		if len(arr) > 2 {
			return strings.Join(arr[0:len(arr)-2], "-")
		}
		return hostname
	} else { //statefulset
		return strings.Join(arr[0:len(arr)-1], "-")
	}
	return hostname
}
