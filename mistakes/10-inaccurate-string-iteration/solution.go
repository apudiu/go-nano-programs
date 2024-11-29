package main

import "fmt"

func main() {
    s := "hêllo"
    fmt.Printf("The str is: %s \n", s)

    for i, v := range s {
        // NOTICE: here we taken whole rune, so it doesn't matter how many bytes it has
        fmt.Printf("position %d: %c\n", i, v)
    }
}
