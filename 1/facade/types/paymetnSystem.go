package types

type PaymentSystem struct {
	Card *Card
	Notification *Notification
	Receipt *Receipt
}

func NewPaymentSystem(card *Card, not *Notification, rec *Receipt) *PaymentSystem{
	return &PaymentSystem{
		Card: card,
		Notification: not,
		Receipt: rec,
	}
}

func (p *PaymentSystem) Pay() {
	p.Card.Pay()
	p.Receipt.Generate()
	p.Receipt.SendReceipt()
	p.Notification.SendNotificatoin()
}