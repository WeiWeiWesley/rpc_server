package rpc

// Service 微服務
type Service struct {
	Name        string      // 名稱
	Description string      // 服務描述
	Instance    interface{} // 實例
}

// ServiceRegisted 註冊清單
var ServiceRegisted = map[string]Service{}
