package blackjack

type Hand struct {
	Cards []Card
}

func (h *Hand) IsSoft() bool {
	for _, card := range h.Cards {
		if card.Name == "A" {
			return true
		}
	}
	return false
}

func (h *Hand) Value() int {
	val := 0
	numa := 0
	for _, card := range h.Cards {
		v := card.Value()
		if v == 11 {
			numa++
		}
		val += v
	}
	// soften hand if need be
	for val > 21 {
		for i := 0; i < numa; i++ {
			val -= 10
		}
		if val < 22 {
			break
		}
	}
	return val
}

func (h *Hand) AddCard(c Card) {
	h.Cards = append(h.Cards, c)
}

func (h *Hand) Split(c Card) Hand {
	other := h.Cards[1]
	h.Cards = h.Cards[:1]
	return Hand{Cards: []Card{other}}
}
