package main

import (
	"fmt"
	"time"
)

// person type
type person struct {
	name string
	age int
}

// golang method
func (p person) sayHi() {
	fmt.Printf("Hola! I'm %s and I'm %d years old.\n", p.name, p.age)
	fmt.Printf("and agaian: Hola! My name is %v and I'm %v years old.\n", p.name, p.age)
}

func main() {
	maksimko := person{
		age: time.Now().Year() - 1983,
		name: "Maksimko",
	}
	maksimko.sayHi()
}
