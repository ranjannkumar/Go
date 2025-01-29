package main

import "fmt"

func main() {
	fmt.Println("welcome to functions in golang")
	greeter()
	res:=adder(5,2)
	fmt.Println("Result is :",res)
	res2,myMessage:=proadder(2,5,8,18);
	fmt.Println("res2 is: ",res2)
	fmt.Println("res2 is: ",myMessage)

}

func greeter(){
	fmt.Println("hello")
}
func adder(v1 int,v2 int)int{
	return v1+v2
}

func proadder(values ...int)(int,string){
	total:=0;
	for _,val:=range values{
		total+=val;
	}
	return total,"hello mfs"
}