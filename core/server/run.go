package server

import (
	"fmt"
	"os"
	"strings"

	"rpc_server/core/rpc"
	"rpc_server/core/boot"
	_ "rpc_server/service"
)

func init() {
	// timezone
	if os.Getenv("TZ") == "" {
		os.Setenv("TZ", "Asia/Taipei")
	}

	// check env
	if os.Getenv("ENV") == "" || os.Getenv("SERVICE") == "" {
		usage(0, "Please use 'ENV' & 'SERVICE' to star service")
	}
}

// Serve Star server
func Serve() {
	conf := boot.LoadConfig()

	// Regist service
	rpc.Active(conf.Service.Name)
}

func usage(exitCode int, extraMessage ...interface{}) {
	builder := new(strings.Builder)
	builder.WriteString("⚙  SERVICE : 服務清單")
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
	可用環境變數：

	*  ENV : 運行環境
		- docker           容器開發
		- local            本機開發
		- gcp-development  GCP開發站
		- gcp-qatest       GCP測試站
		- gcp-production   GCP正式站

	%s

	範例： ENV=local SERVICE=%s go run main.go


`, builder.String(), srvName)

	if len(extraMessage) > 0 {
		fmt.Println("[Server Error]:")
		fmt.Println(extraMessage...)
	}

	os.Exit(exitCode)
}
