package main

import "fmt"

func main() {
	languages := make(map[string]string)
	languages["js"] = "javascript"
	languages["C"] = "C++"

	fmt.Println("all values in map:-", languages)
	fmt.Println("only one key value:-", languages["js"])

	// to delete a key
	delete(languages, "C")

	fmt.Println(languages)

	// loops are interseting
	// loops for map
	for key, value := range languages {
		fmt.Printf("for key %v , value is %v", key, value)
	}

}
