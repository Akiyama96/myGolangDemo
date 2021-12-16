package main

import "fmt"

var (
	name string
)

type People interface {
	SayName()
}

type S struct{}

func (S) SayName() {
	fmt.Println(name)
}

func main() {
	name = "Mio"
	var i People
	i = S{}
	i.People()
}
