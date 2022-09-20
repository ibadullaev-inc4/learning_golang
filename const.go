package main

import "fmt"

const (
	RED = iota + 1
	BLUE
	GREEN
	BLACK
)

func main() {
	fmt.Println(RED)
	fmt.Println(BLUE)
	fmt.Println(GREEN)
	fmt.Println(BLACK)
}
