package main

import "fmt"

func main() {
	s := []int{1, 2, 3}
	fmt.Println(sum(s...))
}

func sum(items ...int) (result int, ok bool) {
	result = 0
	for _, value := range items {
		result += value
	}
	return result, true
}
