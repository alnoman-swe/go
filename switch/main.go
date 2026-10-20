package main

import (
	"fmt"
	"time"
)
// import "runtime"


func main(){
	fmt.Println("Hello, switch!")
	t:=time.Now()
	switch{
		case t.Hour()<12:
			fmt.Println("Good morning!")
		case t.Hour()<17:
			fmt.Println("Good afternoon")
		default:
			fmt.Println("Good evening.")

	}
	// switch os:=runtime.GOOS; os{
	// case "darwin":
	// 	fmt.Println("OS X.")
	// case "linux":
	// 	fmt.Println("Linux.")
	// default:
	// 	fmt.Printf("%s.\n",os)
	// }

}