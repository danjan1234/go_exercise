package main

import "fmt"

func main() {
    ch := make(chan int)
    select {
    case ch <- 1:   // Blocked right away unless line 8 is changed to ch := make(chan int, 1)
        fmt.Println("Stored variable in the channel is: ", <-ch)
    default:
        fmt.Println("Default")
    }
}