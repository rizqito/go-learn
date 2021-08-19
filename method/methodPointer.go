package main

import "fmt"

type person struct {
	name string
	age  int
}

func (p person) changeName1(name string) {
	fmt.Println("---> on changeName1, name changed to", name)
	p.name = name
}

func (p *person) changeName2(name string) {
	fmt.Println("---> on changeName2, name changed to", name)
	p.name = name
}

func methodPointer() {
	var p1 = person{"alexa", 21}
	fmt.Println("s1 before", p1.name)

	p1.changeName1("jonathan")
	fmt.Println("s1 changed name 1", p1.name)

	p1.changeName2("carl")
	fmt.Println("s1 changed name 2", p1.name)
}
