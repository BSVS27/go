package main

import (
	"fmt"
	"encoding/json"
	"io/ioutil"
	"os"
)

type Person struct { <span class="comment">//STRUCT</span>
	name string
	addr string
	phone int
}

func main() {
	fmt.Println(&#34;Hello World&#34;)
	var var1 int32
	var1 = 2
	var2 := 3
	fmt.Println(var1, var2)
	var3 := &#34;casa&#34;
	var4 := (var3 + &#34;hola&#34;)
	fmt.Println(var4)

	var x int = 1
	var y int
	var ip *int

	ip = &amp;x
	y = *ip

	fmt.Println(y)

	ptr := new(int) <span class="comment">//Create a pointer</span>
	*ptr = 3
	fmt.Println(ptr)

	var x1 int32 = 5
	x2 := int64(x1) <span class="comment">//convert variables</span>
	fmt.Println(x2)
	fmt.Println(string(var3[0]))

	const xx = 1.3
	const (
		yy = 4
		z = &#34;H&#34;
	
	)

	type Grades int32
	const (
		Monday Grades = iota
		Tuesday
		Wednesday
		thursday
	)

	if Monday == 0 {
		fmt.Println(&#34;Monday es 0&#34;)
	} else {
	 fmt.Println(&#34;Monday is not 0&#34;)
	}
	
	for i:=0;i&lt;5;i++{
		fmt.Println(i)
		if i == 3{
			fmt.Println(&#34;Saliendo del for&#34;)
			break
		}
	}

	switch xx {
	case 1.3:
		fmt.Println(&#34;Es 1.3&#34;)
	default:
		fmt.Println(&#34;No es 1.3&#34;)
	}
	
	fmt.Println(&#34;Number of apples?&#34;)

	var appleNum int

	fmt.Scan(&amp;appleNum) <span class="comment">// To read de input</span>

	fmt.Println(&#34;The number of apples are: &#34;, appleNum)

	var array [3]int = [3]int{1,2,3}

	array1 := [...]int{0,1}
	
	fmt.Println(array[0],array1[0],len(array))

	for i, v := range array{
		fmt.Println(&#34;Index: &#34;, i,&#34;value: &#34;, v)
	}

	arr := [...]string{&#34;a&#34;,&#34;b&#34;,&#34;c&#34;,&#34;d&#34;}

>	<span class="comment">//s1 := arr[1:2]</span>
>	s2 := arr[2:]
>
>	fmt.Println(<span class="comment">/*s1,*/</span>&#34;   &#34;, s2, &#34;   &#34;, len(s2), cap(s2))
>
>	sli := []int{1,2,3}
>
>	sli1 := make([]int,3)
>
>	fmt.Println(sli1,sli)
>
>	sli1 = append(sli1, 100)
>
>	fmt.Println(sli1)
>
>	mapp := map[string]int{
>		&#34;joe&#34;:123, 
>		&#34;jane&#34;: 331}
>	mapp[&#34;joe&#34;] = 100
>	delete(mapp,&#34;jane&#34;) <span class="comment">//to delete a key</span>
>	_ , p := mapp[&#34;jane&#34;] <span class="comment">//p is true if jane is a key inside mapp</span>
>	fmt.Println(mapp[&#34;joe&#34;],mapp,p)
>
>	for key, val := range mapp {
>		fmt.Println(key, val)
>	}
>
>	var p1 Person
>	p1.name = &#34;joe&#34;
>	p1.addr = &#34;a st&#34;
>	p1.phone = 123
>
>	fmt.Println(p1)
>
>	barr, err := json.Marshal(p1) <span class="comment">//STRUCT TO JSON</span>
>
>	var p2 Person
>	err1 := json.Unmarshal(barr, &amp;p2)
>
>	fmt.Println(barr,err,p2,err1)
>
>	dat, _ := ioutil.ReadFile(&#34;test.txt&#34;)
>
>	fmt.Println(string(dat),len(dat))
>
>	<span class="comment">//err2 := ioutil.WriteFile(&#34;test.txt&#34;,byte(&#34;Bye World&#34;),0777)</span>
>	
>	f, _ := os.Open(&#34;test.txt&#34;)
>	barr1 := make([]byte,78)
>	nb, _ := f.Read(barr1)
>	<span class="comment">//os.Write()</span>
>	fmt.Println(string(nb))
>	f.WriteString(&#34;Hi&#34;)
>	f.Close()
>
>
>
>	var yyy = []int{1,2,3,4,5}
>	fmt.Println(yyy)
>	yyy = yyy[:0]
>	fmt.Println(yyy)
>	yyy = append(yyy,4)
>	fmt.Println(yyy)
>
>	<span class="comment">//var rt = [3]int{1,2,3}</span>
>	<span class="comment">//rt1, rt2 := rt</span>
>
>	var rtrt = &#34;a&#34;
>	<span class="comment">//dada := fmt.Sprintf(&#34;%s%s&#34;, rtrt, &#34;a&#34;)</span>
>	dada := rtrt + &#34;a&#34;
>	fmt.Println(dada)
>}