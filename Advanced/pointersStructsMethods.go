package main

import "fmt"

type person struct {
	name string
	age  int
}

func CreateStructPointer(name string) *person {
	p := person{name: name}
	p.age = 30
	return &p
}

func main() {
	fmt.Println(person{"Sumit", 28})
	fmt.Println(person{"Sam", 29})
	fmt.Println(person{name: "Sid"})
	fmt.Println(CreateStructPointer("Nicolas"))
	fmt.Println(&person{"Nihal", 30})
	testP := person{name: "Tester"}
	fmt.Println(testP.name)
	fmt.Println(testP.age)
	testP.age = 20
	fmt.Println(testP.age)
	testPP := &testP
	testPP.name = "Tester1"
	testPP.age = 21
	fmt.Println(testPP.name)
	fmt.Println(testPP.age)
	(*testPP).name = "Tester2"
	(*testPP).age = 22
	fmt.Println(testPP.name)
	fmt.Println(testPP.age)
}
