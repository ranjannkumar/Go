package main
import "fmt"

func main()  {
	fmt.Println("Welcome to pointer in go")
	myNumber:=23
	var ptr = &myNumber
	fmt.Println("value of actual pointer is", ptr)
	fmt.Println("value of actual pointer is", *ptr)


}