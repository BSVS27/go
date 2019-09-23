package main

import (
	"fmt"
)

func Swap(slice []int, index int) {
	var help int
	help = slice[index]
	slice[index] = slice[index+1]
	slice[index+1] = help
}

func BubbleSort(slice []int) {
	for i:=0;i<len(slice);i++{
		for j:=0;j<(len(slice)-1);j++{
			if slice[j] > slice[j+1] {
				Swap(slice,j)
			}
		}
	}
}

func main(){
	var number int
	sequence := make([]int,0)
	fmt.Println("Type a sequence of 10 integers")
	for i:=0; i < 10; i++{
		fmt.Print("Type the number ", i+1," : ")
		fmt.Scan(&number)
		sequence = append(sequence,number)
	}
	BubbleSort(sequence)

	fmt.Println("The sort sequence is: ", sequence)
}
