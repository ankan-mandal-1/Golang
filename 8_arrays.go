package main

import "fmt"

func main() {

	var nums [4]int
	nums[0] = 345
	fmt.Println(len(nums))

	var name [3]string
	name[0] = "golang"
	fmt.Println(name, nums[0])

	numeric := [3]int{1, 2, 3}
	fmt.Println(numeric)

	// 2D Arrays
	number := [2][2]int{{3, 4}, {5, 6}}
}
