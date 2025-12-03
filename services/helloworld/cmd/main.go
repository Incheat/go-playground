package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	servergen "github.com/incheat/go-playground/services/helloworld/internal/api/gen/server"
	"github.com/incheat/go-playground/services/helloworld/internal/config"
	"github.com/incheat/go-playground/services/helloworld/internal/handler"
)

func main() {
	cfg := config.MustLoad()

	fmt.Printf("ENV: %s\n", os.Getenv("APP_ENV"))
	fmt.Printf("Server port: %d\n", cfg.Server.Port)
	fmt.Printf("DB host: %s\n", cfg.Database.Host)

	r := gin.Default()
	srv := handler.NewServer()
	handler := servergen.NewStrictHandler(srv, nil)
	servergen.RegisterHandlers(r, handler)

	s := &http.Server{
		Handler: r,
		Addr:    fmt.Sprintf(":%d", cfg.Server.Port),
	}

	log.Fatal(s.ListenAndServe())
	
}