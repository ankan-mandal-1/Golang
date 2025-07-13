package main

import "fmt"

func main() {
	const name string = "golang"

	const (
		port = 5000
		host = "localhost"
	)

	fmt.Println(port, host)
}
