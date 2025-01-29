package main

import "fmt"

func main() {
	fmt.Println("welcome to maps in go")

	lan := make(map[string]string)
	lan["js"]="javascript"
	lan["RB"]="Ruby"
	lan["PY"]="python"
	lan["GO"]="Golang"

	fmt.Println("List of languages: ",lan)
	fmt.Println("Js shorts for: ",lan["js"])

	delete(lan,"RB")
	fmt.Println("List of all languages: ",lan) 


}