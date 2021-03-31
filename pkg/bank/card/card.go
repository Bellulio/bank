package card

import (
	"bank/pkg/bank/types"
)


func PaymentSources(cards []types.Card)[]types.PaymentSource {
	
	var payments_source []types.PaymentSource

	for _, card := range cards {
		if card.Balance <= 0 || !card.Active {
			continue
		} 
		payments_source = append(payments_source, types.PaymentSource {
			Type: "card",
			Number: card.PAN,
			Balance: card.Balance,
		})
	}
	
	return payments_source
}
