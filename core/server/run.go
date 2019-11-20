package server

import (
	"fmt"
	"os"
	"strings"

	"rpc_server/core/rpc"
	_ "rpc_server/service"
)

func init() {
	// timezone
	if os.Getenv("TZ") != "" {
		os.Setenv("TZ", os.Getenv("TZ"))
	}

	// check env
	if os.Getenv("ENV") == "" || os.Getenv("SERVICE") == "" {
		usage(0, "Please use 'ENV' & 'SERVICE' to star service")
	}
}

// Serve Star server
func Serve() {
	// Regist service
	rpc.Active(os.Getenv("SERVICE"))
}

func usage(exitCode int, extraMessage ...interface{}) {
	builder := new(strings.Builder)
	builder.WriteString("*  SERVICE : Service List")
	builder.WriteRune('\n')
	srvName := "<none>"
	whiteSpace := 0

	for name := range rpc.ServiceRegisted {
		if len(name) > whiteSpace {
			whiteSpace = len(name)
		}
	}

	for name := range rpc.ServiceRegisted {
		srv := rpc.ServiceRegisted[name]
		srvName = name
		cutNum := whiteSpace - len(name)
		cut := ""
		for i := 0; i < cutNum; i++ {
			cut += " "
		}
		builder.WriteString(fmt.Sprintf("		✏ %s %s %s\n", name, cut, srv.Description))
	}

	fmt.Printf(`
	Available ENV：

	*  ENV : Environment
		- docker           Docker
		- local            Local
		- gcp-development  GCP Development
		- gcp-qatest       GCP Quality Assurance Test
		- gcp-pro          GCP Production

	%s

	Example： ENV=local SERVICE=%s go run main.go


`, builder.String(), srvName)

	if len(extraMessage) > 0 {
		fmt.Println("[Server Error]:")
		fmt.Println(extraMessage...)
	}

	os.Exit(exitCode)
}
