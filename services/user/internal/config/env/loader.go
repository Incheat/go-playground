package envconfig

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// errMissingEnv is the error returned when a required environment variable is missing.
var errMissingEnv = errors.New("missing env var")

// Load loads the configuration from the environment variables.
func Load() (*Config, error) {
	env := getString("ENV") // if missing, it will be ""

	userInternalPort, err := getIntRequired("USER_INTERNAL_PORT")
	if err != nil {
		return nil, err
	}

	userCorsPublicAllowedOrigins := parseOrigins(getString("USER_CORS_PUBLIC_ALLOWED_ORIGINS"))
	userCorsPublicAllowCredentials, err := getBoolRequired("USER_CORS_PUBLIC_ALLOW_CREDENTIALS")
	if err != nil {
		return nil, err
	}

	userCorsInternalAllowedOrigins := parseOrigins(getString("USER_CORS_INTERNAL_ALLOWED_ORIGINS"))
	userCorsInternalAllowCredentials, err := getBoolRequired("USER_CORS_INTERNAL_ALLOW_CREDENTIALS")
	if err != nil {
		return nil, err
	}

	userMySQLHost := getString("USER_MYSQL_HOST")
	userMySQLUser := getString("USER_MYSQL_USER")
	userMySQLPassword := getString("USER_MYSQL_PASSWORD")
	userMySQLDBName := getString("USER_MYSQL_DB_NAME")
	userMySQLMaxOpenConns, err := getIntRequired("USER_MYSQL_MAX_OPEN_CONNS")
	if err != nil {
		return nil, err
	}
	userMySQLMaxIdleConns, err := getIntRequired("USER_MYSQL_MAX_IDLE_CONNS")
	if err != nil {
		return nil, err
	}
	userMySQLConnMaxLifetime, err := getIntRequired("USER_MYSQL_CONN_MAX_LIFETIME")
	if err != nil {
		return nil, err
	}

	cfg := &Config{
		Env: EnvName(env),
		Server: Server{
			InternalPort: Port(userInternalPort),
		},
		CORS: CORS{
			Internal: CORSRule{
				AllowedOrigins:   userCorsInternalAllowedOrigins,
				AllowCredentials: userCorsInternalAllowCredentials,
			},
			Public: CORSRule{
				AllowedOrigins:   userCorsPublicAllowedOrigins,
				AllowCredentials: userCorsPublicAllowCredentials,
			},
		},
		MySQL: MySQL{
			Host:            userMySQLHost,
			User:            userMySQLUser,
			Password:        userMySQLPassword,
			DBName:          userMySQLDBName,
			MaxOpenConns:    userMySQLMaxOpenConns,
			MaxIdleConns:    userMySQLMaxIdleConns,
			ConnMaxLifetime: userMySQLConnMaxLifetime,
		},
	}

	// Optional sanity checks (keep or remove as you like)
	if err := validate(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}

func getString(name string) string {
	return strings.TrimSpace(os.Getenv(name))
}

func getIntRequired(name string) (int, error) {
	raw := getString(name)
	if raw == "" {
		return 0, fmt.Errorf("%s: %w", name, errMissingEnv)
	}
	v, err := strconv.Atoi(raw)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", name, err)
	}
	return v, nil
}

func getBoolRequired(name string) (bool, error) {
	raw := getString(name)
	if raw == "" {
		return false, fmt.Errorf("%s: %w", name, errMissingEnv)
	}
	v, err := strconv.ParseBool(raw)
	if err != nil {
		return false, fmt.Errorf("%s: %w", name, err)
	}
	return v, nil
}

// parseOrigins supports:
// - "" => nil
// - "*" => []string{"*"}  (so CORS layer can treat it as allow-all)
// - "a,b,c" => []string{"a","b","c"} (trimmed, empties removed)
func parseOrigins(env string) []string {
	env = strings.TrimSpace(env)
	if env == "" {
		return nil
	}
	if env == "*" {
		return []string{"*"}
	}

	parts := strings.Split(env, ",")
	out := make([]string, 0, len(parts))
	for _, p := range parts {
		p = strings.TrimSpace(p)
		if p == "" {
			continue
		}
		out = append(out, p)
	}
	return out
}

func validate(cfg *Config) error {
	if cfg.Server.InternalPort <= 0 || cfg.Server.InternalPort > 65535 {
		return fmt.Errorf("USER_PUBLIC_PORT: must be between 1 and 65535")
	}

	return nil
}
