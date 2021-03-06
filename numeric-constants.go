package main

import "fmt"

const (
    Big = 1 << 100
    Small = Big >> 99
    Test = 4 >> 1
)

func needInt(x int) int {
    return x * 10 + 1
}

func main() {
    fmt.Println(needInt(Small))
    fmt.Println(needFloat(Small))
    fmt.Println(needFloat(Big))
    fmt.Println(Test)
}
func needFloat(x float64) float64 {
    return x * 0.1
}
