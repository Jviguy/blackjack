package main

import (
	"github.com/jviguy/blackjack/blackjack"
)

func main() {
	shoe := blackjack.NewShoe(6)
	shoe.Print()
}