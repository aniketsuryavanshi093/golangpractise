package main

import (
	"fmt"
	"sort"
)

func main() {
	// var fruitList [4]string in array we need to define the length ;
	var fruitlistslice = []string{"banana", "orange", "watermelon", "mango"}
	// in slices we do not define the length of the variable so its called as slices and we need to initialize at start;

	// to add values in slice we use append method
	fruitlistslice = append(fruitlistslice, "papaya", "fig")
	// fruitList[0] = "a"
	// fruitList[1] = "b"
	// fruitList[3] = "c"

	fmt.Println("the fruitlist is :=", fruitlistslice)
	fmt.Println("the len fruitlist is :=", len(fruitlistslice))
	// this syntax is used to slice or split the slice from position
	fruitlistsubs := append(fruitlistslice[1:3])
	// var veggies = [5]string{"a", "b", "c", "d"}

	fmt.Println("fruitlistsubs are :=", fruitlistsubs)

	scores := make([]int, 4)
	scores[0] = 1
	scores[1] = 9
	scores[3] = 3
	scores[2] = 4
	// scores[5] = 4
	fmt.Println("scores before static range :=", scores)

	scores = append(scores, 5, 6)

	fmt.Println("scores after using append range are :=", scores)

	// sorting

	sort.Ints(scores)
	fmt.Println("sorted", scores)

}
