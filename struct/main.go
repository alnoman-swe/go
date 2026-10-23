package main 

import "fmt"


func main(){

	type person struct {
		name string 
		age int
	}

	fmt.Println(person{"Noman",27})
	a:= person{"Ali",30}
	
	a.age=31//change the value
	p:=&a//pointer to a
	p.age=32//change the value using pointer
	fmt.Println(a)
}
