package types

import "fmt"

type Square struct {
    side float64
}

func (s *Square) Draw() {
    fmt.Println("Drawing a square with side:", s.side)
}

type SquareFactory struct {
    side float64
}

func NewSquareFactory(side float64) *SquareFactory {
	return &SquareFactory{side: side}
} 

func (s *SquareFactory) CreateShape() Shape {
    return &Square{side: s.side}
}