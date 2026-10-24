package main

import "fmt"


func main(){
	names := [4]string{
		"Tanvir",
		"Rifat",
		"Shuvo",
		"Rasel",
	}
	fmt.Println(names)
	a:= names[0:2]
	b:=names[1:3] //here length is 2 and capacity is 3 and pointer is pointing to the second element of the array
	fmt.Println(a,b)
	
	b[1]="Noman"//slice is a reference type so if we change the value of slice b then it will also change the value of array names
	fmt.Println(a,b)
	fmt.Println(names)



}