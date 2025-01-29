package main

import "fmt"

func main() {
	fmt.Println("structs in go")
	//no inheritance in go: No super or parent

	ranjan:= User{"Ranjan","ranjan@gmail.com",true,23}
	fmt.Println(ranjan)
	fmt.Printf("Ranjan details are: %+v\n",ranjan)
}

type User struct{
	Name string
	Email string
	Status bool
	Age int
}