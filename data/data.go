package data

import (
	"fmt"
)

var x = 1
func PrintX() {
	fmt.Println(x)
}

type Casa struct{
	X float64
	y float64
}

func (p *Casa) InitMe(xn, yn float64){
	p.X = xn
	p.y = yn
}

func (p Casa) PrintCasa() float64{
	return p.X*p.y
}
