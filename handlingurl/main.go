package main

import (
	"fmt"
	"net/url"
)
const myurl string = "https://chaicode.com/"

func main() {
	fmt.Println("welcome to handling urls in go")

	//parsing
	result,_:= url.Parse(myurl)

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qparams:= result.Query()
	fmt.Printf("The type of query params are : %T\n",qparams)
	fmt.Println(qparams["coursename"])

}
