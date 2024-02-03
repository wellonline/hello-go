package checkout

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTransformFromCheckoutToStripCheckoutSessionParams(t *testing.T) {
	c := cart{
		items: []item{
			item{"1", 2},
			item{"3", 4},
		},
	}

	scp := toStripeCheckout(c)

	assert.Equal(t, 2, len(scp.LineItems), "should have 2 line of items in checkout params after transform")
	// TODO: Finish me...
}
