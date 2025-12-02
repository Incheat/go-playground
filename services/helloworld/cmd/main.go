package main

import (
	"fmt"
	"github.com/incheat/go-playground/config"
	"os"
)

func main() {
	cfg := config.MustLoad()

	fmt.Printf("ENV: %s\n", os.Getenv("APP_ENV"))
	fmt.Printf("Server port: %d\n", cfg.Server.Port)
	fmt.Printf("DB host: %s\n", cfg.Database.Host)
}