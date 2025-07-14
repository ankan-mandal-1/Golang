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
	for j := 0; j < 3; j++ {
		fmt.Println(j)
	}

	for k := range 5 {
		fmt.Println(k)
	}
}
