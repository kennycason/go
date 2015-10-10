package main

import (
    "fmt"
)

const start float64 = 1

func Sqrt(x float64) float64 {
    z := start
    for i := 0; i < 10; i++ {
        z = z - (z * z - x) / (2 * z)
    }
    return z
}

func main() {
    fmt.Println(Sqrt(2))
}
