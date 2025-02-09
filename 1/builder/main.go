package main

import (
	"builder/types"
	"fmt"
)


func main() {
	warriorBuilder := types.NewWarriorBuilder()
	creator := types.NewCharacterCreator(warriorBuilder)
	creator.ConstructWarrior()
	warrior := creator.GetCharacter()
	fmt.Println(warrior)
}