package main

import "fmt"

type Game struct {
    Price,Rating float32
}

var m map[string]Game

func main() {
    m = make(map[string]Game)
    m["Metroid"] = Game{
        Price: 20.0, Rating: 10.0,
    }
    fmt.Println(m["Metroid"])
}
