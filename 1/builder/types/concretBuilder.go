package types

type WarriorBuilder struct {
	character *Character
}

func NewWarriorBuilder() *WarriorBuilder {
	return &WarriorBuilder{character: &Character{}}
}

func (b *WarriorBuilder) SetName(name string) CharacterBuilder {
	b.character.Name = name
	return b
}

func (b *WarriorBuilder) SetClass(class string) CharacterBuilder {
	b.character.Class = class
   return b
}

func (b *WarriorBuilder) SetStrength(strength int) CharacterBuilder {
   b.character.Strength = strength
   return b
}

func (b *WarriorBuilder) SetAgility(agility int) CharacterBuilder {
   b.character.Agility = agility
   return b
}

func (b *WarriorBuilder) SetIntelligence(intelligence int) CharacterBuilder {
   b.character.Intelligence = intelligence
   return b
}

func (b *WarriorBuilder) Build() *Character {
   return b.character
}