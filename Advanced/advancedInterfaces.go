package main

import (
	"fmt"
	"math"
)

type rectangle struct {
	width, height float64
}

type circle struct {
	radius float64
}

type geometry interface {
	area() float64
	perimeter() float64
}

func (c circle) area() float64 {
	return c.radius * c.radius
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func (r rectangle) area() float64 {
	return r.width * r.height
}

func (r rectangle) perimeter() float64 {
	return 2 * (r.width + r.height)
}

func calculate(g geometry) {
	fmt.Printf("Area and Perimeter of type %T :\n", g)
	fmt.Println(g.area())
	fmt.Println(g.perimeter())
}

func testInterfaces() {
	c := circle{radius: 10}
	r := rectangle{width: 20, height: 30}
	calculate(c)
	calculate(r)
}
