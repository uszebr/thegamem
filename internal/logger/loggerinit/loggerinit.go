package loggerinit

import (
	"log"
	"log/slog"
	"os"
	"sync"

	"github.com/uszebr/thegamem/internal/config"
)

var (
	once     sync.Once
	instance *slog.Logger
)

// might be removed sinse logger instance assigned to the default logger in slog
func GetLogger() *slog.Logger {
	if instance == nil {
		log.Fatalf("logger is not initialized")
	}
	return instance
}

func MustInitLogger(env config.Env) *slog.Logger {
	once.Do(func() {
		var clog *slog.Logger
		switch env {
		case config.EnvLocal:
			clog = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
		case config.EnvDev:
			clog = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
		case config.EnvProd:
			clog = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
		default:
			log.Fatalf("Unknown environment: %s", env)
		}
		instance = clog
		slog.SetDefault(clog)
	})
	return instance
}
