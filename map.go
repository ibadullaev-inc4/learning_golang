package main

import "fmt"

func main() {
	strToNum := map[string]int{
		"one":    1,
		"two":    2,
		"threee": 3,
	}
	fmt.Println(strToNum)
	fmt.Println(strToNum["two"])
	fmt.Println(strToNum["four"])

	// for key, val := range strToNum {
	// 	fmt.Println(key)
	// 	fmt.Println(val)
	// }
}
