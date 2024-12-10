package main

import (
	"bufio" // Buffered I/O operations enhance performance in more efficient ways to read or write data compared to raw I/O operations.
	"fmt"
	"os"
	"strings" // Provides utilities to manipulate strings, such as splitting and joining them.
)

func reverse_Sentence(sentence string) string {
	// words is a slice, where each word of input strings is its element
	//sentence := "Go is great"
	// words := strings.Fields(sentence)
	// words = []string{"Go", "is", "great"}  became a slice
	words := strings.Fields(sentence) // splits the input sentence into a slice of words based on whitespace.

	// Initialize i and j
	i := 0
	j := len(words) - 1

	// Loop to swap the characters
	for i < j { // this a also a type of for loop, available only in go, because go doesn't have while loop
		// Swap the words at i and j
		temp := words[i]
		words[i] = words[j]
		words[j] = temp

		// Move i forward and j backward
		i = i + 1
		j = j - 1
	}

	// Join the words back into a sentence
	return strings.Join(words, " ") // combines the words back into a single string, separated by spaces.
	//[]string{"great", "is", "Go"} → "great is Go".
}

func main() {
	// Create a Scanner to Read from Standard Input
	scanner := bufio.NewScanner(os.Stdin)
	// bufio.NewScanner: This function creates a new scanner. The work of scanner is to read & write the user input.
	//A scanner is an object that reads input line-by-line (or token-by-token, depending on its configuration).
	//os.Stdin: This refers to the standard input stream, which is where user input from the keyboard will come from.

	fmt.Println("Enter a sentence you want to reverse:")
	// Scan the entire line of input
	scanner.Scan() //This method reads a line of input from the user

	// Get the text entered by the user
	sentence := scanner.Text() // After calling scanner.Scan(), you use scanner.Text() to write/store the value in the variable.
	// scanner.scan() ---> read the user input
	// variable := scanner.Text() -----> write/store the value in the variable.
	// Reverse the sentence
	reversed := reverse_Sentence(sentence)

	// Print the reversed sentence
	fmt.Println("Original:", sentence)
	fmt.Println("Reversed:", reversed)
}
