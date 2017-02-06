package main

import (
    "fmt"
    "math/cmplx"
)

//var c, python, java bool
var i, j int = 1, 2

var (
    ToBe    bool        = false
    MaxInt  uint64      = 1<<64 - 1
    z       complex128  = cmplx.Sqrt(-5 + 12i)
)

func add(x, y int) int {
    return x + y
}

func swap(x, y string) (string, string) {
    return y, x
}

func split(sum int) (x, y int) {
    x = sum * 4 / 9
    y = sum - x
    return
}

func main() {
    a, b := swap("hello", "world")
    fmt.Println(add(42, 13))
    fmt.Println(a, b)
    fmt.Println(split(17))

    k := 3
    var c, python, java = true, false, "no!"
    fmt.Println(i, j, k, c, python, java)

    const f = "%T(%v)\n"
    fmt.Printf(f, ToBe, ToBe)
    fmt.Printf(f, MaxInt, MaxInt)
    fmt.Printf(f, z, z)
}