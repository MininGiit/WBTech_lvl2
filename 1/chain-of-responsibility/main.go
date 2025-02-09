package main

import(
	"chain/types"
	"fmt"
)

func main() {
	handlerA := &types.ConcreteHandlerA{}
	handlerB := &types.ConcreteHandlerB{}

	handlerA.SetNext(handlerB)

	fmt.Println("Sending request 'A':")
	handlerA.Handle("A")

	fmt.Println("\nSending request 'B':")
	handlerA.Handle("B")

	fmt.Println("\nSending request 'C':")
	handlerA.Handle("C")
}