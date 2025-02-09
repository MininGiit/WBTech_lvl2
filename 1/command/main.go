package main

import "command/types"

func main() {
	light := &types.Light{}
	lightOn := &types.LightOnCommand{Light: light}
	lightOff := &types.LightOffCommand{Light: light}
	remote := &types.RemoteControl{}
	remote.SetCommand(lightOn)
	remote.PressButton()
	remote.SetCommand(lightOff)
	remote.PressButton()
}