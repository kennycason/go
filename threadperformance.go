package main

import (
    "fmt"
    "time"
)

func runTests() int64 {
    start := time.Now()
    for i := 0; i < 1000; i++ {
        go func() {}()
    }
    end := time.Now()
    fmt.Println((end.UnixNano() - start.UnixNano()) / 1000)
    return (end.UnixNano() - start.UnixNano()) / 1000
}

func main() {
    n := 10
    var total int64 = 0
    for i := 0; i < n; i++ {
        total += runTests()
    }
    fmt.Println(total / int64(n))
}
