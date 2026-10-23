package main

import "fmt"
// type person struct {
// 	name string
// 	age int
// }
type Vertext struct {
		X int
		Y int
	}

	var (
		v1 = Vertext{1,2}
		v2 = Vertext{X:1}
		v3 = Vertext{}
		p = &Vertext{1,2}
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
