package blackjack

type TableRule uint8

const (
	SS17 TableRule = iota << 1
	HH17
	BJ32
)

type Table struct {
	Shoe Shoe
	Rules TableRule
	
}