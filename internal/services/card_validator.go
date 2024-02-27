package services

import (
	"errors"

	"github.com/vadimpk/card-validator-golang/internal/domain"
)

type cardValidatorService struct {
	validatorFactory *domain.CardValidatorFactory
}

func NewCardValidatorService() CardValidatorService {
	return &cardValidatorService{
		validatorFactory: domain.NewCardValidatorFactory(),
	}
}

func (s *cardValidatorService) ValidateCard(card *domain.Card, validatorType domain.CardValidatorType) (*ValidateCardOutput, error) {
	cardValidator := s.validatorFactory.CreateCardValidator(validatorType)
	if cardValidator == nil {
		return nil, errors.New("unknown validator type")
	}

	valid, reason := cardValidator.Validate(card)
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
