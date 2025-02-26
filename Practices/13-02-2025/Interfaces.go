package main

import (
	"fmt"
	"math"
)


type shape interface {
	area() float64
	perimeter() float64
}

type rect struct{
	width, height float64
}

type circle struct{
	radius float64
}
func (r rect) area() float64 {
	return r.width * r.height
}

func (r circle) area() float64 {
	return math.Pi * r.radius * r.radius
}
func main () {
	a := circle{
		radius: 2,
	}

	b := rect{
		width: 8,
		height: 12,
	}
	fmt.Println(a.area())
	fmt.Println(b.area())
}