package main

// defined the same way as a struct, but uses `interface` keyword
type Shape interface {
	area() float64 // set of methods instead of fields (must have to implement `Shape` interface)
}

// an interface can be used in function argument
func totalArea(shapes ...Shape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.area()
	}
	return area
}

// can also be a field of a struct
type MultiShape struct {
	shapes []Shape
}

// turning `MultiShape` struct into a `Shape` by giving it an area method
func (m *MultiShape) area() float64 {
	var area float64
	for _, s := range m.shapes {
		area += s.area()
	}
	return area
}
