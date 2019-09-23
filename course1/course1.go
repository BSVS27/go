package main

import (
	"fmt"
	"encoding/json"
	"io/ioutil"
	"os"
)

type Person struct { //STRUCT
	name string
	addr string
	phone int
}

func main() {
	fmt.Println("Hello World")
	var var1 int32
	var1 = 2
	var2 := 3
	fmt.Println(var1, var2)
	var3 := "casa"
	var4 := (var3 + "hola")
	fmt.Println(var4)

	var x int = 1
	var y int
	var ip *int

	ip = &x
	y = *ip

	fmt.Println(y)

	ptr := new(int) //Create a pointer
	*ptr = 3
	fmt.Println(ptr)

	var x1 int32 = 5
	x2 := int64(x1) //convert variables
	fmt.Println(x2)
	fmt.Println("VAR3",string(var3[0]), var3[0])

	const xx = 1.3
	const (
		yy = 4
		z = "H"
	
	)

	type Grades int32
	const (
		Monday Grades = iota
		Tuesday
		Wednesday
		thursday
	)

	if Monday == 0 {
		fmt.Println("Monday es 0")
	} else {
	 fmt.Println("Monday is not 0")
	}
	
	for i:=0;i<5;i++{
		fmt.Println(i)
		if i == 3{
			fmt.Println("Saliendo del for")
			break
		}
	}

	switch xx {
	case 1.3:
		fmt.Println("Es 1.3")
	default:
		fmt.Println("No es 1.3")
	}
	
	fmt.Println("Number of apples?")

	var appleNum int

	fmt.Scan(&appleNum) // To read de input

	fmt.Println("The number of apples are: ", appleNum)

	var array [3]int = [3]int{1,2,3}

	array1 := [...]int{0,1}
	
	fmt.Println(array[0],array1[0],len(array))

	for i, v := range array{
		fmt.Println("Index: ", i,"value: ", v)
	}

	arr := [...]string{"a","b","c","d"}

	//s1 := arr[1:2]
	s2 := arr[2:]

	fmt.Println(/*s1,*/"   ", s2, "   ", len(s2), cap(s2))

	sli := []int{1,2,3}

	sli1 := make([]int,3)

	fmt.Println(sli1,sli)

	sli1 = append(sli1, 100)

	fmt.Println(sli1)

	mapp := map[string]int{
		"joe":123, 
		"jane": 331}
	mapp["joe"] = 100
	delete(mapp,"jane") //to delete a key
	_ , p := mapp["jane"] //p is true if jane is a key inside mapp
	fmt.Println(mapp["joe"],mapp,p)

	for key, val := range mapp {
		fmt.Println(key, val)
	}

	var p1 Person
	p1.name = "joe"
	p1.addr = "a st"
	p1.phone = 123

	fmt.Println(p1)

	barr, err := json.Marshal(p1) //STRUCT TO JSON

	var p2 Person
	err1 := json.Unmarshal(barr, &p2)

	fmt.Println(barr,err,p2,err1)

	dat, _ := ioutil.ReadFile("test.txt")

	fmt.Println(string(dat),len(dat))

	//err2 := ioutil.WriteFile("test.txt",byte("Bye World"),0777)
	
	f, _ := os.Open("test.txt")
	barr1 := make([]byte,78)
	nb, _ := f.Read(barr1)
	//os.Write()
	fmt.Println(string(nb))
	f.WriteString("Hi")
	f.Close()



	var yyy = []int{1,2,3,4,5}
	fmt.Println(yyy)
	yyy = yyy[:0]
	fmt.Println(yyy)
	yyy = append(yyy,4)
	fmt.Println(yyy)

	//var rt = [3]int{1,2,3}
	//rt1, rt2 := rt

	var rtrt = "a"
	//dada := fmt.Sprintf("%s%s", rtrt, "a")
	dada := rtrt + "a"
	fmt.Println(dada)
}
