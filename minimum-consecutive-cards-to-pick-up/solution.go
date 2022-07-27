package minimum_consecutive_cards_to_pick_up

import "math"

func minimumCardPickup(cards []int) int {
	if !areEntryNumbersValid(cards) {
		return -1
	}

	var (
		cardLog = make(map[int]int)
		min     = len(cards) * 2
	)

	for cardIndex, card := range cards {
		if prevCardIndex, ok := cardLog[card]; ok {
			if min > (cardIndex-prevCardIndex)+1 {
				min = (cardIndex - prevCardIndex) + 1
			}

			cardLog[card] = cardIndex
			continue
		}

		cardLog[card] = cardIndex
	}

	if min == len(cards)*2 {
		return -1
	}

	return min
}

func areEntryNumbersValid(cards []int) bool {
	if len(cards) < 1 || float64(len(cards)) > math.Pow10(5) {
		return false
	}

	for _, card := range cards {
		if card < 0 || float64(card) > math.Pow10(6) {
			return false
		}
	}

	return true
}
