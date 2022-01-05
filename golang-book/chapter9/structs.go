package main

import (
	"fmt"
	"math"
)

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

// struct, a type that contain set of fields
type Circle struct {
	x, y, r float64
}

type Rectangle struct {
	x1, y1, x2, y2 float64
}

// struct represents "has a" relationship
// Person has a name
type Person struct {
	Name string
}

type Android struct {
	// Person Person => if we typed like this, it will sounds like "Android has a Person"
	Person // use this instead, so it sounds "Android belongs to a Person" (embedded type/anonymous field/only type, no name)
	Model  string
}

// `func` keyword => "receiver" (method (special type of function)) => functionName(parameter) => return type
// `circleInstanceVariable.area()`
func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

// even though the function name is the same (area)
// but since the "receiver" is different, it will not thrown an error
func (r *Rectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return l * w
}

func (p *Person) Talk() {
	fmt.Println("Hi, my name is", p.Name)
}

func main() {
	// zero in int => 0
	// zero in float => 0.0
	// zero in string => ""
	// zero in pointer => nil
	var circle1 Circle
	circle1 = Circle{0, 0, 5}
	circle2 := Circle{x: 0, y: 0, r: 5} // with fields name
	circle3 := new(Circle)              // returns a pointer, allocate memory required for all fields
	fmt.Println(circle1.x)              // accessing a field
	circle2.r = 6                       // modifying a field
	fmt.Println(circle2)
	fmt.Println(circle3)
	fmt.Println(circle1.area()) // no need to use `&` keyword, Go knows it and will do it for us

	rectangle := Rectangle{0, 0, 10, 10}
	fmt.Println(rectangle.area())

	person := Person{"John Doe"}
	person.Talk()

	android := new(Android)
	android.Person.Name = "Jane Doe"
	android.Person.Talk() // one way (using the `Person` struct explicitly)
	android.Talk()        // another way (directly)
}
