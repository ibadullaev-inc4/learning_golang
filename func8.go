package main

import "fmt"

func main() {
	s := []int{1, 2, 3}
	fmt.Println(s)
	zerofy(s)
	fmt.Println(s)
}

func zerofy(s []int) {
	for i := range s {
		s[i] = 0
	}
}
