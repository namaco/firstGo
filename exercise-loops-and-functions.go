package main

import (
    "fmt"
    "math"
)

func Sqrt(x float64) float64 {
    z := .1
    for i := 0; i < 100; i++{
        t := z - ( z * z - x ) / ( z * 2)
        if d := math.Abs(z - t); d < 1e-6 {
            return t
            break
        }
        fmt.Println(">>>>>t is:", t)
        z = t
    }
    return z
}

func main() {
    fmt.Println(Sqrt(2))
    fmt.Println(math.Sqrt(2))
}