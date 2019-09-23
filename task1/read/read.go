package main

import (
	"fmt"
	"bufio"
	"os"
)

type Name struct{
	fname string
	lname string
}

func main() {
	sli := make([]Name,0)
	var n1 Name

	fmt.Println("Type the name of the text file: ")
	var file string
	fmt.Scan(&file)

	f, err := os.Open(file)

	if err == nil{
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			name := scanner.Text()
			fmt.Println(name)
			for i:=0;i<len(name);i++{
				if string(name[i]) == " "{
					n1.fname = name[:i]
					n1.lname = name[i:]
					sli = append(sli,n1)
				}
			}
		}
		var n2 Name
		for _, val := range sli{
			n2 = val
			fmt.Println("First name: ", n2.fname," Last name: ", n2.lname)
		}
		f.Close()
	} else{
		fmt.Println("There is a problem opening the file: ", file)
	}
}
