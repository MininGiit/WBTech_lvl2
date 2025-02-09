package types

import "fmt"

type OnState struct{}

func (o *OnState) TurnOn(light *Light) {
	fmt.Println("Light is already ON")
}

func (o *OnState) TurnOff(light *Light) {
 	fmt.Println("Turning light OFF")
 	light.SetState(&OffState{})
}

func (o *OnState) Brighten(light *Light) {
 	fmt.Println("Light is already at maximum brightness")
}

func (o *OnState) String() string {
 	return "On"
}
