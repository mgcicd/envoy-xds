package server

import (
	"context"
	"envoy-xds/common"
	"fmt"
	"io"
	"io/ioutil"
	"log"

	cc3 "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	tls3 "github.com/envoyproxy/go-control-plane/envoy/extensions/transport_sockets/tls/v3"
	sd3 "github.com/envoyproxy/go-control-plane/envoy/service/discovery/v3"
	ss3 "github.com/envoyproxy/go-control-plane/envoy/service/secret/v3"
	"github.com/envoyproxy/go-control-plane/pkg/cache/types"
	"github.com/gogo/protobuf/proto"
	"github.com/golang/protobuf/ptypes/any"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type SecretDiscoveryService struct {
}

func (*SecretDiscoveryService) DeltaSecrets(server ss3.SecretDiscoveryService_DeltaSecretsServer) error {
	fmt.Println("DeltaSecrets()方法执行了")
	return status.Errorf(codes.Unimplemented, "method DeltaSecrets not implemented")
}

func (*SecretDiscoveryService) StreamSecrets(server ss3.SecretDiscoveryService_StreamSecretsServer) error {
	defer func() {
		if err := recover(); err != nil {
			common.Error("StreamSecrets", fmt.Sprintf("StreamSecrets err: %s", err), "")
		}
	}()

	fmt.Println("StreamSecrets()方法执行了")
	for {
		_, err := server.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Printf("failed to recv: %v", err)
			return err
		}
		//log.Printf("server recv: %s", in)
		resp := getResp()
		err = server.Send(&resp)
		if err != nil {
			log.Printf("failed to send: %v", err)
			return err
		}
	}
}

func (*SecretDiscoveryService) FetchSecrets(ctx context.Context, request *sd3.DiscoveryRequest) (*sd3.DiscoveryResponse, error) {
	fmt.Println("FetchSecrets()方法执行了")
	return nil, status.Errorf(codes.Unimplemented, "method FetchSecrets not implemented")
}

func getResp() sd3.DiscoveryResponse {
	buff := proto.NewBuffer(nil)
	buff.SetDeterministic(true)

	fileByte, err := ioutil.ReadFile("/opt/publicCerts/34580.com.crt")
	if err != nil {
		log.Printf("获取私钥失败: %v", err)
	}

	file2Byte, err := ioutil.ReadFile("/opt/publicCerts/34580.com.key")
	if err != nil {
		log.Printf("获取证书失败: %v", err)
	}

	resources := []types.Resource{
		&tls3.Secret{
			Name: "suiyi_secret",
			Type: &tls3.Secret_TlsCertificate{
				TlsCertificate: &tls3.TlsCertificate{
					CertificateChain: &cc3.DataSource{
						Specifier: &cc3.DataSource_InlineString{
							InlineString: string(fileByte),
						},
					},
					PrivateKey: &cc3.DataSource{
						Specifier: &cc3.DataSource_InlineString{
							InlineString: string(file2Byte),
						},
					},
				},
			},
		},
	}

	err = buff.Marshal(resources[0])

	if err != nil {
		log.Fatal(err)
	}

	return sd3.DiscoveryResponse{
		VersionInfo: "2.0",
		Resources: []*any.Any{
			{
				TypeUrl: SecretType,
				Value:   buff.Bytes(),
			},
		},
		TypeUrl: SecretType,
	}
}
