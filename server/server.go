package server

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net"

	sdsV3 "github.com/envoyproxy/go-control-plane/envoy/service/secret/v3"
	"github.com/mgcicd/cicd-core/logs"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func NewXdsServer(port int) {
	ctx := context.Background()

	cert, err := tls.LoadX509KeyPair("/opt/privateCerts/server.pem", "/opt/privateCerts/server.key")
	if err != nil {
		panic(err)
	}

	certPool := x509.NewCertPool()
	rootBuf, err := ioutil.ReadFile("/opt/privateCerts/ca.crt")
	if err != nil {
		panic(err)
	}
	if !certPool.AppendCertsFromPEM(rootBuf) {
		panic("Fail to append ca")
	}

	tlsConf := &tls.Config{
		ClientAuth:   tls.RequireAndVerifyClientCert,
		Certificates: []tls.Certificate{cert},
		ClientCAs:    certPool,
	}

	grpcServer := grpc.NewServer(grpc.Creds(credentials.NewTLS(tlsConf)))

	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", port))

	if err != nil {
		logs.DefaultConsoleLog.Error("NewXdsServer", fmt.Sprintf("err : %s", err))
	}

	serverHandle := &SecretDiscoveryService{}
	//注册sds
	sdsV3.RegisterSecretDiscoveryServiceServer(grpcServer, serverHandle)

	logs.DefaultConsoleLog.Info("NewXdsServer", fmt.Sprintf("xds GRPC Server Listening On %v", port))

	go func() {
		if err = grpcServer.Serve(lis); err != nil {
			logs.DefaultConsoleLog.Error("NewXdsServer", fmt.Sprintf("err : %s", err))
		}
	}()

	<-ctx.Done()

	grpcServer.GracefulStop()
}
