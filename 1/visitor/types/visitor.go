package types

import "math"

type Visitor interface {
	VisitCircle(*Circle)
	VisitRectangle(*Rectangle)
}

type AreaVisitor struct {
	Area float64
}

func (a *AreaVisitor) VisitCircle(c *Circle) {
	a.Area = math.Pi * c.Radius * c.Radius
}

func (a *AreaVisitor) VisitRectangle(r *Rectangle) {
	a.Area = r.Width * r.Height
}

type PerimeterVisitor struct {
	Perimeter float64
}

func (p *PerimeterVisitor) VisitCircle(c *Circle) {
	p.Perimeter = 2 * math.Pi * c.Radius
}

func (p *PerimeterVisitor) VisitRectangle(r *Rectangle) {
	p.Perimeter = 2 * (r.Width + r.Height)
}