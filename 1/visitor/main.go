package main

import (
	"fmt"
	"visitor/types"
)

func main() {
	circle := types.NewCircle(5)
	rectangle := types.NewRectangle(4, 3)

	areaVisitor := &types.AreaVisitor{}
	perimeterVisitor := &types.PerimeterVisitor{}


	circle.Accept(areaVisitor)
	fmt.Printf("Circle Area: %.2f\n", areaVisitor.Area)
	circle.Accept(perimeterVisitor)
	fmt.Printf("Circle Perimeter: %.2f\n", perimeterVisitor.Perimeter)

	rectangle.Accept(areaVisitor)
	fmt.Printf("Rectangle Area: %.2f\n", areaVisitor.Area)
	rectangle.Accept(perimeterVisitor)
	fmt.Printf("Rectangle Perimeter: %.2f\n", perimeterVisitor.Perimeter)
}