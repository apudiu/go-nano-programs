package main

import "fmt"

func main() {
	s := "hêllo"
	fmt.Printf("The str is: %s \n", s)

	for i := range s {
		// NOTICE: here (s[i]) is 1st byte (of multi byte) rule, so it leaves all other after 1st
		fmt.Printf("position %d: %c\n", i, s[i])
	}
	fmt.Printf("[ UNEXPECTED ] len=%d\n", len(s))
}
