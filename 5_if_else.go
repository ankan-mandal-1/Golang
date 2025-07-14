package main

import "fmt"

func main() {
	age := 18

	if age >= 18 {
		fmt.Println("Person is Adult")
	} else {
		fmt.Println("Person is child")
	}

	if age >= 18 {
		fmt.Println("Person is Adult")
	} else if age >= 12 {
		fmt.Println("Person is Teenage")
	} else {
		fmt.Println("Person is kid")
	}

	var role = "admin"
	var hasPermission = true

	if role == "admin" || hasPermission { // || or && and
		fmt.Println("Yes")
	}

	if num := 15; num >= 18 {
		fmt.Println("Person is an Adult", num)
	} else if num >= 12 {
		fmt.Println("Person is teenage", num)
	}
}
