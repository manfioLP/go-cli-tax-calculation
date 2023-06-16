package utils

type Order struct {
	Operation string  `json:"operation"`
	UnitCost  float64 `json:"unit-cost"`
	Quantity  int     `json:"quantity"`
}

type TaxResponse struct {
	Tax float64 `json:"tax"`
}

func calculateWeightedAverage(currentWeightedAverage float64, currentQuantity int, purchasePrice float64, purchasedQuantity int) float64 {
	return ((float64(currentQuantity) * currentWeightedAverage) + (float64(purchasedQuantity) * purchasePrice)) / float64(currentQuantity+purchasedQuantity)
}

func CalculateTax(orders []Order) []TaxResponse {
	const taxThreshold = 20000.00

	buy := func(op Order) bool {
		return op.Operation == "buy"
	}

	sell := func(op Order) bool {
		return op.Operation == "sell"
	}

	calculateWeightedAverage := func(currentWeightedAverage float64, currentQuantity int, purchasePrice float64, purchasedQuantity int) float64 {
		return (float64(currentQuantity) * currentWeightedAverage) + (float64(purchasedQuantity)*purchasePrice)/float64(currentQuantity+purchasedQuantity)
	}

	calculateProfit := func(sellPrice, weightedAverage float64, quantity int) float64 {
		return ((sellPrice - weightedAverage) * float64(quantity)) * 0.2
	}

	var weightedAverage float64
	var quantity int
	taxResults := make([]TaxResponse, 0)

	for _, op := range orders {
		if buy(op) {
			weightedAverage = calculateWeightedAverage(
				weightedAverage,
				quantity,
				op.UnitCost,
				op.Quantity,
			)
			quantity += op.Quantity
			taxResults = append(taxResults, TaxResponse{Tax: 0.0})
		} else if sell(op) { // sell
			tax := 0.0
			quantity -= op.Quantity
			if op.UnitCost*float64(op.Quantity) > taxThreshold {
				tax = calculateProfit(op.UnitCost, weightedAverage, op.Quantity)
			}
			taxResults = append(taxResults, TaxResponse{Tax: tax})
		}
	}

	return taxResults
}
