package main

import "factory/types"

func drawShape(factory types.ShapeFactory) {
    shape := factory.CreateShape()
    shape.Draw()
}

func main() {
    circleFactory := types.NewCircleFactory(10)
    squareFactory := types.NewSquareFactory(15)

    drawShape(circleFactory)
    drawShape(squareFactory)
}