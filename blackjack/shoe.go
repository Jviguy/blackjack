package blackjack

import (
	"fmt"
	"math/rand"
	"time"
)

type Shoe struct {
	Cards []Card
	Size int
	Count int
}

func (s *Shoe) Shuffle() {
	s.Count = 0
	cards := make([]Card,0,52*s.Size)
	for i := 0; i < 52*s.Size; i++ {
		cards = append(cards, Card{})
	}
	x := rand.NewSource(time.Now().UnixNano())
	r := rand.New(x)
	for _, card := range s.Cards {
		index := r.Intn(52*s.Size)
		for cards[index].Name != "" {
			index = rand.Intn(52*s.Size)
		}
		cards[index] = card
	}
	s.Cards = cards
}

func (s *Shoe) TakeCard() Card {
	if len(s.Cards) == 0 {
		s.Shuffle()
	}
	c := s.Cards[len(s.Cards)-1]
	s.Cards = s.Cards[:len(s.Cards)-1]
	if c.Value() <= 7 {
		s.Count += 1
	} else if c.Value() > 9 {
		s.Count -= 1
	}
	return c
}

func (s *Shoe) Print() {
	for i, c := range s.Cards {
		if i != 0 {
			fmt.Print(", ")
		}
		c.Print()
	}
	fmt.Println()
}

func NewShoe(size int) *Shoe {
	s := Shoe{Size: size}
	for i := 0; i < size; i++ {
		s.Cards = append(s.Cards, NewDeck().Cards...)
	}
	s.Shuffle()
	return &s
}