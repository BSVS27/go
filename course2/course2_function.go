package main

import (
	"fmt"
)

func PrintHello() {
	fmt.Println("Hello")
}

func foo(x int, y int) (int,int) {
	fmt.Println(x * y, x + y)
	return x*y, x+y
}

func pointer(y *int) {
	*y = *y + 1
}

func array(x [3]int) []int{
	return x[1:]
}

func pointerarray(x *[3]int) {
	(*x)[0] = 100
}

func pointerslice(sli []int){
	sli[0] = sli[0] + 23
}

func incFn(x int) int {
	return x+1
}

func applyIt(afunct func (int) int, val int) int {
	return afunct(val)
}

func getMax(vals ...int) int {
	maxV := -1
	for _, v := range vals{
		if v > maxV{
			maxV = v
		}
	}
	return maxV
}

func main() {
	PrintHello()
	x,y := foo(2,3)
	fmt.Println(x,y)

	pointer(&x)
	fmt.Println(x)

	a := [3]int{1,2,3}
	fmt.Println(array(a))
	
	a1 := [3]int{1,2,3}
	pointerarray(&a1)

	fmt.Println(a1)
	
	b := []int{1,2,3}
	pointerslice(b)
	fmt.Println(b)
	

	var funcVar func(int) int //variable for a function
	funcVar = incFn
	fmt.Println(funcVar(1))
	fmt.Println(applyIt(incFn,2))

	slice := []int{1,3,6,4}
	fmt.Println(getMax(slice...))
	i := 1
	fmt.Print(i)
	i++
	defer fmt.Print(i+1)
	fmt.Print(i)
}
