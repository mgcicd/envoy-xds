package server

import (
	"errors"
	"time"

	v3 "github.com/envoyproxy/go-control-plane/envoy/service/discovery/v3"
	"github.com/gogo/protobuf/proto"
	ptyes "github.com/golang/protobuf/ptypes"
	any "github.com/golang/protobuf/ptypes/any"
)

const (
	apiTypePrefix = "type.googleapis.com/"
	EndpointType  = apiTypePrefix + "envoy.config.endpoint.v3.ClusterLoadAssignment"
	ClusterType   = apiTypePrefix + "envoy.config.cluster.v3.Cluster"
	RouteType     = apiTypePrefix + "envoy.config.route.v3.RouteConfiguration"
	ListenerType  = apiTypePrefix + "envoy.config.listener.v3.Listener"
	SecretType    = apiTypePrefix + "envoy.extensions.transport_sockets.tls.v3.Secret"
	RuntimeType   = apiTypePrefix + "envoy.service.runtime.v3.Runtime"

	// AnyType is used only by ADS
	AnyType = ""
)

var (
	// RefreshDelay for the polling config source
	RefreshDelay = 5 * time.Second
	Timeout      = 5 * time.Second
	Interval     = 5 * time.Second
)

type Resource interface {
	proto.Message
}

type Response struct {
	// Request is the original request.
	Request v3.DiscoveryRequest

	// Version of the resources as tracked by the cache for the given type.
	// Proxy responds with this version as an acknowledgement.
	Version string

	// Resources to be included in the response.
	Resources []Resource
}

func CreateResponse(resp *Response, typeURL string) (*v3.DiscoveryResponse, error) {
	if resp == nil {
		return nil, errors.New("missing response")
	}

	resources := make([]*any.Any, len(resp.Resources))

	for i := 0; i < len(resp.Resources); i++ {

		data, err := ptyes.MarshalAny(resp.Resources[i])
		if err != nil {
			return nil, err
		}
		resources[i] = &any.Any{
			TypeUrl: typeURL,
			Value:   data.Value,
		}
	}
	out := &v3.DiscoveryResponse{
		VersionInfo: resp.Version,
		Resources:   resources,
		TypeUrl:     typeURL,
	}
	return out, nil
}
