package main

import (
	"fmt"
	"math"
)

func sort(s []int, c chan []int){
	var aux int
	for i:=0;i<(len(s)-1);i++{
		for i:=0;i<(len(s)-1);i++{
			if s[i] > s[i+1]{
				aux = s[i]
				s[i] = s[i+1]
				s[i+1] = aux
			}
		}
	}
	fmt.Println(s)
	c <- s
}

func main() {
	fmt.Println("How many integers do you want to enter?")
	var in int
	fmt.Scan(&in)
	array := make([]int,0)
	for i:=0;i<in;i++{
		fmt.Println("Please enter the value ",i," of the array: ")
		var val int
		fmt.Scan(&val)
		array = append(array, val)
	}
	fmt.Println("The initial array is: ",array)
	chunk := int(math.Round(float64(in)/float64(4)))
	c := make(chan []int)
	go sort(array[:chunk],c)
	go sort(array[chunk:2*chunk],c)
	go sort(array[2*chunk:3*chunk],c)
	go sort(array[3*chunk:],c)
	a1 := <-c
	b1 := <-c
	c1 := <-c
	d1 := <-c
	s_array := make([]int,0)
	for i:=0;i<len(a1);i++{
		s_array = append(s_array,a1[i])
	}
	for i:=0;i<len(b1);i++{
		s_array = append(s_array,b1[i])
	}
	for i:=0;i<len(c1);i++{
		s_array = append(s_array,c1[i])
	}
	for i:=0;i<len(d1);i++{
		s_array = append(s_array,d1[i])
	}
	fmt.Println("The sorted array is: ",s_array)
	fmt.Println("Done")
}
