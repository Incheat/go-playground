package config

type EnvName string

const (
    EnvDev     EnvName = "dev"
    EnvStaging EnvName = "staging"
    EnvProd    EnvName = "prod"
)

// Config is the configuration for the application.
type Config struct {

	Env EnvName `koanf:"env"`

	Server struct {
		Port int `koanf:"port"`
	} `koanf:"server"`

	CORS struct {
		Rules []CORSRule `koanf:"rules"`
	} `koanf:"cors"`

}

// CORSRule is a rule that defines the CORS configuration for a specific path.
type CORSRule struct {
	Path           string   `koanf:"path"`
	AllowedOrigins []string `koanf:"allowed_origins"`
}