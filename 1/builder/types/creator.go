package types

type CharacterCreator struct {
	builder CharacterBuilder
}

func NewCharacterCreator(builder CharacterBuilder) *CharacterCreator {
	return &CharacterCreator{builder: builder}
}

func (c *CharacterCreator) ConstructWarrior() {
	c.builder.SetName("Slavik").
		SetClass("Developer").
		SetStrength(100).
		SetAgility(15).
		SetIntelligence(100)
}

func (c *CharacterCreator) ConstructМagician() {
	c.builder.SetName("Vitya").
		SetClass("DevOps").
		SetStrength(1000).
		SetAgility(1500).
		SetIntelligence(1000)
}

func (c *CharacterCreator) GetCharacter() *Character {
	return c.builder.Build()
}