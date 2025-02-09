package types

import "fmt"

type Circle struct {
    radius float64
}

func (c *Circle) Draw() {
    fmt.Println("Drawing a circle with radius:", c.radius)
}


type CircleFactory struct {
    radius float64
}

func NewCircleFactory(radius float64) *CircleFactory {
	return &CircleFactory{radius: radius}
} 

func (c *CircleFactory) CreateShape() Shape {
    return &Circle{radius: c.radius}
}
