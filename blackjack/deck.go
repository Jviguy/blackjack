package blackjack

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type Deck struct {
	Cards []Card
}

func (d *Deck) AppendCard(c Card) {
	d.Cards = append(d.Cards, c)
}

func (d *Deck) TakeTopCard() Card {
	card := d.Cards[len(d.Cards)-1]
	d.Cards = d.Cards[:len(d.Cards)-1]
	return card
}

func (d *Deck) Print() {
	for i, card := range d.Cards {
		if i != 0 {
			fmt.Print(", ")
		}
		card.Print()
	}
	fmt.Println()
}

func (d *Deck) Shuffle() {
	cards := make([]Card,0,52)
	for i := 0; i < 52; i++ {
		cards = append(cards, Card{})
	}
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	for _, card := range d.Cards {
		index := r.Intn(52)
		for cards[index].Name != "" {
			index = rand.Intn(52)
		}
		cards[index] = card
	}
	d.Cards = cards
}

func NewDeck() *Deck {
	deck := Deck{}
	for suit := 0; suit < 4; suit++ {
		deck.AppendCard(Card{Suit: Suit(suit), Name: "A"})
		for i := 2; i < 11; i++ {
			deck.AppendCard(Card{Suit: Suit(suit), Name: strconv.Itoa(i)})
		}
		deck.AppendCard(Card{Suit: Suit(suit), Name: "J"})
		deck.AppendCard(Card{Suit: Suit(suit), Name: "Q"})
		deck.AppendCard(Card{Suit: Suit(suit), Name: "K"})
	}
	return &deck
}
