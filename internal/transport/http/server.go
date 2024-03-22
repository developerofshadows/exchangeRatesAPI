package httpT

import (
	"context"
	"errors"
	"exchangeRatesAPI/internal/models"
	"exchangeRatesAPI/pkg/logger"
	"fmt"
	"net/http"
)

type APIServer struct {
	httpServer *http.Server
}

func InitializeServer(config *models.HTTP) *APIServer {
	var as APIServer

	mux := http.NewServeMux()
	InitializeRouter(mux)

	as.httpServer = &http.Server{
		Addr:    fmt.Sprintf("0.0.0.0:%d", config.Port),
		Handler: mux,
	}
	logger.Logger.Debug("Successfully initialized server")
	return &as
}

func (s *APIServer) StartServer() {
	go func() {
		logger.Logger.Info(fmt.Sprintf("HTTP server started on port: %s", s.httpServer.Addr[len(s.httpServer.Addr)-4:]))
		if err := s.httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Logger.Error("HTTP server ListenAndServe Error: %v", err)
		}
	}()
}

func (s *APIServer) GracefulShutdown() {
	logger.Logger.Debug("Gracefully shutdown HTTP started")
	if err := s.httpServer.Shutdown(context.TODO()); err != nil {
		logger.Logger.Error(fmt.Sprintf("HTTP Server Shutdown Error: %v", err))
	}

	logger.Logger.Debug("HTTP Server shutdown finished")
}
