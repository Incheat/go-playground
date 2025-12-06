package constant

const (
	EnvKey        = "env" // APP_ENV => env => ex: "test" / "staging" / "prod"
	EnvPrefix     = "APP_"    // APP_SERVER_PORT => server.port
	EnvConfigDir  = "config"
	EnvConfigTmpl = "config.%s.yaml"
	APIResponseVersionV1 = "v1"
)