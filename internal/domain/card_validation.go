package domain

import (
	"errors"
	"strings"
	"time"
)

// CardValidator represents a strategy interface for validating a card.
type CardValidator interface {
	// Validate validates a card. Returns true if the card is valid and false otherwise.
	// If the card is invalid, it returns a message explaining why the card is invalid.
	Validate(card *Card) (bool, string)
}

// LiveCardValidator is a card validator that validates real-life cards.
// It encapsulates a list of validators that are used to validate a card with different criteria.
type LiveCardValidator struct {
	validators []CardValidator
}

func NewLiveCardValidator() *LiveCardValidator {
	return &LiveCardValidator{
		validators: []CardValidator{
			&luhnAlgorithmValidator{},
			&lengthValidator{},
			&expirationDateValidator{},
		},
	}
}

func (lcv *LiveCardValidator) Validate(card *Card) (bool, string) {
	for _, validator := range lcv.validators {
		valid, message := validator.Validate(card)
		if !valid {
			return false, message
		}
	}

	return true, ""
}

// luhnAlgorithmValidator is a validator that uses the Luhn algorithm to validate a card.
// The implementation was found online.
type luhnAlgorithmValidator struct {
}

func (*luhnAlgorithmValidator) Validate(card *Card) (bool, string) {
	p := len(card.Number) % 2
	sum, err := calculateLuhnSum(card.Number, p)
	if err != nil {
		return false, err.Error()
	}

	// If the total modulo 10 is not equal to 0, then the number is invalid.
	if sum%10 != 0 {
		return false, "Invalid card number"
	}

	return true, ""
}

const (
	asciiZero = 48
	asciiTen  = 57
)

func calculateLuhnSum(number string, parity int) (int64, error) {
	var sum int64
	for i, d := range number {
		if d < asciiZero || d > asciiTen {
			return 0, errors.New("invalid digit")
		}

		d = d - asciiZero
		// Double the value of every second digit.
		if i%2 == parity {
			d *= 2
			// If the result of this doubling operation is greater than 9.
			if d > 9 {
				// The same final result can be found by subtracting 9 from that result.
				d -= 9
			}
		}

		// Take the sum of all the digits.
		sum += int64(d)
	}

	return sum, nil
}

type lengthValidator struct {
}

func (*lengthValidator) Validate(card *Card) (bool, string) {
	l := len(card.Number)

	// Visa
	if strings.HasPrefix(card.Number, "4") {
		return l == 13 || l == 16, "Invalid card number length. Must be 13 or 16"
	}

	// MasterCard
	if strings.HasPrefix(card.Number, "5") {
		return l == 16, "Invalid card number length. Must be 16"
	}

	// Other ...

	// Base case: for unknown card types we return true since we cannot check the length.
	return true, ""
}

// expirationDateValidator is a validator that checks if the expiration date of a card is in the future.
type expirationDateValidator struct {
}

func (*expirationDateValidator) Validate(card *Card) (bool, string) {
	// Parse the expiration month and year.
	expMonth, err := time.Parse("01", card.ExpMonth)
	if err != nil {
		return false, "Invalid expiration month"
	}

	expYear, err := time.Parse("2006", card.ExpYear)
	if err != nil {
		return false, "Invalid expiration year"
	}

	// Create a time.Time object for the expiration date.
	// We add 1 to month because we cannot check the exact day, so we assume the expiration date is the last day of the month,
	// which is the same as the first day of the next month.
	expiration := time.Date(expYear.Year(), expMonth.Month()+1, 1, 0, 0, 0, 0, time.UTC)

	if !time.Now().Before(expiration) {
		return false, "Invalid expiration date. Must be in the future"
	}

	return true, ""
}

// TestCardValidator is a card validator that validates test cards.
type TestCardValidator struct {
	expirationDateValidator
}

func NewTestCardValidator() *TestCardValidator {
	return &TestCardValidator{}
}

func (tcv *TestCardValidator) Validate(card *Card) (bool, string) {
	if card.Number != "4242424242424242" {
		return false, "Invalid card number. Must be 4242424242424242"
	}

	return tcv.expirationDateValidator.Validate(card)
}
