package main

import (
	"fmt"
	"sync"
)

var x int = 0

func foo(wg *sync.WaitGroup,y int) {
	x = y
	fmt.Println(&#34;GO :routine done&#34;)
	wg.Done()
}

func prod(v1 int,v2 int, c chan int){
	c &lt;- v1*v2 <span class="comment">//Block when the buffer is full</span>
}

func setup(){
	fmt.Println(&#34;Init&#34;)
}

func inc(mut *sync.Mutex, wg1 *sync.WaitGroup, on *sync.Once){
	on.Do(setup)
	mut.Lock()
	y = y+1
	mut.Unlock()
	wg1.Done()
}

var y int = 0

func main(){
	var wg sync.WaitGroup
	wg.Add(1)

	go foo(&amp;wg, 3)
	wg.Wait()
	fmt.Println(&#34;Test Done&#34;)

	c := make(chan int)
	go prod(1,2,c)
	go prod(3,4,c)
	var a int
	a = &lt;- c
	b := &lt;- c <span class="comment">//Block when the buffer is empty</span>
	fmt.Println(a*b)
	c1 := make(chan int)
	go prod(1,8,c)
	go prod(9,0,c1)
	
	select{ <span class="comment">//To use the channel that return faster</span>
		case a1 := &lt;-c:
			fmt.Println(&#34;c&#34;,a1)
		case b1 := &lt;-c1:
			fmt.Println(&#34;c1&#34;,b1)
		default: <span class="comment">//If I use the default I am not blocking</span>
			fmt.Println(&#34;Nothing&#34;)
	}
	<span class="comment">/*select{ //To select if send or receive something
		case a = &lt;- inchan:
			fmt.Println(&#34;Received a&#34;)
		case outchan &lt;- b:
			fmt.Println(&#34;Sent b&#34;)
	}
	*/</span>
	var wg1 sync.WaitGroup
	var on sync.Once
	wg1.Add(1)
	var mut sync.Mutex
	go inc(&amp;mut,&amp;wg1,&amp;on)
	go inc(&amp;mut,&amp;wg1,&amp;on)
	wg1.Add(1)
	wg1.Wait()
	fmt.Println(y)

}
