package main

func main() {
	var peter Person = Person{
		name:    "Peter",
		address: "1 oak street",
		age:     42,
	}
	peter.prettyPrint()
}
