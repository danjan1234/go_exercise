package main

import "fmt"

func say(word string, flag chan int) {
    fmt.Println(word)
    flag <- 1

}
func main() {
    flag := make(chan int)
    go say("Subroutine", flag)
    fmt.Println("Main routine")
    <-flag  // Use flag to prevent main routine from terminating too early
}