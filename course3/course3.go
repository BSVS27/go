package main

import (
	"fmt"
	"sync"
)

var x int = 0

func foo(wg *sync.WaitGroup,y int) {
	x = y
	fmt.Println("GO :routine done")
	wg.Done()
}

func prod(v1 int,v2 int, c chan int){
	c <- v1*v2 //Block when the buffer is full
}

func setup(){
	fmt.Println("Init")
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

	go foo(&wg, 3)
	wg.Wait()
	fmt.Println("Test Done")

	c := make(chan int)
	go prod(1,2,c)
	go prod(3,4,c)
	var a int
	a = <- c
	b := <- c //Block when the buffer is empty
	fmt.Println(a*b)
	c1 := make(chan int)
	go prod(1,8,c)
	go prod(9,0,c1)
	
	select{ //To use the channel that return faster
		case a1 := <-c:
			fmt.Println(c,a1)
		case b1 := <-c1:
			fmt.Println(c1,b1)
		default: //If I use the default I am not blocking
			fmt.Println("Nothing")
	}
	/*select{ //To select if send or receive something
		case a = &lt;- inchan:
			fmt.Println(Received a)
		case outchan &lt;- b:
			fmt.Println(Sent b)
	}
	*/
	var wg1 sync.WaitGroup
	var on sync.Once
	wg1.Add(1)
	var mut sync.Mutex
	go inc(&mut,&wg1,&on)
	go inc(&mut,&wg1,&on)
	wg1.Add(1)
	wg1.Wait()
	fmt.Println(y)

}
