package app

type Config struct {
	BindAddr    string
	DatabaseURL string
	SessionKey  string
	TokenSecret string
}

func NewConfig(port string, dburl string) *Config {
	// TODO: setup session key and secret
	return &Config{
		BindAddr:		":" + port,
		SessionKey:		"sessionkey",
		TokenSecret:	"secret",
		DatabaseURL:	dburl,
	}
}