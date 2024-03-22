package logger

import (
	"github.com/lmittmann/tint"
	"log/slog"
	"os"
	"time"
)

const (
	envLocal = "dev"
	envProd  = "prod"
)

var Logger *slog.Logger

func InitLogger(env string) {
	switch env {
	case envLocal:
		Logger = slog.New(tint.NewHandler(os.Stdout, &tint.Options{Level: slog.LevelDebug, TimeFormat: time.Kitchen}))
	case envProd:
		Logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelInfo,
		}))

	}

}
