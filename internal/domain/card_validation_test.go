package domain

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestLiveCardValidator_Validate(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		card     Card
		expected bool
	}{
		{
			name: "positive: valid card",
			card: Card{
				Number:   "4111111111111111",
				ExpMonth: "01",
				ExpYear:  "2028",
			},
			expected: true,
		},
		{
			name: "negative: invalid card (luhn)",
			card: Card{
				Number: "4242424242424241",
			},
			expected: false,
		},
		{
			name: "negative: invalid card (exp)",
			card: Card{
				Number:   "4242424242424242",
				ExpMonth: "01",
				ExpYear:  "2020",
			},
			expected: false,
		},
		{
			name: "negative: invalid card (length)",
			card: Card{
				Number: "424242424242424",
			},
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			validator := NewLiveCardValidator()

			valid, _ := validator.Validate(&tc.card)
			assert.Equal(t, tc.expected, valid)
		})
	}
}

func TestLuhnAlgorithmValidator_Validate(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		card     Card
		expected bool
	}{
		{
			name: "positive: valid card",
			card: Card{
				Number: "4111111111111111",
			},
			expected: true,
		},
		{
			name: "negative: invalid card",
			card: Card{
				Number: "4242424242424241",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			validator := luhnAlgorithmValidator{}

			valid, _ := validator.Validate(&tc.card)
			assert.Equal(t, tc.expected, valid)
		})
	}
}

func TestCardExpirationValidator_Validate(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		card     Card
		expected bool
	}{
		{
			name: "positive: valid card",
			card: Card{
				ExpMonth: "01",
				ExpYear:  "2028",
			},
			expected: true,
		},
		{
			name: "negative: invalid card (in the past)",
			card: Card{
				ExpMonth: "01",
				ExpYear:  "2020",
			},
			expected: false,
		},
		{
			name: "negative: invalid card (not a valid month)",
			card: Card{
				ExpMonth: "13",
			},
			expected: false,
		},
		{
			name: "negative: invalid card (not a valid month)",
			card: Card{
				ExpMonth: "January",
			},
			expected: false,
		},
		{
			name: "negative: invalid card (not a valid year)",
			card: Card{
				ExpMonth: "01",
				ExpYear:  "-190",
			},
			expected: false,
		},
		{
			name: "negative: invalid card (not a valid year)",
			card: Card{
				ExpMonth: "01",
				ExpYear:  "year",
			},
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			validator := expirationDateValidator{}

			valid, _ := validator.Validate(&tc.card)
			assert.Equal(t, tc.expected, valid)
		})
	}
}

func TestLengthValidator_Validate(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		card     Card
		expected bool
	}{
		{
			name: "positive: valid card length (visa)",
			card: Card{
				Number: "4000000000000000",
			},
			expected: true,
		},
		{
			name: "positive: valid card length (visa)",
			card: Card{
				Number: "4000000000000",
			},
			expected: true,
		},
		{
			name: "negative: invalid card length (visa)",
			card: Card{
				Number: "40000000000000",
			},
			expected: false,
		},
		{
			name: "positive: valid card length (mastercard)",
			card: Card{
				Number: "5000000000000000",
			},
			expected: true,
		},
		{
			name: "negative: invalid card length (mastercard)",
			card: Card{
				Number: "500000000000000",
			},
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			validator := lengthValidator{}

			valid, _ := validator.Validate(&tc.card)
			assert.Equal(t, tc.expected, valid)
		})
	}
}

func TestTestCardValidator_Validate(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		card     Card
		expected bool
	}{
		{
			name: "positive: valid card",
			card: Card{
				Number:   "4242424242424242",
				ExpMonth: "01",
				ExpYear:  "2028",
			},
			expected: true,
		},
		{
			name: "negative: invalid card",
			card: Card{
				Number: "4242424242424241",
			},
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			validator := NewTestCardValidator()

			valid, _ := validator.Validate(&tc.card)
			assert.Equal(t, tc.expected, valid)
		})
	}
}
