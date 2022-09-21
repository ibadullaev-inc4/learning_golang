package main

import "fmt"

func main() {
	strToNum := make(map[string]int)
	// fmt.Println("len(strToNum)=", len(strToNum))
	strToNum["zero"] = 0
	strToNum["one"] = 1
	strToNum["two"] = 2
	strToNum["four"] = 4
	strToNum["five"] = 5

	// fmt.Println("len(strToNum)=", len(strToNum))

	for key, value := range strToNum {
		fmt.Printf(`"%s" = %d`+"\n", key, value)
	}

	slice := []string{"zero", "one", "two", "threee", "four"}
	for key, value := range slice {
		fmt.Printf(`"%d" = "%s"`+"\n", key, value)
	}

}
