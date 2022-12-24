package service

import (
	"context"
	"math"

	"github.com/driver005/gateway/models"
)

type Totals struct {
	h   Handler
	ctx context.Context
}

func NewTotals(r Registry) *Totals {
	return &Totals{
		h:   Handler{r: r},
		ctx: context.Background(),
	}
}

// func (c *Totals) GetTotal(cart *models.Cart, order *models.Order) ([]models.Totals, error) {
// 	var m interface{}

// 	if cart != nil {
// 		m = cart
// 	} else if order != nil {
// 		m = order
// 	}

// }

func (c *Totals) GetPaidTotal(order models.Order) int {
	var total int

	for _, value := range order.Payments {
		total += int(value.Amount)
	}

	return total
}

func (c *Totals) GetSwapTotal(order models.Order) int {
	var swapTotal int

	if len(order.Swaps) < 0 {
		for _, value := range order.Swaps {
			swapTotal += int(value.DifferenceDue)
		}
	}

	return swapTotal
}

func (c *Totals) GetShippingMethodTotals(cart *models.Cart, order *models.Order, shippingMethod models.ShippingMethod, opts GetShippingMethodTotalsOptions) (ShippingMethodTotals, error) {
	var m interface{}

	if cart != nil {
		m = cart
	} else if order != nil {
		m = order
	}

	totals := ShippingMethodTotals{
		Price:            int(shippingMethod.Price),
		OriginalTotal:    int(shippingMethod.Price),
		Total:            int(shippingMethod.Price),
		Subtotal:         int(shippingMethod.Price),
		OriginalTaxTotal: 0,
		TaxTotal:         0,
		TaxLines:         shippingMethod.TaxLines,
	}

	if opts.IncludeTax {
		if isOrder(m) && m.(models.Order).TaxRate != 0 {
			totals.OriginalTaxTotal = int(math.Round(float64(float32(totals.Price) * (m.(models.Order).TaxRate / 100.0))))
			totals.TaxTotal = int(math.Round(float64(float32(totals.Price) * (m.(models.Order).TaxRate / 100.0))))
		} else if len(totals.TaxLines) == 0 {

		}
	}
	return totals, nil
}
