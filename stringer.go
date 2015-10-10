package main

import "fmt"

// example 1
type Person struct {
    Name string
    Age  int
}

func (p Person) String() string {
    return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

// example 2
type IPAddr [4]byte

func (ip IPAddr) String() string {
    return fmt.Sprintf("%d, %d, %d, %d", ip[0], ip[1], ip[2], ip[3])    
}

func main() {
    a := Person{"Arthur Dent", 42}
    z := Person{"Zaphod Beeblebrox", 9001}
    fmt.Println(a, z)
    

    addrs := map[string]IPAddr{
        "loopback":  {127, 0, 0, 1},
        "googleDNS": {8, 8, 8, 8},
    }
    for n, a := range addrs {
        fmt.Printf("%v: %v\n", n, a)
    }
}
