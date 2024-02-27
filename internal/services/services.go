package services

import "github.com/vadimpk/card-validator-golang/internal/domain"

type Services struct {
	CardValidatorService
}

type CardValidatorService interface {
	// ValidateCard validates a card. It returns a ValidateCardOutput that contains the result of the validation.
	ValidateCard(card *domain.Card, validator domain.CardValidator) (*ValidateCardOutput, error)
}

type ValidateCardOutput struct {
	IsValid bool
	Reason  string
}
