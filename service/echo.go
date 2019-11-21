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
func (d *Echo) Ping(req *string, res *string) (rpcErr error) {
	*res = "Pong"
	return
}
