package main

import "fmt"

type People interface {
	Name()
	Age()
}

type G struct {
	age  int
	name string
}

func (receiver *G) Name() {
	fmt.Println(receiver.name)
}
func (receiver *G) Age() {
	fmt.Println(receiver.age)
}

func main() {
	var people People
	people = &G{
		age:  29,
		name: "Gm",
	}
	people.Name()
	people.Age()
}
