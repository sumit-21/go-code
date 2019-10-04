package main

import "fmt"

type rect struct {
	height int
	width  int
}

func (r rect) area() int {
	return r.height * r.width
}

func (r *rect) perimeter() int {
	return 2 * (r.height + r.width)
}

func testMethods() {
	r := rect{height: 20, width: 30}
	fmt.Println(r.area())
	fmt.Println(r.perimeter())
	fmt.Println((&r).perimeter())
}
