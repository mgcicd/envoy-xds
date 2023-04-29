package server

import (
	"context"
	"fmt"
	"net"

	sdsV3 "github.com/envoyproxy/go-control-plane/envoy/service/secret/v3"
	"github.com/mgcicd/cicd-core/logs"
	"google.golang.org/grpc"
)

func NewXdsServer(port int) {
	ctx := context.Background()

	server := grpc.NewServer()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", port))

	if err != nil {
		logs.DefaultConsoleLog.Error("NewXdsServer", fmt.Sprintf("err : %s", err))
	}

	serverHandle := &SecretDiscoveryService{}
	//注册sds
	sdsV3.RegisterSecretDiscoveryServiceServer(server, serverHandle)

	logs.DefaultConsoleLog.Info("NewXdsServer", fmt.Sprintf("xds GRPC Server Listening On %v", port))

	go func() {
		if err = server.Serve(lis); err != nil {
			logs.DefaultConsoleLog.Error("NewXdsServer", fmt.Sprintf("err : %s", err))
		}
	}()

	<-ctx.Done()

	server.GracefulStop()
}
