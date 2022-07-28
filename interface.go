package main

import (
	"fmt"
	"math"
)

//There is no relationship between interface and implementations
type shape interface {
	area() float64
	perimeter() float64
}

//Rect
type recta struct {
	width, height float64
}

//Circle
type circle struct {
	radius float64
}

//implemetations for Shape interface
func (r recta) area() float64 { return r.width * r.height }
func (r recta) perimeter() float64 {
	return 2 * (r.width + r.height)
}

//implemetations for Shape interface
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

// receive all type as empty interface then return string,error
func toString(v interface{}) (string, error) {
	return "test", nil
}

func Interfacetest() {
	fmt.Println("Interface Test")
	shapes := []shape{recta{5, 10}, circle{5}}
	for _, s := range shapes {
		a := s.area()
		p := s.perimeter()
		fmt.Println("Area : ", a)
		fmt.Println("Perimeter :", p)
	}

	var a interface{} = 1 // similar with Object of Java or void* of C++

	i := a       // dynamic type, value is 1, NOT int
	j := a.(int) // means a is not null and a is int type and it's value is 1

	fmt.Println(toString(j))
	fmt.Println(i) // display pointer
	fmt.Println(j) // display value
}
