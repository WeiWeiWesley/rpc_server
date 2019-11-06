package scheme

// Config 設定檔
type Config struct {
	ENV     string
	Service struct {
		Name string `toml:"name"`
		Env  string `toml:"env"`
	} `toml:"service"`
}
