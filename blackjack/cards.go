package blackjack

import "github.com/fatih/color"



type Suit uint8

func (s Suit) Color() color.Attribute {
	if s < 2 {
		return color.FgBlack
	} else {
		return color.FgRed
	}
}

func (s Suit) Char() string {
	switch s {
	case Clubs:
		return "♣"
	case Spades:
		return "♠"
	case Diamonds:
		return "♦"
	case Hearts:
		return "♥"
	default:
		return "ERR"
	}
}

const (
	Spades Suit = iota
	Clubs
	Diamonds
	Hearts
)

// Store cards as their values
type Card struct {
	Name string
	Suit Suit
}

func (c Card) Print() {
	x := color.New(c.Suit.Color())
	x.Print(c.Name+c.Suit.Char())
}