package main

import (
	"log/slog"
	"os"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	//cfg := config.MustLoad()
	//
	//log := setupLogger(cfg.Env)
	//
	//log.Info("starting application",
	//	slog.String("env", cfg.Env),
	//	slog.Any("cfg", cfg),
	//	slog.Int("port", cfg.GRPC.Port),
	//)
	////fmt.Println(cfg)
	//log.Debug("debug message")
	//
	//log.Error("error message")
	//
	//log.Warn("warning message")
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log
}
