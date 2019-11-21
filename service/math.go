package service

import (
	"rpc_server/core/rpc"
)

func init() {
	rpc.Register("Math", "Do so math", &Math{})
}

//Math Math
type Math struct{}

//Double Double int64
func (m *Math) Double(req *int64, res *int64) (rpcErr error) {
	*res = *req * 2
	return
}