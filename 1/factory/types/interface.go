package types

type Shape interface {
    Draw()
}

type ShapeFactory interface {
    CreateShape() Shape
}
