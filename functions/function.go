package main

import "fmt"

func main() {
	fmt.Println("functions in go")
	greeter()
	result := adder(3, 5)
	println("result is : ", result)

	proRes := proadder(2, 5, 6, 7, 8)
	fmt.Println("pro result is:= ", proRes)

}
func freetertwo() {
	fmt.Println("Test")
}

// need to mention thereturn type such as int in this case
func adder(val1 int, val2 int) int {
	return val1 + val2
}

func proadder(values ...int) int {
	total := 0
	for _, val := range values {
		total = total + val
	}
	return total
}

func greeter() {
	fmt.Println("Namastey from golang")
}
