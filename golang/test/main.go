package main

import "fmt"

func main() {
	//var a string = "Mio"
	//fmt.Println(a)
	//var b, c int = 1, 2
	//fmt.Println(b, c)
	//mio()
	//mio2()
	//scan()
	/*result := 0
	for i := 1; i <= 20; i++ {
		result = dg(i)
		fmt.Printf("dg(%d) is: %d\n", i, result)
	}
	*/
	result := 0
	for i := 11; i >= 1; i-- {
		result = dj(i)
		fmt.Println(result)
	}

}
func mio() {
	var b int = 0
	if b == 2 {
		fmt.Println(b)
	} else {
		fmt.Println("b!=0")
	}
}

func mio2() {
	for true {
		fmt.Println("Just Monika")
	}
}

func scan() {
	var (
		name, age string
	)
	fmt.Println("Please enter your name and age:")
	fmt.Scanln(&name, &age)
	fmt.Printf("Hello %s %s!\n", name, age)

}

func dg(n int) (res int) {
	if n <= 1 {
		res = 1
	} else {
		res = dg(n-1) + dg(n-2)
	}
	return
}

func dj(n int) (res int) {
	if n < 1 {
		res = 1
	} else {
		res = (n - 1)
	}
	return
}
