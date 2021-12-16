package main

import (
	"fmt"
	"math/rand"
)

func main() {
	mainGame()
}

func mainGame() {
	var yourNum, myNum int
	fmt.Println("I have a num[0-100], Pleasr enter the first num what you guess!")
	myNum = setNum()
	yourNum = getNum()

PD:
	if myNum > yourNum {
		fmt.Println("Your num is small, Please add")
		yourNum = getNum()
		goto PD
	} else if myNum < yourNum {
		fmt.Println("Your num is big, Pless reduce")
		yourNum = getNum()
		goto PD
	} else {
		fmt.Printf("You Win!\nThe num is %d\n", yourNum)
		playAgain()
	}
}

func setNum() int {
	return rand.Intn(100)
}

func getNum() int {
	var num int
	fmt.Scanln(&num)
	return num
}

func playAgain() {
	var anser string
	fmt.Println("Are you what to play again?(yes/no)")
	fmt.Scanln(&anser)
	if anser == "yes" {
		mainGame()
	} else if anser == "no" {
		fmt.Println("Thinks for Play")
	} else {

	}
}
