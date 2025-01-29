package main

import "fmt"

func main() {
	fmt.Println("If else in go");

	loginCount := 1
	var res string
	if loginCount <10 {
		res="Regular user"
	}else if loginCount>10{
		res="wacth out"
	}else{
		res="exactly 10"
	}

	fmt.Println(res)
}