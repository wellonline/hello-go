package checkout

import (
	"fmt"

	"github.com/stripe/stripe-go/v76"
)

type cart struct {
	TransactionID string
	items         []item
	// TODO: VAT, country
}

type item struct {
	ID    string
	Price int64
	// TODO: Study more about price in stripe
}

func toStripeCheckout(c cart) *stripe.CheckoutSessionParams {
	basePath := fmt.Sprintf("https://well.online/checkout?transactionId=%s&status=", c.TransactionID)
	csp := &stripe.CheckoutSessionParams{
		Mode:       stripe.String(string(stripe.CheckoutSessionModePayment)),
		SuccessURL: stripe.String(basePath + "success"),
		CancelURL:  stripe.String(basePath + "cancel"),
	}

	for _, item := range c.items {
		csp.LineItems = append(csp.LineItems, &stripe.CheckoutSessionLineItemParams{
			Price:    stripe.String(item.ID),
			Quantity: stripe.Int64(item.Price),
		})
	}

	return csp
}
