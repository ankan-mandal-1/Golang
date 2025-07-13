package main

import "fmt"

func main() {
	// while loop
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// infinite loop
	for {
		println("1")
	}

	// FOR Loop
	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}

	for i := range 5 {
		fmt.Println(i)
	}
}
