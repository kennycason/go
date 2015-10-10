package main

import "fmt"

func swap(x, y string) (string, string) {
    return y, x
}

// namd return
func split(sum int) (int, int) {
    x := sum * 2
    y := sum - x
    return x, y
}

func main() {
    a, b := swap("hello", "world")
    fmt.Println(a, b)
    fmt.Println(split(17))
}
