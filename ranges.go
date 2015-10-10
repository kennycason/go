package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
    for x, p := range pow {
        fmt.Printf("%d = %d\n", x, p)
    }
}
