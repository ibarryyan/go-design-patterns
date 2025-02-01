package _1_options

import "time"

type RpcServer struct {
	Name    string
	MaxConn int
	Address []string
	TimeOut time.Duration
}

type RpcServerOption func(server *RpcServer)

func WithName(name string) RpcServerOption {
	return func(server *RpcServer) {
		server.Name = name
	}
}

func WithMaxConn(max int) RpcServerOption {
	return func(server *RpcServer) {
		server.MaxConn = max
	}
}

func WithAddress(addr []string) RpcServerOption {
	return func(server *RpcServer) {
		server.Address = addr
	}
}

func WithTimeOut(timeout time.Duration) RpcServerOption {
	return func(server *RpcServer) {
		server.TimeOut = timeout
	}
}

func NewRpcServer(opts ...RpcServerOption) *RpcServer {
	server := &RpcServer{}

	for _, opt := range opts {
		opt(server)
	}

	return server
}
