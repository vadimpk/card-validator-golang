package services

import (
	"github.com/vadimpk/card-validator-golang/internal/domain"
)

type cardValidatorService struct {
}

func NewCardValidatorService() CardValidatorService {
	return &cardValidatorService{}
}

func (s *cardValidatorService) ValidateCard(card *domain.Card, validator domain.CardValidator) (*ValidateCardOutput, error) {
	valid, reason := validator.Validate(card)
	if !valid {
		return &ValidateCardOutput{
			IsValid: false,
			Reason:  reason,
		}, nil
	}

	return &ValidateCardOutput{
		IsValid: true,
	}, nil
}
