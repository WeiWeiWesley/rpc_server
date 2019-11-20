package boot

import (
	"log"
	"os"
	"strings"

	"rpc_server/scheme"

	"github.com/BurntSushi/toml"
)

var conf *scheme.Config

//LoadConfig Load config by env & service's name
func LoadConfig() *scheme.Config {
	if conf != nil {
		return conf
	}

	env := os.Getenv("ENV")
	service := "default"

	// default config
	cf := "config/env/" + env + "/" + service + ".toml"
	if _, err := toml.DecodeFile(cf, &conf); err != nil {
		log.Fatalf("No default config file: '%s'", cf)
	}

	// service's config
	if spf := strings.ToLower(os.Getenv("SERVICE")); len(spf) > 0 {
		service = spf
	}
	cf = "config/env/" + env + "/" + service + ".toml"
	toml.DecodeFile(cf, &conf)

	conf.ENV = env
	return conf
}
