package main

import "fmt"

type Shape interface{
	Area() int
	Peri() int
}

type rect struct{
	l, b int
}

type sqr struct{
	side int
}
func (r rect)Area() int{
	return r.l*r.b
}

func (s sqr)Area() int{
	return s.side*s.side
}
func (r rect)Peri() int{
	return 2*(r.l+r.b)
}
func (s sqr)Peri() int{
	return 2*(s.side+s.side)
}

func main(){
	r := rect{10,20}
	s:= sqr{100}
	var shape Shape
	shape = r
	fmt.Println("Area of Rectange: ",shape.Area())
	fmt.Println("Peri of Rectangle: ",shape.Peri())
	shape = s
	fmt.Println("Area of sqr: ",shape.Area())
	fmt.Println("Peri of sqr: ",shape.Peri())

}
