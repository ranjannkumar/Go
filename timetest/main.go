package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("time testing in go")
	presentTime := time.Now()
	fmt.Println(presentTime)
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))
}