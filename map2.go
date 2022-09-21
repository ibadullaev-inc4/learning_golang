package main

import "fmt"

func main() {
	users := map[string]int{
		"nariman": 35,
		"aydin":   38,
		"babak":   28,
	}
	fmt.Println(users)

	name, ok := users["nariman"]
	if ok {
		fmt.Println(name)
	} else {
		fmt.Println("Can't find key")
	}

	fmt.Println(name, ok)
}
