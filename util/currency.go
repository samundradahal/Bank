package util

const (
	USD = "USD"
	EUR = "EUR"
	NPR = "NPR"
	INR = "INR"
)

func SuportedCurrency(currency string) bool {
	switch currency {
	case USD, EUR, NPR, INR:
		return true
	}
	return false
}
