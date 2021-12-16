package main

import "fmt"

type Inter interface {
	ping()
	pang()
}

type Anter interface {
	Inter
	String()
}

type St struct {
	Name string
}

func (St) ping() {
	println("ping")
}

func (*St) pang() {
	println("pang")
}

func main() {
	fmt.Println("The program has begin!")
	st := &St{"andes"}
	var i interface{} = st
	o := i.(Inter)
	o.ping()
	o.pang()

	s := i.(Anter)
	fmt.Printf("%s", s)
}
