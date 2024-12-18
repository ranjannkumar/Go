package main

import "fmt"

func main() {
	fmt.Println("welcome to functions in golang")
	greeter()
	res:=adder(5,2)
	fmt.Println("Result is :",res)
}

func greeter(){
	fmt.Println("hello")
}
func adder(v1 int,v2 int)int{
	return v1+v2
}