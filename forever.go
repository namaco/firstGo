package main

import (
    "time"
    "fmt"
)

func main() {
    for {
        time.Sleep(1000)
        fmt.Println("Hello")
    }
}
