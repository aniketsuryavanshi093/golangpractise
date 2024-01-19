package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("enter rating")
	input, _ := reader.ReadString('\n')
	fmt.Printf(input)
}
