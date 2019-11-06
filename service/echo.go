package service

// Echo echo service
type Echo struct{}

//Ping Ping
func (d *Echo) Ping(args *string, result *string) (rpcErr error) {
	*result = "Pong"
	return
}
