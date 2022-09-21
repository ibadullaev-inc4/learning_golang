package main

import "fmt"

func main() {
	m := map[string]int{}
	fmt.Println(m)
	fil(m)
	fmt.Println(m)
}

func fil(m map[string]int) {
	m["one"] = 1
}
