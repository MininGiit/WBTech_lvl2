package types

type LightState interface {
	TurnOn(light *Light)
	TurnOff(light *Light)
	Brighten(light *Light)
	String() string 
}