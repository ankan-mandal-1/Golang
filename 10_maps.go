package main

import "fmt"

func main() {
	m := make(map[string]string)

	m["name"] = "golang"
	m["area"] = "backend"

	fmt.Println(m["name"], m["area"])

	delete(m, "name")
	clear(m)

	object := map[string]int{"price": 40, "phone": 3}
	fmt.Println(object)

	v, ok := object["price"]

	fmt.Println(v)

	if ok {
		fmt.Println("All OK")
	} else {
		fmt.Println("Not Ok")
	}
}
