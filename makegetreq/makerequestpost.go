package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	// postjsonrequest()
	postjsonformdatarequest()
}

func postjsonrequest() {
	const myurl = "http://localhost:8000/post"

	// fake json payload

	requestBody := strings.NewReader(`
		{
			"coursename" : "Lets go with go lang",
			"price":0,
			"platform":"learncodeonline.in"
		}
	`)

	response, err := http.Post(myurl, "application/json", requestBody)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	result, err := ioutil.ReadAll(response.Body)
	fmt.Println(string(result))
}

func postjsonformdatarequest() {
	const myurl = "http://localhost:8000/postform"

	// formdata
	data := url.Values{}
	data.Add("firstname", "Aniket")
	data.Add("lastname", "suryavanshi")
	data.Add("email", "hitesh@dev.com")

	response, err := http.PostForm(myurl, data)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	content, err := ioutil.ReadAll(response.Body)

	fmt.Println(string(content))
}
