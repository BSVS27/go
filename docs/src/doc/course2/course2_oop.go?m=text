package main

import (
	"fmt";
	"math";
	"data";
)

type MyInt int

func (mi MyInt) Double() int{
	return int(mi*2)
}

type Point struct{
	x float64
	y float64
}

func (p Point) DisToOrig() float64 {
	t := math.Pow(p.x,2)+math.Pow(p.y,2)
	return math.Sqrt(t)
}

type Shape2D interface{ <span class="comment">//polymorphism or  to create a function that can resive diffentes types of inputs</span>
	Area() float64
	Perim() float64
}

type Triangle struct{
	base float64
	height float64
}

type Circle struct{
	radius float64
}

func (t Triangle) Area() float64{
	return t.base*t.height
}

func (t Triangle) Perim() float64{
	return t.base*3
}

func (c Circle) Area() float64{
	return c.radius*math.Pi*c.radius
}

func (t Triangle) Side() float64{
	return t.base 
}

func (c Circle) Perim() float64{
	return 2*c.radius*math.Pi
}

func measure(s Shape2D){
	fmt.Println(s)
	fmt.Println(s.Area())
}

func FitInYard(s Shape2D) bool {
	if s.Area() &gt; 100 &amp;&amp; s.Perim() &gt; 100 {
		return true
	}
	return false
}

func PrintMe(val interface{}){ <span class="comment">//It is a function that can resive any type of input</span>
	fmt.Println(val)
}

func DrawShape(s Shape2D) {
	tri, ok := s.(Triangle)
	if ok{
		fmt.Println(&#34;It is a Triangle &#34;,tri)
	}
	cir, ok := s.(Circle)
	if ok {
		fmt.Println(&#34;It is a Circle &#34;, cir)
	}

	switch s.(type){
	case Triangle:
		fmt.Println(&#34;Triangle&#34;)
	case Circle:
		fmt.Println(&#34;Circle&#34;)
	}
}

func main() {
	var v1 MyInt
	v1 = 2
	fmt.Println(v1)
	v := MyInt(3)
	fmt.Println(v.Double())
	fmt.Println(v)
>	
>	var p Point
>	p.x = 3
>	p.y = 4
>	fmt.Println(p.DisToOrig())
>	data.PrintX()
>
>	var ca data.Casa
>	ca.InitMe(5,5)
>	ca.X = 0
>	fmt.Println(ca.PrintCasa())
>	
>	t := Triangle{2,5}
>	c := Circle{100}
>
>	fmt.Println(t.Area())
>	fmt.Println(c.Area())
>
>	measure(t)
>	measure(c)
>
>	var s1 Shape2D
>	d1 := Circle{2}
>	s1 = d1 <span class="comment">//because s1 has a method of d1</span>
>	fmt.Println(s1.Area())
>
>	fmt.Println(FitInYard(c))
>	DrawShape(c)
>	PrintMe(123)
>}
