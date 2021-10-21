package example

import (
	"fmt"
	"reflect"
)

type Vertex struct {
	X int
	Y int
}

func Pointerttest() {
	v := Vertex{1, 2}
	p := &v
	r := &p
	s := &r
	fmt.Println(reflect.TypeOf(p).String())
	fmt.Println(&p)
	fmt.Println(reflect.TypeOf(r).String())
	fmt.Println(&r)
	fmt.Println(reflect.TypeOf(s).String())
	fmt.Println(&s)
	p.X = 1
	fmt.Println(v)
	(*r).X = 2
	fmt.Println(v)
	(*(*s)).X = 3
	fmt.Println(v)
}
