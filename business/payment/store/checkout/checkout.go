// Package checkout contains all apis related to stripe checkout
package checkout

import (
	"github.com/stripe/stripe-go/v76/checkout/session"
)

// Checkout api
// when able connect to always return redirect url
func Checkout(c cart, price int64) (string, error) {
	params := toStripeCheckout(c)

	s, err := session.New(params)

	if err != nil {
		return "", err
	}

	return s.URL, nil
}
