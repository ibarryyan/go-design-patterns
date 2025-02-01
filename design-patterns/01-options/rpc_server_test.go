package _1_options

import (
	"fmt"
	"testing"
	"time"
)

func TestCreateRpcServerByOptions(t *testing.T) {
	rpcServer := NewRpcServer(
		WithAddress([]string{"127.0.0.1"}),
		WithName("rpcServer"),
		WithMaxConn(1),
		WithTimeOut(time.Second),
	)

	fmt.Println(*rpcServer)
}
