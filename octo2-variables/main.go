package main

import (
	"fmt"
	// "math/cmplx"
)

// var c,python,java bool//same type can be declared in one time at the end
// var(
// 		beboolean bool = true
// 		MaxInt uint64 = 1<<64 - 1
// 		z  complex128= cmplx.Sqrt(-5 + 12i)

// 	)
//const can be only three types bool, string, numeric
//can not reassign the value of const
const Pi=3.1416
func main(){
	// var i int
	const myname ="al noman"
	// myname="abdullah"//error because const can not be reassigned
	//only inside of the function can be declared
	// a:=10 //same as var a = 10
	// var b string = "variable b"//if you type the variable while assigning the value must be the same type
	// fmt.Println(i,c,python,java,a,b)
	// fmt.Println("Type %T value %v\n",beboolean,beboolean)
	// fmt.Println("Type %T value %v\n",MaxInt,MaxInt)
	// fmt.Println("Type %T value %v\n",z,z)
	fmt.Println(Pi)
	fmt.Println(myname)
}