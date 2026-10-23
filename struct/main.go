package main

import "fmt"
// type person struct {
// 	name string
// 	age int
// }
type Vertex struct {
		X int//all fields must be capitalized to be exported
		Y int
	}

	var (
		v1 = Vertex{1,2}
		v2 = Vertex{X:1}
		v3 = Vertex{}
		p = &Vertex{1,2}
)

func main(){

	

	

	// fmt.Println(person{"Noman",27})
	// a:= person{"Ali",30}
	
	// a.age=31//change the value
	// p:=&a//pointer to a
	// p.age=32//change the value using pointer
	// fmt.Println(a)

	fmt.Println(v1,v2,v3,p)
}
