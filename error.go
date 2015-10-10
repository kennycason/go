package main

import (
    "fmt"
    "time"
    "strconv"
)

type MyError struct {
    When time.Time
    What string
}

func (e *MyError) Error() string {
    return fmt.Sprintf("at %v, %s",
        e.When, e.What)
}

func run() error {
    return &MyError{
        time.Now(),
        "it didn't work",
    }
}

func main() {
    if err := run(); err != nil {
        fmt.Println(err)
    }

    // example 2
    i, err := strconv.Atoi("42df")
    if err != nil {
        fmt.Printf("couldn't convert number: %v\n", err)
    } else {
        fmt.Println("Converted integer:", i)
    }
}
