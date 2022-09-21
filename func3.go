package main

import "fmt"

func main() {
	fmt.Println(sum([]int{10, 20, 30, 40}))
}

func sum(items []int) int {
	result := 0
	for _, value := range items {
		result += value
	}
	return result
}
