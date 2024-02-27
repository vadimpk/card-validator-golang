package main

import (
	"fmt"
	"net/http"

	"github.com/vadimpk/card-validator-golang/config"
	httpcontroller "github.com/vadimpk/card-validator-golang/internal/handlers/http"
	"github.com/vadimpk/card-validator-golang/internal/services"
	"github.com/vadimpk/card-validator-golang/pkg/logging"
)

func main() {
	cfg := config.Get()
	logger := logging.New(cfg.Log.Level)

	srvs := services.Services{
		CardValidatorService: services.NewCardValidatorService(),
	}

	h := httpcontroller.NewHandler(httpcontroller.Options{
		Services: srvs,
		Logger:   logger,
		Config:   cfg,
	})

	if err := http.ListenAndServe(fmt.Sprint(":", cfg.HTTP.Port), h); err != nil {
		logger.Fatal("failed to start server", "err", err)
	}
}
