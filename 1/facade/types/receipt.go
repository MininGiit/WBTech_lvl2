package types

import "fmt"

type Receipt struct{}

func (r *Receipt) Generate() {
	fmt.Println("receipt created")
}

func (r *Receipt) SendReceipt() {
	fmt.Println("receipt sent")
}

func NewReceipt() *Receipt {
	return &Receipt{}
}