package main

import(
	"fmt"
	"sync"
)

type Chop struct{sync.Mutex}

type Philo struct{
	left, right *Chop
}

var mut sync.Mutex

func (p Philo) Eat(x int) {
	var in string
	var pass = true
	for pass==true{
		mut.Lock()
		fmt.Println("Can the ",x,"Philosopher eat? type y or n")
		fmt.Scan(&in)
		if in == "y"{
			for i:=0;i<3;i++{
				p.left.Lock()
				p.right.Lock()
				fmt.Println("Starting to eat <",x,">")
				fmt.Println("Finishing  eating <",x,">")
				p.left.Unlock()
				p.right.Unlock()
			}
			pass = false
		}
		mut.Unlock()
	}
	wg.Done()
}

var wg sync.WaitGroup

func main(){
	Strick := make([]*Chop,5)
	for i:=0;i<5;i++{
		Strick[i]=new(Chop)
	}

	philo := make([]*Philo,5)
	for i:=0;i<5;i++{
		philo[i] = &Philo{Strick[i],Strick[(i+1)%5]}
	}
	wg.Add(5)
	for i:=0;i<5;i++{
		go philo[i].Eat(i+1)
	}
	wg.Wait()
	fmt.Println("Done every philosopher have eaten")

}
