package main

import (
	"fmt"
)

func main() {
	var Num int = 10 //Num is assign with some integer values
	var Ptr *int     // Declare a pointer as int data type
	Ptr = &Num       // Assign the address of Num to Ptr
	fmt.Println("Value Stored is :", Num)
	fmt.Println("The address of Num is :", Ptr)
	fmt.Println(*Ptr)
}
