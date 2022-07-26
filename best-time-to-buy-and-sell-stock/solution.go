package best_time_to_buy_and_sell_stock

import "math"

func maxProfit(prices []int) int {
	if !areEntryPricesValid(prices) {
		return 0
	}

	var (
		min    = prices[0]
		profit int
	)

	for _, price := range prices {
		if price < min {
			min = price
			continue
		}

		if price-min > profit {
			profit = price - min
		}
	}

	return profit
}

func areEntryPricesValid(prices []int) bool {
	if len(prices) < 1 || float64(len(prices)) > math.Pow10(5) {
		return false
	}

	for _, price := range prices {
		if price < 0 || float64(price) > math.Pow10(4) {
			return false
		}
	}

	return true
}
