package main

import "fmt"

func main() {

	fmt.Println("welcome to array in go")

	var fruitList [4]string
	fruitList[0] = "apple"
	fruitList[1] = "tomato"
	fruitList[3] = "Grapes"

	fmt.Println("Fruit list is: ",fruitList);
	fmt.Println("lengthL ", len(fruitList));

}