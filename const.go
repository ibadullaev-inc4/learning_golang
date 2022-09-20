package main

import "fmt"

const (
	RED = iota + 1
	BLUE
	GREEN
	BLACK
	_
	WHITE
	_
	GRAY
	YELLOW = 100
)

func main() {
	fmt.Println(RED)
	fmt.Println(BLUE)
	fmt.Println(GREEN)
	fmt.Println(BLACK)
	fmt.Println(WHITE)
	fmt.Println(GRAY)
	fmt.Println(YELLOW)
}
