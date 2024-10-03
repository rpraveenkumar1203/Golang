package utils

import (
	"testing"
	"unicode"

	"github.com/stretchr/testify/require"
)

func TestRandomInt(t *testing.T) {
	min, max := int64(1), int64(10)
	for i := 0; i < 100; i++ {
		result := RandomInt(min, max)
		require.GreaterOrEqual(t, result, min, "RandomInt result should be >= min")
		require.LessOrEqual(t, result, max, "RandomInt result should be <= max")
	}
}

func TestRandomString(t *testing.T) {
	n := 10
	result := RandomString(n)
	require.Len(t, result, n, "RandomString result length should match the requested length")

	for _, char := range result {
		require.True(t, unicode.IsLower(rune(char)), "RandomString result should only contain lowercase alphabetic characters")
	}
}

func TestRandomOwner(t *testing.T) {
	result := RandomOwner()
	require.Len(t, result, 6, "RandomOwner result should be a 6 character string")
	for _, char := range result {
		require.True(t, unicode.IsLower(rune(char)), "RandomOwner result should only contain lowercase alphabetic characters")
	}
}

func TestRandomMoney(t *testing.T) {
	for i := 0; i < 100; i++ {
		result := RandomMoney()
		require.GreaterOrEqual(t, result, int64(0), "RandomMoney result should be >= 0")
		require.LessOrEqual(t, result, int64(1000), "RandomMoney result should be <= 1000")
	}
}

func TestRandomCurrency(t *testing.T) {
	validCurrencies := []string{"USD", "INR", "EUR"}
	result := RandomCurrency()
	require.Contains(t, validCurrencies, result, "RandomCurrency result should be one of USD, INR, EUR")
}
