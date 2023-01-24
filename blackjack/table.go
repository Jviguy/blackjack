package blackjack

import (
	"bufio"
	"fmt"
	"strings"
)

type TableRule uint8

const (
	SS17 TableRule = iota << 1
	HH17
	BJ32
)

type Table struct {
	Shoe *Shoe
	Rules TableRule
	Hand [][]Hand
	TurnIdx int
	Scanner *bufio.Scanner
}

func (t *Table) StartRound() {
	if len(t.Hand) == 0 {
		t.Hand = append(t.Hand, []Hand{})
		t.Hand = append(t.Hand, []Hand{})
	}
	t.Hand[0] = []Hand{{}}
	t.Hand[1] = []Hand{{}}
	t.Hand[0][0].AddCard(t.Shoe.TakeCard())
	t.Hand[0][0].AddCard(t.Shoe.TakeCard())
	fmt.Print("Dealer Hand: ")
	t.Hand[0][0].Cards[0].Print()
	fmt.Println()
	
	t.Hand[1][0].AddCard(t.Shoe.TakeCard())
	t.Hand[1][0].AddCard(t.Shoe.TakeCard())

	fmt.Print("Your Hand: ")
	t.Hand[1][0].Print()
	fmt.Println()
	for _,h := range t.Hand[1] {
		for h.Status == Active {
			actions := []string{"H", "S"}
			if len(h.Cards) == 2 {
				actions = append(actions, "D")
				if h.Cards[0] == h.Cards[1] {
					actions = append(actions, "SP")
				}
			}
			fmt.Print(strings.Join(actions, ", ")+": ")
		}
	}
	t.TurnIdx = 0
}

