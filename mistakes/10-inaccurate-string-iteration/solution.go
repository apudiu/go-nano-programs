package main

import (
	"fmt"
)

func main() {
	s := "hêllo"
	fmt.Printf("The str is: %s \n", s)

	for i, v := range s {
		// NOTICE: here we taken whole rune, so it doesn't matter how many bytes it has
		fmt.Printf("position %d: %c\n", i, v)
	}
}

// The other approach is to convert the string into a slice of runes and iterate over it:
// like:
//s := "hêllo"
//runes := []rune(s)
// range over rules
