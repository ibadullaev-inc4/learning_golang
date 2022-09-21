package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5}
	fmt.Println(s)
	zerofy(s)
	fmt.Println(s)
}

func zerofy(s []int) {
	s = append(s, 6)
	for i, _ := range s {
		s[i] = 0
	}
}
