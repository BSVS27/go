package main

import (
	"fmt"
	"math"
	"time"
)

func GenDisplaceFn(acc, vel, dis float64) func (float64) float64{
	fn := func(time float64) float64{
		return (0.5)*(acc)*(math.Pow(time,2)) + vel*time + dis
	} 
	return fn
}

func main(){
	fmt.Print("Enter the acceleration: ")
	var acce float64
	fmt.Scan(&acce)

	fmt.Print("Enter the initial velocity: ")
	var vel float64
	fmt.Scan(&vel)

	fmt.Print("Enter the initial displacement: ")
	var dis float64
	fmt.Scan(&dis)

	fmt.Print("Enter the time: ")
	var tim float64
	fmt.Scan(&tim)

	time.Sleep(time.Duration(tim) * time.Second)

	Dist := GenDisplaceFn(acce,vel,dis)
	fmt.Println(Dist(tim))
}

