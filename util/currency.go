package util

// constants for all supported currencies
const (
	USD  = "USD"
	EUR  = "EUR"
	RUB  = "RUB"
	YUAN = "YUAN"
)

func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, EUR, RUB, YUAN:
		return true
	}
	return false
}
