package main

import "fmt"

func main() {
	var fruitList [4]string
	fruitList[0] = "a"
	fruitList[1] = "b"
	fruitList[3] = "c"

	fmt.Println("the fruitlist is :=", fruitList)
	fmt.Println("the len fruitlist is :=", len(fruitList))

	var veggies = [5]string{"a", "b", "c", "d"}

	fmt.Println("vegies are :=", veggies)

}
