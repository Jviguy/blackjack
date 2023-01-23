package blackjack

type Shoe struct {
	ActiveDecks []*Deck
	DeckIndex   int
}

func (s *Shoe) Shuffle() {
	for i := range s.ActiveDecks {
		deck := NewDeck()
		deck.Shuffle()
		s.ActiveDecks[i] = deck
	}
}

func (s *Shoe) NextDeck() {
	s.DeckIndex++
	if s.DeckIndex > len(s.ActiveDecks) {
		s.Shuffle()
	}
}

func (s *Shoe) TakeCard() Card {
	deck := s.ActiveDecks[s.DeckIndex]
	if len(deck.Cards) < 1 {
		s.NextDeck()
		return s.TakeCard()
	}
	return deck.TakeTopCard()
}

func (s *Shoe) Print() {
	for _, deck := range s.ActiveDecks {
		deck.Print()
	}
}

func NewShoe(size int) *Shoe {
	s := Shoe{DeckIndex: 0,ActiveDecks: make([]*Deck, 0 , size)}
	for i := 0; i < size; i++ {
		s.ActiveDecks = append(s.ActiveDecks, NewDeck())
		s.ActiveDecks[i].Shuffle()
	}
	return &s
}