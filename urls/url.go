package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=adwadadws"

func main() {
	fmt.Println("url", myurl)
	// parsing
	result, _ := url.Parse(myurl)
	fmt.Println(result.Port())
	fmt.Println(result.Path)
	fmt.Println(result.RawQuery)
	fmt.Println(result.Host)
	fmt.Println(result.Scheme)
	qparams := result.Query()
	fmt.Println("type of params are %T\n ", qparams)
	fmt.Println(qparams["coursename"])

	// constructing the url

	partsofUrl := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "user=hitesh",
	}
	anotherUrl := partsofUrl.String()
	fmt.Println(anotherUrl)
}
