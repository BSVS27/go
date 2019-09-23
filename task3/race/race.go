package main

import (
	"fmt"
)

func getNumber() int{
	var x int
	go func(){
		x=10
	}()
	return x
}

func main(){
	fmt.Println(getNumber())
}
