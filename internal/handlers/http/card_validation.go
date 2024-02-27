package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vadimpk/card-validator-golang/internal/domain"
	"github.com/vadimpk/card-validator-golang/pkg/logging"
)

type validateCardResponseBody struct {
	IsValid bool           `json:"valid"`
	Error   *responseError `json:"error,omitempty"`
}

func (h *handler) validateLiveCard(c *gin.Context) {
	h.validateCard(c, h.logger.Named("validateLiveCard"), domain.LiveCardValidatorType)
}

func (h *handler) validateTestCard(c *gin.Context) {
	h.validateCard(c, h.logger.Named("validateTestCard"), domain.TestCardValidatorType)
}

func (h *handler) validateCard(c *gin.Context, logger logging.Logger, validatorType domain.CardValidatorType) {
	var card domain.Card
	if err := c.ShouldBindJSON(&card); err != nil {
		logger.Error("failed to bind json", "err", err, "request", c.Request.Body)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	logger = logger.With("card", card)

	output, err := h.services.ValidateCard(&card, validatorType)
	if err != nil {
		h.logger.Error("failed to validate card", "err", err)
		c.JSON(http.StatusInternalServerError, responseError{
			Code:    unexpectedErrorCode,
			Message: "internal server error",
		})
		return
	}
	logger = logger.With("output", output)
	logger.Info("card validated successfully")

	if !output.IsValid {
		c.JSON(http.StatusOK, validateCardResponseBody{
			IsValid: false,
			Error: &responseError{
				Code:    invalidCardCode,
				Message: output.Reason,
			},
		})
		return
	}

	c.JSON(http.StatusOK, validateCardResponseBody{
		IsValid: true,
	})
}
