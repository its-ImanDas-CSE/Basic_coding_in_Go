package main

import (
	"fmt"
)

func Reverse_String(S string) string {
	//The string S is converted into a slice of rune. This allows us to handle characters that may require multiple bytes (like emojis or accented characters).
	reverse := []rune(S) // we convert the string into a slice, rune type.
	// Initialize i and j
	i := 0
	j := len(reverse) - 1

	// Loop to swap the characters
	for i < j {
		// Swap the characters at i and j
		temp := reverse[i]
		reverse[i] = reverse[j]
		reverse[j] = temp

		// Move i forward and j backward
		i = i + 1
		j = j - 1
	}

	return string(reverse)
}
func rev() {
	fmt.Println("Give the String you want to reverse")
	var S string
	fmt.Scanln(&S)
	result := Reverse_String(S)
	fmt.Println(result)
}
