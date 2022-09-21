package main

import "fmt"

func main() {
	l := pr()
	fmt.Println(l)
}

func pr() func() {
	return callBack
}

func callBack() {
	fmt.Println("2")
}
