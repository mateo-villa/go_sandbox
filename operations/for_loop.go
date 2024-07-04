package main

import "fmt"

func main() {

	name := "skipper"
	for x := 0; x < len(name); x++ {
		fmt.Printf("%c\n", name[x])
	}
	for _, char := range name {
		fmt.Printf("%c\n", char)
	}
}
