package rpc

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"

	"rpc_server/core/boot"
)

func init() {
	s := boot.LoadConfig().Service

	ServiceRegisted[s.Name] = Service {
		Name: s.Name,
		Description: s.Description,
	}
}

// Register 服務註冊
func Register(name string) {
	if service, ok := ServiceRegisted[name]; ok {
		// 註冊服務
		err := rpc.RegisterName(name, service.Instance)
		if err != nil {
			log.Println("[rpc_error]:", err.Error())
		}
	} else {
		log.Fatalf("[Fatal error] Service '%s' not exist", name)
	}

	log.Printf("Start service: %s; success", name)

	// 建立監聽
	l, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Println("[rpc_error]:", err.Error())
	}

	// 啟動服務
	for {
		conn, _ := l.Accept()
		defer conn.Close()

		go jsonrpc.ServeConn(conn)
	}
}
