package main

import "fmt"

func main() {
	var m map[string]int = make(map[string]int, 10)
	m["one"] = 1
	fmt.Println(m)
	fmt.Println(len(m))
}
