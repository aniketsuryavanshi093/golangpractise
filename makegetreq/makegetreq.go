package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	request()
}

func request() {
	const url = "http://localhost:8000/"

	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	fmt.Println("bytecount is :", response.ContentLength)
	fmt.Println("status is :", response.StatusCode)
	var responseString strings.Builder
	res, _ := ioutil.ReadAll(response.Body)
	bytecontent, _ := responseString.Write(res)
	fmt.Println("byte content is :=", bytecontent)
	fmt.Println("response string :=-", responseString.String())
}
