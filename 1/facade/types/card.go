package types

import("fmt")

type Card struct {}

func (c *Card) Pay() {
	fmt.Println("payment completed")
}

func NewCard() *Card{
	return &Card{}
}