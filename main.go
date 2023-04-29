package main

import (
	"envoy-xds/common"
	"envoy-xds/server"
	"flag"
	"fmt"
	"strconv"
	"time"
)

var portFlag = flag.String("xds_port", "50051", "EnvoyXds-Port")

func main() {

	port, err := strconv.Atoi(*portFlag)

	if err != nil {
		common.Error("", fmt.Sprintf("envoy-xds start fail err : %s", err), "")
		panic(err)
	}
	fmt.Println("Application starting in 50051," + time.Now().Format(time.RFC3339))
	server.NewXdsServer(port)
}
