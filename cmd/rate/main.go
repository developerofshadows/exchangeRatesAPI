package main

import (
	"context"
	"exchangeRatesAPI/internal/repository"
	httpT "exchangeRatesAPI/internal/transport/http"
	"exchangeRatesAPI/pkg/config"
	"exchangeRatesAPI/pkg/logger"
	"os/signal"
	"syscall"
)

// @title ExchangeRate API
// @version 1.0
// @description API server for checking exchange rates

// @host localhost:3443
// @BasePath  /

func main() {
	// Loading config file
	cfg := config.MustLoadConfig()

	// Initializing logger
	logger.InitLogger(cfg.Logger.LoggerLevel)

	// Connecting to DB
	repository.InitRepository(&cfg.Database)

	// Initializing HTTP server
	httpServer := httpT.InitializeServer(&cfg.HTTP)

	// Creating listener for SIGINT,SIGTERM system events that triggers graceful shutdown
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// Starting HTTP server
	httpServer.StartServer()

	// On SIGINT,SIGTERM start shutdown
	<-ctx.Done()
	logger.Logger.Warn("Gracefully shutdown started")
	httpServer.GracefulShutdown()

	logger.Logger.Debug("Exit")
}
