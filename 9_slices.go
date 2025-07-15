package main

import "fmt"

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
}
