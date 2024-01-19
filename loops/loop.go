package main

import "fmt"

func main() {
	days := []string{"sunday", "friday", "tuesday"}
	fmt.Println((days))

	// simple loop
	for d := 0; d < len(days); d++ {
		fmt.Println(days[d])
	}

	//some advance

	for i := range days {
		fmt.Println(days[i])
	}

	// for each type

	// for index, day := range days {
	// 	fmt.Println("index is %v , value is %v", index, day)
	// }

	// while loop
	roughvalue := 1
	for roughvalue < 10 {
		fmt.Println("values is :", roughvalue)
		roughvalue++
	}
}
