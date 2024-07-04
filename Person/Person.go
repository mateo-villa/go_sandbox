package main

import "fmt"

type Person struct {
	name    string
	address string
	age     int
}

func (p *Person) getName() string {
	return p.name
}

func (p *Person) getAddress() string {
	return p.address
}

func (p *Person) getAge() int {
	return p.age
}

func (p Person) prettyPrint() {
	fmt.Printf("%s lives on %s and is %d years old", p.name, p.address, p.age)
}
