package main

import (
	"fmt"
)

type speaker interface { // we make a interface, and declare a method "Speak" in it.
	Speak() string
}
type person struct { // we make a struct
	Name string
}

func (P person) Speak() string {
	return "Hello, my name is" + P.Name
}
func main() {
	Person := person{Name: "John"}
	var Speaker speaker
	Speaker = Person
	fmt.Println(Speaker.Speak())

}
