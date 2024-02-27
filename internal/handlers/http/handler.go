package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vadimpk/card-validator-golang/config"
	"github.com/vadimpk/card-validator-golang/internal/services"
	"github.com/vadimpk/card-validator-golang/pkg/logging"
)

type handler struct {
	services services.Services
	logger   logging.Logger
	config   *config.Config
}

type Options struct {
	Services services.Services
	Logger   logging.Logger
	Config   *config.Config
}

func NewHandler(opts Options) http.Handler {
	h := &handler{
		services: opts.Services,
		logger:   opts.Logger,
		config:   opts.Config,
	}

	r := gin.New()
	r.Use(gin.Recovery())

	api := r.Group("/api")
	{
		api.POST("/validate", h.validateLiveCard)
	}

	test := r.Group("/test")
	{
		test.POST("/validate", h.validateTestCard)
	}

	return r
}

type responseError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

const (
	unexpectedErrorCode = 1000
	invalidCardCode     = 4000
)
