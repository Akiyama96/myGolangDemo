package main

import "fmt"

var nums = []int{1, 2, 3, 5, 6, 7, 9}
var k = 4

func rotate(nums []int, k int) {
	newNums := make([]int, len(nums))
	for i, v := range nums {
		newNums[(i+k)%len(nums)] = v
		fmt.Println((i + k) % len(nums))
	}
	copy(nums, newNums)
}

func main() {
	rotate(nums, k)
}
