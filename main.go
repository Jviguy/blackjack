package main

import (
	"github.com/jviguy/blackjack/blackjack"
)

func main() {
	table := blackjack.Table{Rules: blackjack.SS17,Shoe: blackjack.NewShoe(6)}
	table.StartRound()
}