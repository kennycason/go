package main

import (
    "fmt"
)

type Parent struct {
    Foo string
}

func (p *Parent) bar() {
	fmt.Println("bar1")
}

type Child struct {
	Parent
	Foo string
}

func (c *Child) bar() {
	fmt.Println("bar2")
}

func main() {
	parent := Parent{"foo1"}
	child := Child{Foo : "foo1"}
	//child := Child{"foo1"}
	fmt.Println(parent)
	fmt.Println(child)
	parent.bar()
	child.bar()
}