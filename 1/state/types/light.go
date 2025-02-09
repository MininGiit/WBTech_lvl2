package types

type Light struct {
	state LightState
}
   
func NewLight() *Light {
	return &Light{state: &OffState{}}
}
   
func (l *Light) SetState(state LightState) {
	l.state = state
}
   
func (l *Light) TurnOn() {
	l.state.TurnOn(l)
}
   
func (l *Light) TurnOff() {
	l.state.TurnOff(l)
}

func (l *Light) Brighten() {
	l.state.Brighten(l)
}
   
func (l *Light) GetState() string {
	return l.state.String()
}