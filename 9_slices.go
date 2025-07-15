package main

import (
	"fmt"
	"slices"
)

func main() {

	var nums []int // value = nil
	fmt.Println(len(nums))

	var number = make([]int, 2, 5)
	arr := []int{}

	arr[0] = 3

	number = append(number, 1)
	number = append(number, 2)
	number = append(number, 3)

	fmt.Println(number)
	fmt.Println(cap(number))

	// Copy Function
	num1 := make([]int, 0, 5)
	num1 = append(num1, 2)

	var num2 = make([]int, len(num1))

	copy(num2, num1)

	// Slice Operator
	var arr1 = []int{1, 2, 3}
	fmt.Println(arr1[0:2])

	//Slices Package
	var lio1 = []int{1, 2, 3}
	var lio2 = []int{1, 2, 3}

	fmt.Println(slices.Equal(lio1, lio2))

	var odd = [][]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(odd)
}
