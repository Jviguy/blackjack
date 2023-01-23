package blackjack

import "strconv"
import "fmt"

type Deck struct {
	Cards []Card
}

func (d *Deck) AddCard(c Card) {
	d.Cards = append(d.Cards, c)
}

func (d Deck) Print() {
	for i, card := range d.Cards {
		if i == 0 { 
			fmt.Print(", ")
		}
		card.Print()
	}
}

func NewDeck() *Deck {
	deck := Deck{}
	for suit := 0; suit < 3; suit++ {
		deck.AddCard(Card{Suit: Suit(suit), Name: "A"})
		for i := 2; i < 10; i++ {
			deck.AddCard(Card{Suit: Suit(suit), Name: strconv.Itoa(i)})
		}
		deck.AddCard(Card{Suit: Suit(suit), Name: "J"})
		deck.AddCard(Card{Suit: Suit(suit), Name: "Q"})
		deck.AddCard(Card{Suit: Suit(suit), Name: "K"})
	}
	return &deck
}