package main

import "fmt"

type PointWithComments struct {
	Point    Point
	Comments []string
}

type Point struct {
	X int
	Y int
}

func main() {
	p := Point{X: 1}
	fmt.Println(p)
	fmt.Printf("%+v\n", p)
	fmt.Printf("%#+v\n", p)

	var pc PointWithComments
	fmt.Printf("%+v\n", pc)
	fmt.Println(pc.Comments)
	fmt.Println(pc.Point.X)
	fmt.Println(pc.Point.Y)

	fmt.Println(p == Point{1, 0})
}
