package main

import "fmt"

func main() {
	fmt.Println("structs in go")
	//no inheritance in go: No super or parent

	ranjan:= User{"Ranjan","ranjan@gmail.com",true,23}
	fmt.Println(ranjan)
	fmt.Printf("Ranjan details are: %+v\n",ranjan)
	ranjan.GetStatus()
	ranjan.NewMail()
}

type User struct{
	Name string
	Email string
	Status bool
	Age int
}

func (u User) GetStatus(){
	fmt.Println("Is user active",u.Status)
}

func (u User) NewMail(){
	u.Email = "test@go.dev"
	fmt.Println("Email for this user is : ",u.Email)
}