package main

import "fmt"

func main() {

	var a int = 10
	var ip *int

	ip = &a

	fmt.Printf("a ip is:%x\n", &a)
	fmt.Printf("ip's ip:%x\n", ip)
	fmt.Printf("*ip's value:%d\n", *ip)

}
