package main

import "fmt"

// 返回一个“返回int的函数”
// The returned function has attached global variables a and b
func fibonacci() func() int {
    a, b := -1, 1
    return func() int {
        a, b = b, a+b
        return b
    }
}

func main() {
    f := fibonacci()
    for i := 0; i < 10; i++ {
        fmt.Println(f())
    }
}