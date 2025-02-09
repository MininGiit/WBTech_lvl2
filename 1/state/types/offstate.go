package types

import "fmt"

type OffState struct{}

func (o *OffState) TurnOn(light *Light) {
 	fmt.Println("Turning light ON")
 	light.SetState(&OnState{})
}

func (o *OffState) TurnOff(light *Light) {
 	fmt.Println("Light is already OFF")
}

func (o *OffState) Brighten(light *Light) {
 	fmt.Println("Cannot brighten light when it's OFF")
}

func (o *OffState) String() string {
 	return "Off"
}