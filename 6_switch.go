package main

import (
	"fmt"
	"time"
)

func main() {
	i := 5

	switch i {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	default:
		fmt.Println("Other")
	}

	// Multiple Condition Switch

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's weekend")
	default:
		fmt.Println("It's workday")
	}

	// type switch
	whoAmI := func(i interface{}) {
		switch t := i.(type) { // switch i.(type) {}
		case int:
			fmt.Println("It's an integer")
		case string:
			fmt.Println("It's a String")
		case bool:
			fmt.Println("It's a boolean")
		default:
			fmt.Println("Other", t)
		}
	}

	whoAmI("Golang")
}
