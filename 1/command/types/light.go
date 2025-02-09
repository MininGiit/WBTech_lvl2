package types

import "fmt"

type Light struct{
	state bool
}

func (l *Light) On() {
	l.state = true
	fmt.Println("Light is On")
}

func (l *Light) Off() {
	l.state = false
	fmt.Println("Light is Off")
}
