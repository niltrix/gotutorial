package main

import "fmt"

type rect struct {
	width, height int
}

func newRect() *rect {
	r := rect{}
	r.height = 0
	r.width = 0
	return &r
}

func (rp *rect) area() int {
	return rp.width * rp.height
}

func (r rect) perimeter() int {
	return 2 * (r.width + r.height)
}

func main() {
	r2 := newRect()
	fmt.Printf("Rect %p %v \n", &r2, r2)
	fmt.Println("Area: ", r2.area())
	fmt.Println("perimeter: ", r2.perimeter())

	r := rect{width: 10, height: 5}

	fmt.Printf("Rect %p %v \n", &r, r)
	fmt.Println("Area: ", r.area())
	fmt.Println("perimeter: ", r.perimeter())

	rp := &r
	fmt.Printf("Rect %p %v \n", &rp, rp)
	fmt.Println("Area: ", rp.area())
	fmt.Println("perimeter: ", rp.perimeter())
}
