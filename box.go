package golang_united_school_homework

import "fmt"

// box contains list of shapes and able to perform operations on them
type box struct {
	shapes         []Shape
	shapesCapacity int // Maximum quantity of shapes that can be inside the box.
}

// NewBox creates new instance of box
func NewBox(shapesCapacity int) *box {
	return &box{
		shapesCapacity: shapesCapacity,
	}
}

// AddShape adds shape to the box
// returns the error in case it goes out of the shapesCapacity range.
func (b *box) AddShape(shape Shape) error {
	if len(b.shapes) == b.shapesCapacity {
		return fmt.Errorf("box is full")
	}
	b.shapes = append(b.shapes, shape)
	return nil
}

// GetByIndex allows getting shape by index
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) GetByIndex(i int) (Shape, error) {
	if i >= b.shapesCapacity || b.shapes[i] == nil {
		return nil, fmt.Errorf("shape doesnt exist or index went out of range")
	}
	return b.shapes[i], nil
}

// ExtractByIndex allows getting shape by index and removes this shape from the list.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ExtractByIndex(i int) (Shape, error) {

	sh, err := b.GetByIndex(i)
	if err != nil {
		return nil, fmt.Errorf("shape doesnt exist or index went out of range")
	}
	end := b.shapes[i+1:]
	b.shapes = append(b.shapes[:i], end...)
	return sh, nil
}

// ReplaceByIndex allows replacing shape by index and returns removed shape.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ReplaceByIndex(i int, shape Shape) (Shape, error) {
	sh, err := b.GetByIndex(i)
	if err != nil {
		return nil, fmt.Errorf("shape doesnt exist or index went out of range")
	}
	b.shapes[i] = shape
	return sh, nil

}

// SumPerimeter provides sum perimeter of all shapes in the list.
func (b *box) SumPerimeter() float64 {
	var res float64
	for _, shape := range b.shapes {
		res += shape.CalcPerimeter()
	}
	return res
}

// SumArea provides sum area of all shapes in the list.
func (b *box) SumArea() float64 {

	var res float64
	for _, shape := range b.shapes {
		res += shape.CalcArea()
	}
	return res

}

// RemoveAllCircles removes all circles in the list
// whether circles are not exist in the list, then returns an error
func (b *box) RemoveAllCircles() error {
	var res []Shape
	var count int
	for _, circle := range b.shapes {
		if _, ok := circle.(*Circle); !ok {
			res = append(res, circle)
		} else {
			count++
		}
	}
	b.shapes = res
	if count == 0 {
		return fmt.Errorf("no any circles in the box")
	}
	return nil
}
