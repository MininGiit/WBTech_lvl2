package types

type CharacterBuilder interface {
	SetName(name string) CharacterBuilder
	SetClass(class string) CharacterBuilder
	SetStrength(strength int) CharacterBuilder
	SetAgility(agility int) CharacterBuilder
	SetIntelligence(intelligence int) CharacterBuilder
	Build() *Character
}
