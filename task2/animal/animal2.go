package main

import (
	"fmt"
)

type Animal interface{
	Eat()
	Move()
	Speak()
}

type Cow struct {
	food string
	locomotion string
	noise string
}

type Bird struct {
	food string
	locomotion string
	noise string
}

type Snake struct {
	food string
	locomotion string
	noise string
}
func (c Cow) Eat(){
	fmt.Println(c.food)
}
func (b Bird) Eat(){
        fmt.Println(b.food)
}
func (s Snake) Eat(){
	fmt.Println(s.food)
}
func (c Cow) Move(){
	fmt.Println(c.locomotion)
}
func (b Bird) Move(){
	fmt.Println(b.locomotion)
}
func (s Snake) Move(){
	fmt.Println(s.locomotion)
}
func (c Cow) Speak(){
	fmt.Println(c.noise)
}
func (b Bird) Speak(){
	fmt.Println(b.noise)
}
func (s Snake) Speak(){
	fmt.Println(s.noise)
}
func Eat_inter(a Animal){
	a.Eat()
}
func Move_inter(a Animal){
	a.Move()
}
func Speak_inter(a Animal){
	a.Speak()
}

func main() {
	cow := Cow{"grass","walk","moo"}
	bird := Bird{"worms","fly","peep"}
	snake := Snake{"mice","slither","hsss"}
	
	var command, name, request string

	animal := make(map[string]Animal)

	for {
		fmt.Println("<")
		fmt.Scanln(&command,&name,&request)
		if command == "newanimal"{
			switch request{
			case "cow":
				animal[name] = cow
			case "bird":
				animal[name] = bird
			case "snake":
				animal[name] = snake
			}
			fmt.Println("Created it!")
		} else if command == "query"{
			switch request{
			case "eat":
				Eat_inter(animal[name])
			case "move":
				Move_inter(animal[name])
			case "speak":
				Speak_inter(animal[name])
			}
		} else{
			fmt.Println("Unknow command!!")
		}
	}
}
