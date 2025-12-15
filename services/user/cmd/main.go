// Package main defines the main function for the user service.
package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	servergen "github.com/incheat/go-playground/services/user/internal/api/oapi/gen/private/server"
	envconfig "github.com/incheat/go-playground/services/user/internal/config/env"
	userhandler "github.com/incheat/go-playground/services/user/internal/handler/http"
	chimiddleware "github.com/incheat/go-playground/services/user/internal/middleware/chi"
	userrepo "github.com/incheat/go-playground/services/user/internal/repository/mysql"
	userservice "github.com/incheat/go-playground/services/user/internal/service/user"
	nethttpmiddleware "github.com/oapi-codegen/nethttp-middleware"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
)

func main() {

	cfg, err := envconfig.Load()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}
	logger := initLogger(envconfig.EnvName(cfg.Env))

	logger.Info("Starting user service", zap.String("env", string(cfg.Env)))
	logger.Info("Http server internal port", zap.Int("port", int(cfg.Server.InternalPort)))

	// Get OpenAPI definition from embedded spec
	openAPISpec, err := servergen.GetSwagger()
	if err != nil {
		log.Fatalf("Error loading OpenAPI spec: %v", err)
	}

	// HTTP router
	router := chi.NewRouter()
	router.Use(nethttpmiddleware.OapiRequestValidatorWithOptions(
		openAPISpec,
		chimiddleware.NewValidatorOptions(chimiddleware.ValidatorConfig{
			ProdMode: cfg.Env == envconfig.EnvProd,
		}),
	))
	// router.Use(chimiddleware.PathBasedCORS(convertCORSRules(cfg)))
	router.Use(chimiddleware.RequestID())
	router.Use(chimiddleware.ZapLogger(logger))
	router.Use(chimiddleware.ZapRecovery(logger))

	// Initialize MySQL connection
	dbDSN := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", cfg.MySQL.User, cfg.MySQL.Password, cfg.MySQL.Host, cfg.MySQL.DBName)
	logger.Info("Initializing MySQL connection", zap.String("dsn", dbDSN))
	dbConn, err := sql.Open("mysql", dbDSN)
	if err != nil {
		log.Fatalf("Error opening MySQL connection: %v", err)
	}
	dbConn.SetMaxOpenConns(cfg.MySQL.MaxOpenConns)
	dbConn.SetMaxIdleConns(cfg.MySQL.MaxIdleConns)
	dbConn.SetConnMaxLifetime(time.Duration(cfg.MySQL.ConnMaxLifetime) * time.Second)
	if err != nil {
		log.Fatalf("Error opening MySQL connection: %v", err)
	}
	defer func() {
		if err := dbConn.Close(); err != nil {
			logger.Warn("Failed to close MySQL connection", zap.Error(err))
		}
	}()

	// Check if the connection is working
	if err := dbConn.Ping(); err != nil {
		log.Fatalf("Error pinging MySQL: %v", err)
	}

	// user components
	userRepository := userrepo.NewUserRepository(dbConn)

	userService := userservice.New(userRepository)
	userImpl := userhandler.New(userService)

	strict := servergen.NewStrictHandler(userImpl, nil)
	apiHandler := servergen.HandlerFromMux(strict, router)

	var g errgroup.Group

	g.Go(func() error {
		return http.ListenAndServe(fmt.Sprintf(":%d", cfg.Server.InternalPort), apiHandler)
	})

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}

}

func initLogger(env envconfig.EnvName) *zap.Logger {
	switch env {
	case envconfig.EnvDev, envconfig.EnvStaging:
		return zap.Must(zap.NewDevelopment())
	default:
		return zap.Must(zap.NewProduction())
	}
}
