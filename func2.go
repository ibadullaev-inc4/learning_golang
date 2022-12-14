package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(sum(s...))
}

func sum(items ...int) (int, bool) {
	result := 0
	for _, value := range items {
		result += value
	}
	return result, true
}
