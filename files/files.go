package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	content := "test"
	file, err := os.Create("./mytestfile.txt")
	if err != nil {
		panic(err)
	}

	length, er := io.WriteString(file, content)

	checkNilErr(er)

	fmt.Println("length is:", length)
	defer file.Close()
	readFile("./mytestfile.txt")
}

func readFile(filename string) {
	dataBytes, err := ioutil.ReadFile(filename)
	checkNilErr(err)
	fmt.Println("text data in file \n ", string(dataBytes))

}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
