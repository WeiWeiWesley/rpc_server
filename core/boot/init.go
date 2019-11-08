package boot

import (
	"log"
	"os"
	"strings"

	"rpc_server/scheme"

	"github.com/BurntSushi/toml"
)

var conf *scheme.Config

//LoadConfig 載入 config
func LoadConfig() *scheme.Config {
	if conf != nil {
		return conf
	}

	env := os.Getenv("ENV")
	service := "default"

	// 載入預設 config
	cf := "config/env/" + env + "/" + service + ".toml"
	if _, err := toml.DecodeFile(cf, &conf); err != nil {
		log.Fatalf("No default config file: '%s'", cf)
	}

	// 載入特定 config
	if spf := strings.ToLower(os.Getenv("SERVICE")); len(spf) > 0 {
		service = spf
	}
	cf = "config/env/" + env + "/" + service + ".toml"
	if _, err := toml.DecodeFile(cf, &conf); err != nil {
		log.Fatalf("No default config file: '%s'", cf)
	}

	conf.ENV = env
	return conf
}
