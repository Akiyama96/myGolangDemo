package main

import "fmt"

func main() {
	var a int = 20
	var ip *int
	ip = &a
	fmt.Printf("This is \"a\"'s ip:%s\n", &a)
	fmt.Printf("This is value:%d\n", *ip)
}
