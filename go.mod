module envoy-xds

go 1.14

replace github.com/mgcicd/cicd-core => ../cicd-core

require (
	github.com/envoyproxy/go-control-plane v0.9.5
	github.com/gogo/protobuf v1.1.1
	github.com/golang/protobuf v1.4.3
	github.com/mgcicd/cicd-core v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.31.0
)
