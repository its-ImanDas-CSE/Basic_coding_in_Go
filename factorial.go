package main

import (
	"fmt"
)

func Factorial(Num int) (Fact int) {
	Fact = 1
	for i := Num; i >= 1; i-- {
		Fact = Fact * i
	}
	return Fact
}
func main() {
	var Num int
	fmt.Println("Give the number")
	fmt.Scanln(&Num)

	fmt.Println(Factorial(Num))
}
