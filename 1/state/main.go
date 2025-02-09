package main

import (
	"fmt"
	"state/types"
)

func main() {
	light := types.NewLight()
   
	fmt.Println("Initial state:", light.GetState())
	light.TurnOn()                
	fmt.Println("Current state:", light.GetState()) 
   
	light.Brighten()
	fmt.Println("Current state:", light.GetState())
   
	light.TurnOff()
	fmt.Println("Current state:", light.GetState())
   
	light.TurnOff()
	fmt.Println("Current state:", light.GetState())
   }
