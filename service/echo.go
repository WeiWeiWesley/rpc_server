package service

import (
	"rpc_server/core/rpc"
)

func init() {
	rpc.Register("Echo", "A Echo Service", &Echo{})
}

// Echo echo service
type Echo struct{}

//Ping Ping
func (d *Echo) Ping(args *string, result *string) (rpcErr error) {
	*result = "Pong"
	return
}
