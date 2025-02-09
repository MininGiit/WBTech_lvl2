package main

import "facade/types"

func main(){
	card := types.NewCard()
	notifications := types.NewHotification()
	receipt := types.NewReceipt()
	paymentSystem := types.NewPaymentSystem(card, notifications, receipt)
	paymentSystem.Pay()
}