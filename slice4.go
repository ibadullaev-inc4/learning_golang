package main

import "fmt"

func main() {
	// s := []int{1, 2, 3}

	// c := 10
	// s := make([]int, 0, c)

	var s []int

	fmt.Println(s)
	fmt.Println("len(s)", len(s))
	fmt.Println("cap(s)", cap(s))
}
