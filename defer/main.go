package main

import "fmt"

func main(){
	// defer fmt.Println("Go first in the stack")//defer wil pop at when the function returns 
	// defer fmt.Println("World")
	// fmt.Println("Hello")
	fmt.Println("Counting")
	for i:=0;i<10;i++{
		defer fmt.Println(i)
	}
	fmt.Println("Done")
}