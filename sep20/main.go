package main

import (
	"fmt"
	"math"
	// "math/rand"
)

func swap(x,y string) (string,string){
	return y,x
}

//can return without writing the variable name
// func split(sum int)(x,y int){
// 	x=sum*4/9
// 	y=sum-x
// 	return//not readable so better to write return x,y explicitly
// }
func split (sum int)(int,int){
	x:=sum*4/9
	y:=sum-x
	return x,y

}

func add(x,y int) int{ //x int and y int or  x,y int
	return x+y
}
func main(){

	fmt.Println("Hello world")
	// fmt.Println("Random number:", rand.Intn(100))
	fmt.Println(math.Pi)//var so use Pi as P capital as its form a outside package
	fmt.Println(add(2,3))

	//swap func
	a,b := swap("hello","world")
	fmt.Println(a,b)

	//split func
	k,l := split(21)
	fmt.Println(k,l)
}