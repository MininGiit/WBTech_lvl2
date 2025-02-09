package types

type Command interface {
	Execute()
}

type LightOnCommand struct {
	Light *Light
}

func (c *LightOnCommand) Execute() {
	c.Light.On()
}

type LightOffCommand struct {
	Light *Light
}

func (c *LightOffCommand) Execute() {
	c.Light.Off()
}

