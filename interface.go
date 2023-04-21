package main

import "fmt"

type shape interface {
	area() float32
}
type rectangle struct {
	length, breadth float32
}

func (r rectangle) area() float32 {
	return r.length * r.breadth
}
func main() {
	var s shape
	s = rectangle{10, 14}
	fmt.Println("area of the rectangle is:", s.area())
}
