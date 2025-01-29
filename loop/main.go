package main

import "fmt"

func main() {
	fmt.Println("welcome to loops in go")

	days := []string{"sun","mon","tues","wed","thurs","fri","sat"}
	fmt.Println(days)

	// for d:=0;d<len(days);d++ {
	// 	fmt.Println(days[d])
	// }

	for i:=range days {
		fmt.Println(days[i])
	}

	count:=1;
	for count<10{
		if count==3 {
			goto lco;
		}
		fmt.Println("value is: ",count);
		count++;
	}

	lco:
	  fmt.Println("helod shhdduh hdhsh")
}