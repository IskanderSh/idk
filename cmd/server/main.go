package main

import (
	"awesomeProject/internal/app"
	"awesomeProject/internal/config"
	"fmt"
	"log/slog"
	"os"
)

func main() {
	// init logger
	log := slog.New(slog.NewTextHandler(os.Stdout, nil))
	log.Info("logger initialized successfully")

	// init config
	cfg := config.MustLoad()
	log.Info(fmt.Sprintf("config: %+v", cfg))

	// init and start all deps
	if err := app.StartService(log, cfg); err != nil {
		panic(err)
	}
}
