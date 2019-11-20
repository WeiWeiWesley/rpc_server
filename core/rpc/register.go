package rpc

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

// Active Active Service
func Active(name string) {
	if name == "All" {
		// Active all registed services
		for serviceName, service := range ServiceRegisted {
			err := rpc.RegisterName(serviceName, service.Instance)
			if err != nil {
				log.Println("[rpc_error]:", err.Error())
				continue
			}
			log.Printf("Start service: %s; success\n", serviceName)
		}
	} else {
		// Active service by name
		if service, ok := ServiceRegisted[name]; ok {
			// Active from registed service
			err := rpc.RegisterName(name, service.Instance)
			if err != nil {
				log.Println("[rpc_error]:", err.Error())
			}
		} else {
			log.Fatalf("[Fatal error] Service '%s' not exist", name)
		}
		log.Printf("Start service: %s; success\n", name)
	}

	// Star listen
	l, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Println("[rpc_error]:", err.Error())
	}

	// Active Service
	for {
		conn, _ := l.Accept()
		defer conn.Close()

		go jsonrpc.ServeConn(conn)
	}
}

//Register Regist Services
func Register(serviceName, description string, instance interface{}) error {
	ServiceRegisted[serviceName] = Service{
		Name:        serviceName,
		Description: description,
		Instance:    instance,
	}

	return nil
}
