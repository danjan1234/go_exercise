package main

import (
    "fmt"
    "math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) String() string {
    return fmt.Sprintf("Haha")
}

func (e ErrNegativeSqrt) Error() string {
    if e < 0 {
        // Infinite loop if one does fmt.Sprintf(e)
        return fmt.Sprintf("Cannot Sqrt negative number: %v", float64(e))
    }
    return ""
}

func Sqrt(x float64) (float64, error) {
    if x >= 0 {
        return math.Sqrt(x), nil
    }
    return 0, ErrNegativeSqrt(x)
}

func main() {
    fmt.Println(Sqrt(2))
    fmt.Println(Sqrt(-2))
}
