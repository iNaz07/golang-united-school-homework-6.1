package golang_united_school_homework

import "math"

// Circle must satisfy to Shape interface
type Circle struct {
	Radius float64
}

func (c *Circle) CalcPerimeter() float64 {
	return math.Pi * 2 * c.Radius
}

// CalcArea returns calculation result of area
func (c *Circle) CalcArea() float64 {
	return math.Pi * c.Radius * c.Radius
}
