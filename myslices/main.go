package main

import "fmt"

func main() {
	fmt.Println("welcome to slices in go");

	var fruitList=[]string{}
	fmt.Printf("Type of fruitList is %T\n",fruitList);

	fruitList = append(fruitList,"Mango","Banana")
	fmt.Println(fruitList)

	highscores:=make([]int,4)
	highscores[0]=234
	highscores[1]=334
	highscores[2]=434
	highscores[3]=534
	fmt.Println(highscores);

	var courses =[]string{"reactjs","javascript","swift","python","ruby","golang"};
	fmt.Println(courses)
	var index int =2
	courses = append(courses[:index],courses[index+1:]...)
	fmt.Println(courses)

}