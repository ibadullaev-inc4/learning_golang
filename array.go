package main

import "fmt"

func main() {
	var point1 [3]float64
	// point1 = [3]float64{1, 2, 3}
	var point2 [2]int64 = [2]int64{100, 200}
	var users [2]string = [2]string{"nariman", "aydin"}
	var active [2]bool = [2]bool{true, false}
	test := [3]float64{1}

	fmt.Println(point1)
	fmt.Println(point2)
	fmt.Println(users)
	fmt.Println(active)
	fmt.Println(test)

	fmt.Println(point1[0])
	fmt.Println(point2[1])
	fmt.Println(users[0])
	fmt.Println(active[1])
	fmt.Println(test[2])
}
