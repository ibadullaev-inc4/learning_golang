package main

import "fmt"

func main() {
	fmt.Println(sum([4]int{1, 2, 3, 4}))
}

func sum(items [4]int) int {
	result := 0
	for _, value := range items {
		result += value
	}
	return result
}
