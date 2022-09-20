package main

import "fmt"

func main() {
	array := [5]float64{1, 2, 3, 4, 5}
	slice := array[:1]
	fmt.Println(slice)
	// fmt.Println("len(array)", len(array))
	// fmt.Println("cap(array)", cap(array))
	fmt.Println("len(slice)", len(slice))
	fmt.Println("cap(slice)", cap(slice))
}
