package main

import "fmt"

type Point struct {
	X int
	Y int
}

func main() {
	var p Point
	fmt.Println(p)
	fill(&p)
	fmt.Println(p)
}

func fill(p *Point) {
	p.X = 1
	p.Y = 2
}
