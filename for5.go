package main

import "fmt"

func main() {
	var sum int = 0

outer:
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			// break outer
			continue outer
		}
		if i == 5 {
			continue
		}
		if i%2 == 1 {
			fmt.Println(i, "is odd")
		} else {
			fmt.Println(i, "is even")
		}
		sum += i
	}

	fmt.Println("sum is: ", sum)
}
