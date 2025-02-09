package types

type Shape interface {
	Accept(Visitor)
}

type Circle struct {
	Radius float64
}

func NewCircle(radius float64) *Circle {
	return &Circle{Radius: radius}
}

func (c *Circle) Accept(v Visitor) {
	v.VisitCircle(c)
}

type Rectangle struct {
	Width, Height float64
}

func NewRectangle(width, height float64) *Rectangle {
	return &Rectangle{Width: width, Height: height}
}

func (r *Rectangle) Accept(v Visitor) {
	v.VisitRectangle(r)
}