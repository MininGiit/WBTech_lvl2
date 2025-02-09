package types

import("fmt")

type Notification struct {}

func (n *Notification) SendNotificatoin() {
	fmt.Println("notification sent")
}

func NewHotification() *Notification{
	return &Notification{}
}